package model

import (
	"fmt"
	"ginblog/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 声明变量便于调用
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

// 目的就是为了归类，说白了就是将这些内容归纳到Sp中，遍历执行函数，将自己所设置的数据库参数传入进去
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

func InitMysql(options ...Sp) {
	q := &SqlParam{
		Host:     utils.DbHost,
		Port:     utils.DbPort,
		Password: utils.DbPassWord,
		DataBase: utils.DbName,
		UserName: utils.DbUser,
	}
	//遍历调用函数方法
	for _, option := range options {
		option(q)
	}
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", q.UserName, q.Password, q.Host, q.Port, q.DataBase)
	fmt.Println(dataSourceName)
	SqlDb, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	// AutoMigrate 会创建表、缺失的外键、约束、列和索引。 如果大小、精度、是否为空可以更改，则 AutoMigrate 会改变列的类型。
	SqlDb.SingularTable(true)
	SqlDb.AutoMigrate(&User{}, &Article{}, &Category{})
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	SqlDb.DB().SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	SqlDb.DB().SetMaxOpenConns(20)
	//设置连接可以复用的最大时间
	SqlDb.DB().SetConnMaxLifetime(0)
	SqlDb.SingularTable(true)
	//SqlDb.Close()
}
