package db

import (
	"commonPackage"
	"commonPackage/model/clouddisk"
	"commonPackage/viewModel"
	"errors"
	"fmt"
	"service/uuid"
	"strings"
	"time"

	"github.com/go-xorm/xorm"
)

var cf *clouddisk.Lctb_cloudFiles = &clouddisk.Lctb_cloudFiles{}
var fm *clouddisk.Lctb_filesMap = &clouddisk.Lctb_filesMap{}
var sf *clouddisk.Lctb_sharedFiles = &clouddisk.Lctb_sharedFiles{}
var su *clouddisk.Lctb_sharedUsers = &clouddisk.Lctb_sharedUsers{}
var si *clouddisk.Lctb_sharedInfos = &clouddisk.Lctb_sharedInfos{}
var ucs *clouddisk.Lctb_userCloudSize = &clouddisk.Lctb_userCloudSize{}
var rf *clouddisk.Lctb_recycleBinFiles = &clouddisk.Lctb_recycleBinFiles{}

func InitCloudDiskTables() {
	fmt.Println("-----InitCloudDiskTables -----")
	engine, err := postgresEngine(clouddiskDB)
	checkError(err)
	defer engine.Close()

	err = engine.Sync(cf)
	checkError(err)

	err = engine.Sync(fm)
	checkError(err)

	err = engine.Sync(su)
	checkError(err)

	err = engine.Sync(si)
	checkError(err)

	err = engine.Sync(sf)
	checkError(err)

	err = engine.Sync(ucs)
	checkError(err)

	err = engine.Sync(rf)
	checkError(err)
}

//新建文件夹
func AddFloder(userId, parentId int64, folderName string) (int64, error) {
	cfile := &clouddisk.Lctb_cloudFiles{}
	cParentfile := &clouddisk.Lctb_cloudFiles{}

	//判断父亲存在不
	if parentId != 0 {
		_, err := GetCloudFileById(parentId, cParentfile)
		if err != nil {
			return 0, err
		}
	}
	//检查文件名的有效性
	err := checkFileName(userId, parentId, folderName)
	if err != nil {
		return 0, err
	}

	cfile.Lc_userInfoId = userId
	cfile.Lc_createTime = time.Now()
	cfile.Lc_fileName = folderName
	cfile.Lc_fileType = 1 //1表示文件夹
	cfile.Lc_parentId = parentId
	_, err = addCloudFile(cfile)
	if err != nil {
		return 0, err
	}
	return cfile.Id, nil
}

func GetFileMapKey(fileMapId int64) (string, bool) {
	engine, err := GetCloudDiskEng()
	checkError(err)

	fm := clouddisk.Lctb_filesMap{}
	has, err := engine.Id(fileMapId).Get(&fm)
	if has && err == nil {
		return fm.Lc_fileSystemId, true
	} else {
		return "", false
	}
}

func getFileMapId(fileMapKey string) (int64, bool) {
	engine, err := GetCloudDiskEng()
	checkError(err)

	fm := clouddisk.Lctb_filesMap{}
	has, err := engine.Where("lc_file_system_id = ?", fileMapKey).Get(&fm)
	if has && err == nil {
		return fm.Id, has
	} else {
		return 0, false
	}
}

func UpdateLinkFileMapId(fileId int64, linkFilemapId string) error {
	engine, err := GetCloudDiskEng()
	checkError(err)

	cfile := clouddisk.Lctb_cloudFiles{}
	has, err := engine.Id(fileId).Get(&cfile)
	if err == nil && has {
		lfid, has := getFileMapId(linkFilemapId)
		if has {
			cfile.Lc_linkFilesMapId = lfid
			_, err = engine.Update(&cfile)
		} else {
			return commonPackage.NewErr(commonPackage.ErrFileId)
		}

	} else {
		return commonPackage.NewErr(commonPackage.ErrFileId)
	}

	return err
}

