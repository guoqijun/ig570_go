package conf

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	Cnf *Conf
	Env string
)

func DefaultInit() {
	initLogRus()
	CnfInit()
}

var Logger = logrus.New()

func initLogRus() (err error) { // 初始化log的函数
	Logger.Formatter = &logrus.JSONFormatter{}                                    // 设置为json格式的日志
	f, err := os.OpenFile("./zdc.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644) // 创建一个log日志文件
	if err != nil {
		return
	}
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999999999",
	}) // 设置日志格式为json格式
	Logger.Out = f // 设置log的默认文件输出
	//gin.SetMode(gin.ReleaseMode) // 线上模式，控制台不会打印信息
	//gin.DefaultWriter = log.Out  // gin框架自己记录的日志也会输出
	Logger.Level = logrus.InfoLevel // 设置日志级别
	return
}

func CnfInit() {
	cf := &Conf{
		Url570: "www.baidu.com",
	}

	files, _ := filepath.Glob("./env.*.yaml")
	dev := false
	prod := false

	for _, v := range files {
		switch v {
		case "env.dev.yaml":
			dev = true
		case "env.prod.yaml":
			prod = true
		default:
			continue
		}
	}

	var fileName string
	var env string

	if dev {
		fileName = "/env.dev.yaml"
		env = "dev"
	} else if prod {
		fileName = "/env.prod.yaml"
		env = "prod"
	} else {
		fileName = "default"
		env = "dev"
	}

	if fileName == "default" {
		Cnf = cf
		Env = env
		return
	}

	res, err := filepath.Abs(filepath.Dir("./main.go"))
	if err != nil {
		Logger.WithFields(logrus.Fields{
			"msg": "ERROR",
		}).Error(err.Error())
	}

	//读取yaml配置文件
	yamlFile, err := ioutil.ReadFile(res + fileName)
	if err != nil {
		Logger.WithFields(logrus.Fields{
			"msg": "ERROR",
		}).Error(err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &cf)
	if err != nil {
		Logger.WithFields(logrus.Fields{
			"msg": "ERROR",
		}).Error(err.Error())
	}

	Cnf = cf
	Env = env
	return
}
