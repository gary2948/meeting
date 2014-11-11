package main

import (
	. "commonPackage"
	"commonPackage/model/clouddisk"
	"commonPackage/viewModel"
	"service/db"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"
)

//指定获取文件夹儿子数量
func GetFilesCount(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	uId := jsonBody.Get(JSON_USERID).MustInt64()
	pId := jsonBody.Get(JSON_PARENTID).MustInt64()
	reJson := simplejson.New()
	count, err := db.GetUserFolderFilesCount(uId, pId)
	if err != nil {
		Println(err.Error())
		reJson.Set(JSON_ERR, ErrSys)
		reJson.Set(JSON_OK, false)
	} else {
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_COUNT, count)
	}

	return reJson.Encode()

}

//指定获取文件夹大小
func GetFolderSize(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	uId := jsonBody.Get(JSON_USERID).MustInt64()
	pId := jsonBody.Get(JSON_PARENTID).MustInt64()
	reJson := simplejson.New()

	count, err := db.GetUserFolderSizes(uId, pId)
	if err != nil {
		Println(err.Error())
		reJson.Set(JSON_ERR, ErrSys)
		reJson.Set(JSON_OK, false)
	} else {
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_COUNT, count)
	}
	return reJson.Encode()
}

//添加文件夹
func AddFolder(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	uId := jsonBody.Get(JSON_USERID).MustInt64()
	pId := jsonBody.Get(JSON_PARENTID).MustInt64()
	fname, err := jsonBody.Get(JSON_FOLDERNAME).String()
	CheckError(err)
	reJson := simplejson.New()

	id, err := db.AddFloder(uId, pId, fname)
	if err != nil {
		reJson.Set(JSON_ERR, ErrSys)
		reJson.Set(JSON_OK, false)
	} else {
		reJson.Set(JSON_FLODERID, id)
		reJson.Set(JSON_OK, true)
	}

	return reJson.Encode()
}

func GetUploadtoken(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	fileKey, err := jsonBody.Get(JSON_FILEMAPID).String()
	CheckError(err)
	reJson := simplejson.New()
	uptoken := Uptoken(fileKey)
	reJson.Set(JSON_CONTENT, uptoken)
	reJson.Set(JSON_OK, true)
	return reJson.Encode()
}

func UpdateLinFileMapId(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	fileId := jsonBody.Get(JSON_FILEID).MustInt64(0)
	linkFileMapId, err := jsonBody.Get(JSON_LINKFILEMAPID).String()
	CheckError(err)
	reJson := simplejson.New()
	err = db.UpdateLinkFileMapId(fileId, linkFileMapId)
	if err != nil {
		reJson.Set(JSON_OK, false)
		reJson.Set(JSON_ERR, ErrFileId)
	} else {
		reJson.Set(JSON_OK, true)
	}

	return reJson.Encode()
}

func GetDownLoadUrl(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	fiileMapId := jsonBody.Get(JSON_FILEMAPID).MustInt64(0)
	key, has := db.GetFileMapKey(fiileMapId)
	reJson := simplejson.New()
	if has {
		url := DownloadUrl(key)
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_CONTENT, url)
	} else {
		reJson.Set(JSON_OK, false)
		reJson.Set(JSON_ERR, ErrFileId)
	}

	return reJson.Encode()
}

func GetCouldFileInfo(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	fileId := jsonBody.Get(JSON_FILEID).MustInt64(0)
	cfile := clouddisk.Lctb_cloudFiles{}
	reJson := simplejson.New()
	has, err := db.GetCloudFileById(fileId, &cfile)
	if has && err == nil {
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_CONTENT, cfile)
	} else {
		reJson.Set(JSON_OK, false)
		reJson.Set(JSON_ERR, err.Error())
	}

	return reJson.Encode()
}

func AddFileInfo(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	uId := jsonBody.Get(JSON_USERID).MustInt64(0)
	size := jsonBody.Get(JSON_FILESIZE).MustInt64(0)
	folderId := jsonBody.Get(JSON_FLODERID).MustInt64(0)
	fileName, _ := jsonBody.Get(JSON_FILENAME).String()
	fileMapId := jsonBody.Get(JSON_FILEMAPID).MustString("")
	linkFileMapId := jsonBody.Get(JSON_LINKFILEMAPID).MustString("")
	reJson := simplejson.New()

	cloudfileId, err := db.AddFile(uId, folderId, size, fileMapId, linkFileMapId, fileName)

	if err != nil {
		reJson.Set(JSON_OK, false)
		reJson.Set(JSON_ERR, err.Error())
	} else {
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_FILEID, cloudfileId)
	}

	return reJson.Encode()
}