func AddFile(userId, parentId, size int64, fileMapId, linkFilemapId, fileName string) (int64, error) {
	cfile := &clouddisk.Lctb_cloudFiles{}
	cParentfile := &clouddisk.Lctb_cloudFiles{}

	fmId, has := getFileMapId(fileMapId) //文件必须有
	if !has {
		return 0, commonPackage.NewErr(commonPackage.ErrSys)
	}

	lfmId, _ := getFileMapId(linkFilemapId) //附属文件不一定有

	//判断父亲存在不
	if parentId != 0 {
		_, err := GetCloudFileById(parentId, cParentfile)
		if err != nil {
			return 0, err
		}
	}
	ss := strings.Split(fileName, ".")
	//检查文件名的有效性
	//err := checkFileName(userId, parentId, ss[0])
	//if err != nil {
	//	return 0, err
	//}

	cfile.Lc_userInfoId = userId
	cfile.Lc_createTime = time.Now()
	cfile.Lc_fileName = ss[0]
	cfile.Lc_fileExtension = ss[1]
	cfile.Lc_fileType = 0 //0表示文件
	cfile.Lc_parentId = parentId
	cfile.Lc_filesMapId = fmId
	cfile.Lc_linkFilesMapId = lfmId
	cfile.Lc_fileSize = size
	_, err := addCloudFile(cfile)
	if err != nil {
		return 0, err
	}
	addFileMapCount(fmId)
	return cfile.Id, nil
}

func addFileMapCount(fileMapId int64) (err error) {
	engine, err := GetCloudDiskEng()
	checkError(err)
	fm := &clouddisk.Lctb_filesMap{}
	has, err := engine.Id(fileMapId).Get(fm)
	if err != nil {
		return commonPackage.NewErr(commonPackage.ErrSys)
	}

	if has {
		fm.Lc_usedCount = fm.Lc_usedCount + 1
		engine.Id(fm.Id).Update(fm)
	}

	return err
}

func RenameCloudFile(userId, parentId, fileid int64, rename string) error {
	engine, err := GetCloudDiskEng()
	checkError(err)

	err = checkFileName(userId, parentId, rename)

	if err == nil {
		f := clouddisk.Lctb_cloudFiles{Lc_fileName: rename}
		_, err = engine.Id(fileid).Cols("lc_file_name").Update(f)
	}

	return err
}

//判断文件名是否重复
func checkFileName(userId, parentId int64, folderName string) error {
	cfiles := &[]clouddisk.Lctb_cloudFiles{}
	err := GetUserFolderViewFiles(userId, parentId, cfiles)
	if err != nil {
		return err
	}
	//判断文件名重复
	for _, v := range *cfiles {
		if v.Lc_fileName == folderName {
			err = errors.New(commonPackage.ErrFileName)
			return err
		}
	}

	return nil
}

// Add a new cloudfiles
func addCloudFile(cfile *clouddisk.Lctb_cloudFiles) (int64, error) {
	engine, err := GetCloudDiskEng()
	checkError(err)

	return engine.Insert(cfile)
}

//得到用户文件夹下所有文件
func GetUserFolderFiles(userId, parentId int64, cfiles *[]clouddisk.Lctb_cloudFiles) (err error) {
	engine, err := GetCloudDiskEng()
	checkError(err)

	return engine.Where("lc_user_info_id = ? and lc_parent_id = ? ", userId, parentId).Find(cfiles)
}

//得到用户文件夹下文件大小
func GetUserFolderSizes(userId, parentId int64) (size int64, err error) {
	cfiles := &[]clouddisk.Lctb_cloudFiles{}
	err = GetUserFolderFiles(userId, parentId, cfiles)
	if err != nil {
		return 0, err
	}
	//算文件夹大小
	var count int64
	for _, v := range *cfiles {
		count = count + v.Lc_fileSize
	}

	return count, nil
}

//得到用户文件夹下所有文件数量
func GetUserFolderFilesCount(userId, parentId int64) (count int64, err error) {
	engine, err := GetCloudDiskEng()
	checkError(err)
	var f clouddisk.Lctb_cloudFiles
	return engine.Where("lc_user_info_id = ? and lc_parent_id = ?", userId, parentId).Count(f)
}

