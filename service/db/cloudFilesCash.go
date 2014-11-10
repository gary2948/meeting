package db

import (
	"commonPackage"
)

const (
	DB = "cloudFiles.db"
)

type UploadingFile struct {
	UserId   int64
	FolderId int64
	Position int64
	Size     int64
	Context  []byte
	FileId   int64
	EOF      bool   `xorm:"'EOF'"`
	CRC      string `xorm:"'CRC'" unique`
}

var f *UploadingFile = &UploadingFile{}

func IniTable() {
	commonPackage.Println("-----IniCloudFileCash-----")
	engine, err := sqliteEngine(DB)
	commonPackage.CheckError(err)
	defer engine.Close()

	err = engine.Sync(f)
	commonPackage.CheckError(err)
}

func AddCash(f *UploadingFile) error {
	engine, err := sqliteEngine(DB)
	commonPackage.CheckError(err)
	defer engine.Close()

	_, err = engine.Insert(f)
	commonPackage.CheckError(err)
	return err
}

func GetCash(userId int64, key string, f *UploadingFile) (has bool, err error) {
	engine := GetCloudFileCashEng()
	return engine.Where("CRC = ? and user_id = ?", key, userId).Get(f)
}

func UpdateCash(f *UploadingFile) error {
	engine := GetCloudFileCashEng()
	_, err := engine.Update(f)
	return err
}