func GetUserFolderViewFiles(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	uId := jsonBody.Get(JSON_USERID).MustInt64()
	fId := jsonBody.Get(JSON_FILEID).MustInt64()
	tag := jsonBody.Get(JSON_TAG).MustString("")
	reJson := simplejson.New()

	files := []clouddisk.Lctb_cloudFiles{}
	err := db.GetUserFolderViewFiles(uId, fId, &files)
	if err != nil {
		reJson.Set(JSON_ERR, ErrSys)
		reJson.Set(JSON_OK, false)
	} else {
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_CONTENT, files)
		reJson.Set(JSON_TAG, tag)
	}

	return reJson.Encode()
}

func GetSharedFiles(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	uId := jsonBody.Get(JSON_USERID).MustInt64()
	reJson := simplejson.New()

	files := []viewModel.ShareViewModel{}
	err := db.GetShareCloudFiles(uId, &files)
	if err != nil {
		reJson.Set(JSON_ERR, ErrSys)
		reJson.Set(JSON_OK, false)
	} else {
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_CONTENT, files)
	}

	return reJson.Encode()
}

func ShareFiles(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	uId := jsonBody.Get(JSON_USERID).MustInt64()
	sType := jsonBody.Get(JSON_SHAREFILETYPE).MustInt()
	fIds, _ := Int64Array(jsonBody.Get(JSON_FILEIDS))
	sname, err := jsonBody.Get(JSON_SHARENAME).String()
	CheckError(err)
	fileExt, _ := jsonBody.Get(JSON_FILEEXT).String()
	fileSize := jsonBody.Get(JSON_FILESIZE).MustInt64()
	reJson := simplejson.New()
	shareCode, err := db.ShareCloudFileByCode(uId, sType, fIds, sname, fileExt, fileSize)
	if err != nil {
		reJson.Set(JSON_ERR, ErrSys)
		reJson.Set(JSON_OK, false)
	} else {
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_CONTENT, shareCode)
	}

	return reJson.Encode()
}

func UnshareFiles(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	sId := jsonBody.Get(JSON_SHAREINFOID).MustInt64(0)
	reJson := simplejson.New()
	err := db.UnshareCloudFile(sId)
	if err != nil {
		reJson.Set(JSON_ERR, ErrSys)
		reJson.Set(JSON_OK, false)
	} else {
		reJson.Set(JSON_OK, true)
	}

	return reJson.Encode()
}

func GetRecycleFiles(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	uId := jsonBody.Get(JSON_USERID).MustInt64(0)
	reJson := simplejson.New()

	files := []clouddisk.Lctb_cloudFiles{}
	err := db.GetInRecycelCloudFile(uId, &files)
	if err != nil {
		reJson.Set(JSON_ERR, ErrSys)
		reJson.Set(JSON_OK, false)
	} else {
		reJson.Set(JSON_OK, true)
		reJson.Set(JSON_CONTENT, files)
	}

	return reJson.Encode()
}

func InRecycleFiles(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	uId := jsonBody.Get(JSON_USERID).MustInt64()
	fIds, err := Int64Array(jsonBody.Get(JSON_FILEIDS))
	reJson := simplejson.New()
	err = db.InRecycelCloudFile(uId, fIds)
	if err != nil {
		reJson.Set(JSON_ERR, ErrSys)
		reJson.Set(JSON_OK, false)
	} else {
		reJson.Set(JSON_OK, true)
	}

	return reJson.Encode()
}

func OutRecycleFiles(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	fIds, err := Int64Array(jsonBody.Get(JSON_FILEIDS))
	reJson := simplejson.New()
	err = db.OutRecycelCloudFile(fIds)
	if err != nil {
		reJson.Set(JSON_ERR, ErrSys)
		reJson.Set(JSON_OK, false)
	} else {
		reJson.Set(JSON_OK, true)
	}

	return reJson.Encode()
}

func DelRecycleFiles(ws *websocket.Conn, jsonBody *simplejson.Json) ([]byte, error) {
	fIds, err := Int64Array(jsonBody.Get(JSON_FILEIDS))
	reJson := simplejson.New()
	err = db.RecycelCloudFile(fIds)
	if err != nil {
		reJson.Set(JSON_ERR, ErrSys)
		reJson.Set(JSON_OK, false)
	} else {
		reJson.Set(JSON_OK, true)
	}

	return reJson.Encode()
}