//得到单个文件
func GetCloudFileById(fileId int64, cfile *clouddisk.Lctb_cloudFiles) (bool, error) {
	engine, err := GetCloudDiskEng()
	checkError(err)

	has, err := engine.Id(fileId).Get(cfile)

	if !has {
		err = errors.New(commonPackage.ErrFileId)
	}

	return has, err
}

func addUserCloudSize(userCloud *clouddisk.Lctb_userCloudSize) (id int64, err error) {
	engine, err := GetCloudDiskEng()
	checkError(err)

	_, err = engine.Insert(userCloud)

	return userCloud.Id, err
}

//更新用户云盘大小
func UpdateUserCloudSize(userCloud *clouddisk.Lctb_userCloudSize) error {
	engine, err := GetCloudDiskEng()
	checkError(err)

	_, err = engine.Id(userCloud.Id).Update(userCloud)

	return err
}

//得到用户空间大小
func GetUserSizeByUserId(userId int64) (maxSize, usedSzie int64, err error) {
	engine, err := GetCloudDiskEng()
	checkError(err)

	u := &clouddisk.Lctb_userCloudSize{}
	has, err := engine.Where("lc_user_info_id = ?", userId).Get(u)
	commonPackage.CheckError(err)
	if !has {
		err = commonPackage.NewErr(commonPackage.ErrUserId)
	} else if err != nil {
		err = commonPackage.NewErr(commonPackage.ErrSys)
	} else {
		maxSize = u.Lc_maxCloudSize
		usedSzie = u.Lc_usedCloudSize
	}

	return maxSize, usedSzie, err
}

//得到用户文件夹下所有可视文件，不包括回收站和已经删除的文件
func GetUserFolderViewFiles(userId, parentId int64, cfiles *[]clouddisk.Lctb_cloudFiles) (err error) {
	engine, err := GetCloudDiskEng()
	checkError(err)

	return engine.Where("lc_user_info_id = ? and lc_parent_id = ? and ( lc_file_status = ? or lc_file_status = ?)", userId, parentId, commonPackage.CommFile, commonPackage.ShareFile).Find(cfiles)
}

//得到回收站内文件
func GetInRecycelCloudFile(userId int64, cfiles *[]clouddisk.Lctb_cloudFiles) error {
	engine, err := GetCloudDiskEng()
	checkError(err)

	rfiles := []clouddisk.Lctb_recycleBinFiles{}
	err = engine.Where("lc_user_info_id = ? ", userId).Find(&rfiles)
	commonPackage.CheckError(err)

	fileIds := make([]int64, len(rfiles))

	for i, v := range rfiles {
		fileIds[i] = v.Lc_cloudFilesId
	}

	return engine.In("id", fileIds).Find(cfiles)
}

//处理由于删除时候需要处理共享的问题
func unShareByRecycle(engine *xorm.Engine, cloudFile *clouddisk.Lctb_cloudFiles) {
	if cloudFile.Lc_fileStatus == commonPackage.ShareFile {
		sf := clouddisk.Lctb_sharedFiles{}
		_, err := engine.Where("lc_cloud_files_id = ?", cloudFile.Id).Get(&sf)
		commonPackage.CheckError(err)
		si := clouddisk.Lctb_sharedInfos{}
		_, err = engine.Id(sf.Lc_sharedInfoId).Get(&si)
		commonPackage.CheckError(err)
		//如果是分享的单个文件，直接删除
		if si.Lc_sharedFileType == commonPackage.OneFileShare ||
			si.Lc_sharedFileType == commonPackage.OneFolderShare {
			//del shareinfo
			engine.Id(si.Id).Delete(si)
			//del users
			_, err = engine.Where("lc_shared_info_id = ?", si.Id).Delete(su)
			//del share files
			_, err = engine.Where("lc_shared_info_id = ?", si.Id).Delete(sf)
		} else { //只删除自己
			//del share files
			_, err = engine.Where("lc_cloud_files_id = ?", cloudFile.Id).Delete(sf)
		}

	}
}

