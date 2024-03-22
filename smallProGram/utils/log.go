package utils

import (
	"errors"
	"github.com/Penglq/QLog"
	"log"
	"runtime"
	"smallProGarm/conf"
	"strings"
)

type MLogParam struct {
	FilePath     string
	FileName     string
	FileSuffix   string //后缀
	FileMaxSize  int64
	FileMaxNSize int
	TimeZone     string
}

var mLogParam *MLogParam

type MyParam func(param *MLogParam) interface{}

func (mlp *MyParam) SetFilePath(fp string) MyParam {
	return func(mlp *MLogParam) interface{} {
		f := mlp.FilePath
		mlp.FilePath = fp
		return f
	}
}
func (mlp *MyParam) SetFileName(fn string) MyParam {
	return func(mlp *MLogParam) interface{} {
		f := mlp.FileName
		mlp.FileName = fn
		return f
	}

}
func (mlp *MyParam) SetFileSuffix(fx string) MyParam {
	return func(mlp *MLogParam) interface{} {
		f := mlp.FileSuffix
		mlp.FileSuffix = fx
		return f
	}

}
func (mlp *MyParam) SetMaxSize(fms int64) MyParam {
	return func(mlp *MLogParam) interface{} {
		f := mlp.FileMaxSize
		mlp.FileMaxSize = fms
		return f

	}

}
func (mlp *MyParam) SetFileMaxNSize(fmns int) MyParam {
	return func(mlp *MLogParam) interface{} {
		f := mlp.FileMaxNSize
		mlp.FileMaxNSize = fmns
		return f
	}
}
func (mlp *MyParam) SetTimeZone(tz string) MyParam {
	return func(mlp *MLogParam) interface{} {
		f := mlp.TimeZone
		mlp.TimeZone = tz
		return f
	}
}

var Mylog QLog.LoggerInterface

func (myp *MLogParam) MLogInit(options ...MyParam) error {
	q := &MLogParam{
		FilePath:     conf.LOGFILEPATH,
		FileName:     conf.LOGFILENAME,
		FileSuffix:   conf.LOGFILESUFFIX,
		FileMaxSize:  conf.LOGFILEMAXSIZE,
		FileMaxNSize: conf.LOGFILEMAXNSIZE,
		TimeZone:     conf.LOGTIMEZONE,
	}
	for _, option := range options {
		option(q)
	}
	mLogParam = q
	if mLogParam == nil {
		log.Fatalf("Mlog not init err %s", errors.New("日志没有初始化 - "))
	}
	l := QLog.GetLogger()
	l.SetConfig(QLog.INFO, mLogParam.TimeZone,
		QLog.WithFileOPT(mLogParam.FilePath, mLogParam.FileName, mLogParam.FileSuffix, mLogParam.FileMaxSize, mLogParam.FileMaxNSize),
		QLog.WithConsoleOPT(),
	)
	Mylog = l
	return nil
}

func MLog() QLog.LoggerInterface {
	//TODO:: i want to add prefix for log ,but the package is not written by myself
	//runtime.Caller用来获取当前goroutine函数调用栈的程序计数器及其相关信息
	funcName, _, _, ok := runtime.Caller(1)
	if ok {
		//获取包含给定PC地址的函数，无效则返回nil，调用name获取函数名称，fileline,源码文件名和行号，Entry获取对饮函数地址。
		fName := runtime.FuncForPC(funcName).Name()
		arrStr := strings.Split(fName, "/")
		Mylog.SetTextPrefix("method", arrStr[len(arrStr)-1])
	}
	return Mylog
}
