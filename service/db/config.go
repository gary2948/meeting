package db

import (
	"commonPackage"
	"errors"
	"fmt"

	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const (
	drive        = "mysql"
	sqlitedrive  = "sqlite3"
	pqdrive      = "postgres"
	userName     = "root"
	passwd       = "qwe123!@#"
	server       = "tcp(db.lichengsoft.com:3306)"
	helperEmail  = "lchelper@lichengsoft.com"
	helpPasswd   = "qwe123!@#"
	redisNetwork = "tcp"
	redisAddr    = "db.lichengsoft.com:6378"
	//库
	accountDB   = "lcdb_account"
	clouddiskDB = "lcdb_clouddisk"
	meetingDB   = "lcdb_meeting"
	socialDB    = "lcdb_social"
)

var userSqliteEng *xorm.Engine
var accountEng *xorm.Engine
var clouddiskEng *xorm.Engine
var meetingEng *xorm.Engine
var socialEng *xorm.Engine
var cloudFilesEng *xorm.Engine
var redisConn redis.Conn

func GetRedisConn() redis.Conn {
	if redisConn == nil {
		rc, err := redis.Dial(redisNetwork, redisAddr)
		commonPackage.CheckError(err)
		redisConn = rc
	}
	return redisConn
}

func GetUserSqliteEng() *xorm.Engine {
	if userSqliteEng == nil {
		userSqliteEng, _ = sqliteEngine("userCash.db")
	}
	return userSqliteEng
}

func GetCloudFileCashEng() *xorm.Engine {
	if cloudFilesEng == nil {
		cloudFilesEng, _ = sqliteEngine("cloudFiles.db")
	}
	return cloudFilesEng
}

func GetAccountEng() (*xorm.Engine, error) {
	if accountEng == nil {
		accountEng, err := postgresEngine(accountDB)
		return accountEng, err
	}
	return accountEng, nil
}

func GetCloudDiskEng() (*xorm.Engine, error) {
	if clouddiskEng == nil {
		clouddiskEng, err := postgresEngine(clouddiskDB)
		return clouddiskEng, err
	}
	return clouddiskEng, nil
}

func GetMeetingEng() (*xorm.Engine, error) {
	if meetingEng == nil {
		meetingEng, err := postgresEngine(meetingDB)
		return meetingEng, err
	}
	return meetingEng, nil
}

func GetSocialEng() (*xorm.Engine, error) {
	if socialEng == nil {
		socialEng, err := postgresEngine(socialDB)
		return socialEng, err
	}
	return socialEng, nil
}

func sqliteEngine(dbName string) (*xorm.Engine, error) {
	orm, err := xorm.NewEngine(sqlitedrive, dbName)
	return orm, err
}

func mysqlEngine(dbName string) (*xorm.Engine, error) {
	connStr := userName + ":" + passwd + "@" + server + "/" + dbName + "?charset=utf8"
	orm, err := xorm.NewEngine(drive, connStr)
	return orm, err
}

func postgresEngine(dbName string) (*xorm.Engine, error) {
	//pqconnStr := "user=postgres password=qwe123!@# port=5433 dbname=" + dbName + " host=db.lichengsoft.com sslmode=disable" //内网
	pqconnStr := "user=postgres password=qwe123!@# port=5432 dbname=" + dbName + " host=115.29.172.111 sslmode=disable" //外网
	return xorm.NewEngine(pqdrive, pqconnStr)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		err = errors.New(commonPackage.ErrSys)
		return
	}
}