//放入回收站
func InRecycelCloudFile(userId int64, fileIds []int64) error {
	engine, err := GetCloudDiskEng()
	checkError(err)

	fins := make([]clouddisk.Lctb_recycleBinFiles, len(fileIds))
	for i, v := range fileIds {
		of := &clouddisk.Lctb_cloudFiles{}
		has, err := engine.Id(v).Get(of)
		if has && err == nil {
			if of.Lc_fileType == commonPackage.File { //文件直接删除
				//设置状态
				f := clouddisk.Lctb_cloudFiles{Lc_fileStatus: commonPackage.RecycleFile}
				_, err = engine.Id(v).Cols("lc_file_status").Update(f)

			} else if of.Lc_fileType == commonPackage.Folder { //文件夹递归删除子文件以及目录
				sql := `WITH RECURSIVE  r AS ( 
       					SELECT * FROM lctb_cloud_files WHERE id = ?
     					union   ALL 
				        SELECT lctb_cloud_files.* FROM lctb_cloud_files, r WHERE lctb_cloud_files.lc_parent_id = r.id AND lctb_cloud_files.lc_file_status != ?
     				   ) 
					   UPDATE lctb_cloud_files f SET lc_file_status = ? FROM r WHERE f.id = r.id`
				_, err = engine.Exec(sql, v, commonPackage.DelFile, commonPackage.RecycleFile)
			}
			fin := clouddisk.Lctb_recycleBinFiles{}
			fin.Lc_cloudFilesId = of.Id
			fin.Lc_userInfoId = of.Lc_userInfoId
			fins[i] = fin
			//如果是分享的文件需要单独处理
			unShareByRecycle(engine, of)
		}
	}
	_, err = engine.Insert(fins)
	return err
}

//回复回收站文件
func OutRecycelCloudFile(fileIds []int64) error {
	engine, err := GetCloudDiskEng()
	checkError(err)

	for _, v := range fileIds {
		of := &clouddisk.Lctb_cloudFiles{}
		has, err := engine.Id(v).Get(of)
		if has && err == nil {
			if of.Lc_fileType == commonPackage.File { //文件直接恢复
				f := clouddisk.Lctb_cloudFiles{Lc_fileStatus: commonPackage.CommFile}
				_, err = engine.Id(v).Cols("lc_file_status").Update(f)
			} else if of.Lc_fileType == commonPackage.Folder { //文件夹递归回复子文件以及目录,已彻底删除的不恢复
				sql := `WITH RECURSIVE  r AS ( 
       					SELECT * FROM lctb_cloud_files WHERE id = ?
     					union   ALL 
				        SELECT lctb_cloud_files.* FROM lctb_cloud_files, r WHERE lctb_cloud_files.lc_parent_id = r.id AND lctb_cloud_files.lc_file_status != ?
     				   ) 
					   UPDATE lctb_cloud_files f SET lc_file_status = ? FROM r WHERE f.id = r.id`
				_, err = engine.Exec(sql, v, commonPackage.DelFile, commonPackage.CommFile)
			}

			//恢复上级目录
			sql := `WITH RECURSIVE r AS (
			    				SELECT * FROM lctb_cloud_files WHERE id = ?
			  				union   ALL
			 			    SELECT lctb_cloud_files.* FROM lctb_cloud_files, r WHERE lctb_cloud_files.id = r.lc_parent_id
			  				)
					UPDATE lctb_cloud_files f SET lc_file_status = ? FROM r WHERE f.id = r.id`
			_, err = engine.Exec(sql, v, commonPackage.CommFile)

		}
	}
	_, err = engine.In("lc_cloud_files_id", fileIds).Delete(rf)
	return err
}

//彻底删除文件
func RecycelCloudFile(fileIds []int64) error {
	engine, err := GetCloudDiskEng()
	checkError(err)

	fin := clouddisk.Lctb_recycleBinFiles{}

	//删除回收站记录
	_, err = engine.In("lc_cloud_files_id", fileIds).Delete(fin)
	//设置删除状态
	f := clouddisk.Lctb_cloudFiles{Lc_fileStatus: commonPackage.DelFile}
	_, err = engine.In("id", fileIds).Cols("lc_file_status").Update(f)

	return err
}

