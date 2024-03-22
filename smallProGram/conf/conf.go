package conf

import "time"

const (
	DBHOST     = "127.0.0.1"
	DBPORT     = "3306"
	DBPASSWORD = "ljy12345..."
	DBUSERNAME = "root"
	DBDATABASE = "program"

	HASHIDSALT      = "salt"
	HASHIDMINLENGTH = 8

	REDISADDR = "175.178.72.197:6379"
	REDISPWD  = "ljy123..."
	REDISDB   = 0

	JWTISS       = "Eternal"
	JWTAUDIENCE  = "smallProGram"
	JWTJTI       = "Eternal"
	JWTSECRETKEY = "Eternal"
	JWTTOKENKEY  = "login:token:"
	JWTTOKENLIFE = time.Hour * time.Duration(72)
)

// Log
const (
	LOGFILEPATH     = "./log"
	LOGFILENAME     = "zog"
	LOGFILESUFFIX   = "log"
	LOGFILEMAXSIZE  = 0
	LOGFILEMAXNSIZE = 1
	LOGTIMEZONE     = "Asia/Chongqing"
)

const (
	BackUpDest        = "./backup"
	BackUpDuration    = "0 0 0 * * *"
	BackUpSqlFileName = "-sql-backup.sql"
	BackUpFilePath    = "./backup/"
)
