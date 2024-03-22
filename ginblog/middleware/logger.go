package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

//洋葱模型中间件

func Logger() gin.HandlerFunc {
	filePath := "log/log"
	LinkName := "latest_log.log"
	//三个参数，一个地址文件地址，第二个操作，os.O_RDWR 表示以读写方式打开文件，os.O_CREATE 表示如果文件不存在，则创建文件。
	//第三个参数是权限对于所有者可读、可写、可执行，对于组和其他用户可读、可执行，但不可写。
	src, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("err:", err)
	}
	logger := logrus.New()

	logger.Out = src

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	//日志分割
	logWriter, _ := retalog.New(
		filePath+"%Y%m%d.log",
		//保存的最长时间
		retalog.WithMaxAge(7*24*time.Hour),
		//每24小时分割一次
		retalog.WithRotationTime(24*time.Hour),
		//软链接
		retalog.WithLinkName(LinkName),
	)
	//写什么进去,这里可以根据自己的要求选择不通的信息，存放在不同的路径上
	//writeMap 是一个字典，用于指定不同级别的日志应该被写入到哪个文件中
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	//钩子：钩子"（Hook）通常指的是一种机制，允许开发者在某个特定事件发生时插入自定义代码或功能。
	//钩子允许你在其他代码执行的特定时刻进行拦截或者添加额外的逻辑，以便实现定制化的行为.
	logger.AddHook(Hook)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0)))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"Hostname":  hostName,
			"status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		}
		if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