//分享文件-提取码
func ShareCloudFileByCode(userId int64, sharefileType int, fileIds []int64, shareName, fileExt string, fileSize int64) (shareCode string, err error) {
	engine, err := GetCloudDiskEng()
	checkError(err)

	//add shareinfo
	shareCode, _ = uuid.GenUUID4()
	si := &clouddisk.Lctb_sharedInfos{}
	si.Lc_userInfoId = userId
	si.Lc_sharedName = shareName
	si.Lc_sharedExt = fileExt
	si.Lc_shareCode = shareCode
	si.Lc_sharedSize = fileSize
	si.Lc_sharedType = commonPackage.ShareByCode
	si.Lc_sharedFileType = sharefileType
	_, err = engine.Insert(si)
	commonPackage.CheckError(err)
	sid := si.Id

	//add share files
	sfs := make([]clouddisk.Lctb_sharedFiles, len(fileIds))
	for i, v := range fileIds {
		sf := clouddisk.Lctb_sharedFiles{}
		sf.Lc_sharedInfoId = sid
		sf.Lc_cloudFilesId = v
		sfs[i] = sf
	}
	_, err = engine.Insert(&sfs)

	//update files status
	cf := &clouddisk.Lctb_cloudFiles{Lc_fileStatus: commonPackage.ShareFile}
	engine.In("id", fileIds).Cols("lc_file_status").Update(cf)

	return
}

//分享文件-给用户
func ShareCloudFileByUser(userId int64, userIds, fileIds []int64, shareName, fileExt string, fileSize int64) (err error) {
	engine, err := GetCloudDiskEng()
	checkError(err)

	//add shareinfo
	si := &clouddisk.Lctb_sharedInfos{}
	si.Lc_userInfoId = userId
	si.Lc_sharedName = shareName
	si.Lc_sharedExt = fileExt
	si.Lc_sharedSize = fileSize
	si.Lc_sharedType = commonPackage.ShareByUser
	_, err = engine.Insert(si)
	commonPackage.CheckError(err)
	sid := si.Id
	//add share files
	sfs := make([]clouddisk.Lctb_sharedFiles, len(fileIds))
	for i, v := range fileIds {
		sf := clouddisk.Lctb_sharedFiles{}
		sf.Lc_sharedInfoId = sid
		sf.Lc_cloudFilesId = v
		sfs[i] = sf
	}
	_, err = engine.Insert(&sfs)

	//add share users
	sus := make([]clouddisk.Lctb_sharedUsers, len(userIds))
	for i, v := range userIds {
		su := clouddisk.Lctb_sharedUsers{}
		su.Lc_sharedInfoId = sid
		su.Lc_userInfoId = userId
		su.Lc_toUserInfoId = v
		sus[i] = su
	}
	_, err = engine.Insert(&sus)
	//update files status
	cf := &clouddisk.Lctb_cloudFiles{Lc_fileStatus: commonPackage.ShareFile}
	engine.In("id", fileIds).Cols("lc_file_status").Update(cf)

	return
}

//得到已分享文件
func GetShareCloudFiles(userId int64, shareInfos *[]viewModel.ShareViewModel) error {
	engine, err := GetCloudDiskEng()
	checkError(err)

	si := &clouddisk.Lctb_sharedInfos{}
	return engine.Table(si).Where("lc_user_info_id = ?", userId).Find(shareInfos)
}

func GetShareInfo(shareInfoId int64, shareInfo *viewModel.ShareViewModel) error {
	engine, err := GetCloudDiskEng()
	checkError(err)

	si := &clouddisk.Lctb_sharedInfos{}
	_, err = engine.Where("id = ?", shareInfoId).Get(si)
	shareInfo.DownloadCount = si.Lc_downloadCount
	shareInfo.FILEEXT = si.Lc_sharedExt
	shareInfo.ShareCode = si.Lc_shareCode
	shareInfo.ShareFileType = si.Lc_sharedFileType
	shareInfo.ShareId = si.Id
	shareInfo.ShareName = si.Lc_sharedName
	shareInfo.ShareSize = si.Lc_sharedSize
	shareInfo.ShareTime = si.Lc_sharedTime
	shareInfo.ShareUser = si.Lc_userInfoId

	return err

}

