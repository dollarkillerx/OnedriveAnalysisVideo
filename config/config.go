/**
 * @Author: DollarKiller
 * @Description: Config
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:37 2019-10-24
 */
package config

import (
	"github.com/dollarkillerx/easyutils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type base struct {
	App struct {
		Host        string `yaml:"host"`
		Debug       bool   `yaml:"debug"`
		OneDriveUrl string `yaml:"one_drive_url"`
	}
	Pgsql struct {
		Dsn     string        `yaml:"dsn"`
		MaxIdle int           `yaml:"max_idle"`
		MaxOpen int           `yaml:"max_open"`
		TimeOut time.Duration `yaml:"time_out"`
	}
}

var (
	MyConfig *base
)

func init() {
	MyConfig = &base{}

	b, i := easyutils.PathExists("./config.yml")
	if i != nil || b == false {
		// 如果配置文件不存在
		createConfig()
	}

	bytes, e := ioutil.ReadFile("./config.yml")
	if e != nil {
		panic(e.Error())
	}

	e = yaml.Unmarshal(bytes, MyConfig)
	if e != nil {
		panic(e.Error())
	}

	if MyConfig.App.Debug {
		log.Println(MyConfig)
	}
}

func createConfig() {
	err := ioutil.WriteFile("config.yml", []byte(conf), 00755)
	if err != nil {
		panic(err)
	}
	log.Fatalln("请填写配置文件!")
}

var conf = `
# 通用配置文件

app:
  host: "0.0.0.0:8082"
  debug: true
  one_drive_url: "https://vipmail-my.sharepoint.cn"  // OneDriveUrl (国内版本:https://vipmail-my.sharepoint.cn)

pgsql:
  dsn: "host= user= dbname= sslmode=disable password="
  max_idle: 10 # 最大闲置链接数
  max_open: 100 # 最大链接数
  time_out: 4 # 超时 Hour
`
