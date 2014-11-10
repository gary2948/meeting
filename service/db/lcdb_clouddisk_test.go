package db

import (
	"commonPackage"
	"commonPackage/model/clouddisk"
	"commonPackage/viewModel"
	"fmt"
	"testing"
)

//func TestAddFolder(t *testing.T) {
//	//
//	_, err := AddFloder(1, -1, "test")
//	if err == nil {
//		t.Errorf("err")
//	}

//	floderName := commonPackage.GetUnixTime()
//	var i int64
//	i, err = AddFloder(1, 0, floderName)
//	if err != nil {
//		t.Errorf("err")
//	}

//	fmt.Printf("new folderId:%d\n", i)

//	_, err = AddFloder(1, 0, "test")
//	if err == nil {
//		t.Errorf("err")
//	}

//}

//func TestAddFile(t *testing.T) {

//	_, err := AddFile(1, 0, "test.txt")
//	if err != nil {
//		commonPackage.Printf(err)
//		t.Errorf("err")
//	}

//}

func TestFolderCount(t *testing.T) {
	//
	i, err := GetUserFolderFilesCount(1, 0)
	if err != nil {
		t.Errorf("err")
	}
	fmt.Printf("files count:%d\n", i)
}

func TestFolderSize(t *testing.T) {
	//
	i, err := GetUserFolderSizes(1, 0)
	if err != nil {
		t.Errorf("err")
	}
	fmt.Printf("folder size:%d\n", i)
}

//func TestGetUserSizeByUserId(t *testing.T) {
//	maxS, usedS, err := GetUserSizeByUserId(9)
//	if err != nil {
//		t.Errorf("err")
//	}
//	fmt.Printf("Max:%d;Used:%d\n", maxS, usedS)

//}

//func TestRenameFiles(t *testing.T) {
//	f := &clouddisk.Lctb_cloudFiles{}
//	_, err := GetCloudFileById(1, f)
//	if err != nil {
//		t.Errorf("err")
//	}

//	err = RenameCloudFile(1, 0, 1, f.Lc_fileName)
//	if err == nil {
//		t.Errorf("err")
//	}

//	err = RenameCloudFile(1, 0, 1, f.Lc_fileName+"1")
//	if err != nil {
//		t.Errorf("err")
//	}

//}

func TestGetUserFolderViewFiles(t *testing.T) {
	f := &[]clouddisk.Lctb_cloudFiles{}
	err := GetUserFolderViewFiles(1, 80, f)
	if err != nil {
		t.Errorf("err")
	}
	//fmt.Println(len(*f))
	commonPackage.Printf(f)
}

//func TestInRecycelCloudFile(t *testing.T) {
//	files := []int64{2}
//	err := InRecycelCloudFile(1, files)
//	if err != nil {
//		commonPackage.Printf(err)
//		t.Error(err)
//	}
//}

//func TestOutRecycelCloudFile(t *testing.T) {
//	files := []int64{2, 3}
//	err := OutRecycelCloudFile(files)
//	if err != nil {
//		t.Error(err)
//	}
//}

//func TestRecycelCloudFile(t *testing.T) {
//	files := []int64{8}
//	err := RecycelCloudFile(files)
//	if err != nil {
//		t.Error(err)
//	}
//}

//func TestShareCloudFileByCode(t *testing.T) {
//	files := []int64{2}

//	code, err := ShareCloudFileByCode(1, commonPackage.OneFileShare, files, "share2", "txt", 0)
//	commonPackage.Println(code)
//	if err != nil {
//		commonPackage.Printf(err)
//		t.Error(err)
//	}
//}

//func TestShareCloudFileByUsers(t *testing.T) {
//	files := []int64{1}
//	users := []int64{4, 5}
//	err := ShareCloudFileByUser(1, users, files, "testsharecode", "txt", 2048)

//	if err != nil {
//		t.Error(err)
//	}
//}

//func TestUnShareCloudFile(t *testing.T) {

//	err := UnshareCloudFile(3)

//	if err != nil {
//		t.Error(err)
//	}
//}

//func TestGetShareCloudFile(t *testing.T) {

//	sfiles := []viewModel.ShareViewModel{}
//	err := GetShareCloudFiles(1, &sfiles)

//	if err != nil {
//		t.Error(err)
//	}

//	commonPackage.Printf(sfiles)
//}

//func TestClearAll(t *testing.T) {

//	err := ClearAllRecycle(1)

//	if err != nil {
//		t.Error(err)
//	}
//}

func TestGetShareInfo(t *testing.T) {
	shareInfo := viewModel.ShareViewModel{}
	err := GetShareInfo(90, &shareInfo)
	if err != nil {
		t.Error(err)
	}
	commonPackage.Printf(shareInfo)
}

//func TestAddFile(t *testing.T) {
//	_, err := AddFile(1, 1, "111.txt", "2222222")
//	if err != nil {
//		t.Error(err)
//	}
//}