//取消分享
func UnshareCloudFile(shareInfoId int64) error {
	engine, err := GetCloudDiskEng()
	checkError(err)

	//del shareinfo
	si := &clouddisk.Lctb_sharedInfos{}
	engine.Id(shareInfoId).Delete(si)

	//del users
	su := clouddisk.Lctb_sharedUsers{}
	_, err = engine.Where("lc_shared_info_id = ?", shareInfoId).Delete(su)

	//del files
	sfs := make([]clouddisk.Lctb_sharedFiles, 0)
	//sf := clouddisk.Lctb_sharedFiles{}
	err = engine.Where("lc_shared_info_id = ?", shareInfoId).Find(&sfs)
	fileIds := make([]int64, len(sfs))
	for i, v := range sfs {
		fileIds[i] = v.Lc_cloudFilesId
	}
	_, err = engine.Where("lc_shared_info_id = ?", shareInfoId).Delete(sf)

	//update files status
	cf := &clouddisk.Lctb_cloudFiles{Lc_fileStatus: commonPackage.CommFile}
	_, err = engine.In("id", fileIds).Cols("lc_file_status").Update(cf)

	return err
}

//得到分享文件
func GetShareFiles(shareInfoId int64, files *[]clouddisk.Lctb_cloudFiles) error {
	engine, err := GetCloudDiskEng()
	checkError(err)

	sfs := []clouddisk.Lctb_sharedFiles{}
	err = engine.Where("lc_shared_info_id = ?", shareInfoId).Find(&sfs)
	commonPackage.CheckError(err)
	n := len(sfs)
	if n > 0 {
		fileIds := make([]int64, n)
		for i, v := range sfs {
			fileIds[i] = v.Lc_cloudFilesId

		}
		err = engine.In("id", fileIds).Find(files)
		commonPackage.CheckError(err)
	}

	return err
}

//清理所有
func ClearAllRecycle(userId int64) error {
	engine, err := GetCloudDiskEng()
	checkError(err)

	fin := clouddisk.Lctb_recycleBinFiles{}
	files := []clouddisk.Lctb_recycleBinFiles{}
	err = engine.Where("lc_user_info_id = ?", userId).Find(&files)

	//删除回收站记录
	_, err = engine.Where("lc_user_info_id = ?", userId).Delete(fin)
	//设置删除状态
	n := len(files)
	if n > 0 {
		fileIds := make([]int64, n)
		for i, v := range files {
			fileIds[i] = v.Lc_cloudFilesId
		}

		f := clouddisk.Lctb_cloudFiles{Lc_fileStatus: commonPackage.DelFile}
		_, err = engine.In("id", fileIds).Cols("lc_file_status").Update(f)
	}

	return err
}

//查找文件是否存在,如果存在，返回id
func IsFilesMapExist(fileSystemId string) (has bool, id int64, err error) {
	engine, err := GetCloudDiskEng()
	checkError(err)
	fm := clouddisk.Lctb_filesMap{}
	has, err = engine.Where("lc_file_system_id  = ?", fileSystemId).Get(&fm)
	if !has {
		return has, 0, err
	}
	return has, fm.Id, err
}

//添加文件
func AddFileMap(fileSystemId string) (id int64, err error) {
	engine, err := GetCloudDiskEng()
	checkError(err)

	fm := clouddisk.Lctb_filesMap{}

	has, err := engine.Id(fileSystemId).Get(&fm)
	if err == nil {
		if has {
			return fm.Id, nil
		} else {
			fm.Lc_fileSystemId = fileSystemId
			fm.Lc_usedCount = 1
			_, err = engine.Insert(&fm)
		}
	} else {
		return 0, err
	}

	return fm.Id, err
}
