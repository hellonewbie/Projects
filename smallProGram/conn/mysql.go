package conn

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"smallProGarm/conf"
	"time"
)

//声明变量便于调用
var (
	SqlDb *gorm.DB
	err   error
)

type SqlParam struct {
	Host     string
	Port     string
	DataBase string
	UserName string
	Password string
}

//目的就是为了归类
type Sp func(*SqlParam) interface{}

func (p *Sp) SetDbHost(host string) Sp {
	return func(p *SqlParam) interface{} {
		h := p.Host
		p.Host = host
		return h
	}
}

func (p *Sp) SetDbPort(port string) Sp {
	return func(p *SqlParam) interface{} {
		pt := p.Port
		p.Port = port
		return pt
	}
}

func (p *Sp) SetDbDataBase(dataBase string) Sp {
	return func(p *SqlParam) interface{} {
		db := p.DataBase
		p.DataBase = dataBase
		return db
	}
}

func (p *Sp) SetDbPassword(pwd string) Sp {
	return func(p *SqlParam) interface{} {
		password := p.Password
		p.Password = pwd
		return password
	}

}

func (p *Sp) SetDbUserName(u string) Sp {
	return func(p *SqlParam) interface{} {
		name := p.UserName
		p.UserName = u
		return name
	}
}

func InitMysql(options ...Sp) (*gorm.DB, error) {
	q := &SqlParam{
		Host:     conf.DBHOST,
		Port:     conf.DBPORT,
		Password: conf.DBPASSWORD,
		DataBase: conf.DBDATABASE,
		UserName: conf.DBUSERNAME,
	}
	//遍历调用函数方法
	for _, option := range options {
		option(q)
	}
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", q.UserName, q.Password, q.Host, q.Port, q.DataBase)
	SqlDb, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	SqlDb.DB().SetMaxIdleConns(3)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	SqlDb.DB().SetMaxOpenConns(20)
	//设置连接可以复用的最大时间
	SqlDb.DB().SetConnMaxLifetime(0)
	SqlDb.SingularTable(true)
	//设置周期触发定时的计时器它会按照一个时间间隔往channel发送系统当前时间
	timer := time.NewTicker(time.Minute * 10)
	go func(conn *gorm.DB) {
		for _ = range timer.C {
			if err = SqlDb.DB().Ping(); err != nil {
				MySQLAutoConnect()
			}
		}
	}(SqlDb)
	return SqlDb, err
}

func autoConnectMySQL(tryTimes int, maxTryTimes int) int {
	tryTimes++
	if tryTimes <= maxTryTimes {
		if SqlDb.DB().Ping() != nil {
			message := fmt.Sprintf("数据库连接失败，已重连%d次", tryTimes)
			log.Print(message)
		}
		tryTimes = autoConnectMySQL(tryTimes, maxTryTimes)
	}
	return tryTimes
}

func MySQLAutoConnect() {
	autoConnectMySQL(0, 0)
}
