package common

import (
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Mysql struct {
	Url string `yaml:"url"`
}

type Server struct {
	Ip   string `yaml:"ip"`
	Port string `yaml:"port"`
}

type BaseInfo struct {
	Mysql  Mysql  `yaml:"mysql"`
	Server Server `yaml:"server"`
}

var (
	Info BaseInfo
)

func init() {
	path, err := os.Executable()
	if err != nil {
		panic(err.Error())
	}
	dir := filepath.Dir(path)
	yamlFile, err := ioutil.ReadFile(dir + "\\application.yml")
	if err != nil {
		panic(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &Info)
	if err != nil {
		panic(err.Error())
	}
}

var (
	Db  *gorm.DB
	err error
)
