package main

import (
	//"commonPackage"
	//"commonPackage/viewModel"
	"fmt"
	"runtime"
	"service/db"
)

func main() {
	runtime.GOMAXPROCS(2)

	fmt.Println("-----start syn all tables -----")
	db.InitAccountTables()
	db.InitCloudDiskTables()
	db.InitMeetingTables()
	db.InitSocialTables()
	db.IniTable()
}
