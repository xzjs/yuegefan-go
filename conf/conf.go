package conf

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Map struct {
	AK string `yaml:"ak"`
}

type Conf struct {
	Map Map
}

var conf *Conf

func GetConf() *Conf {
	if conf == nil {
		conf = &Conf{}
	}
	file, err := ioutil.ReadFile("./app.yml")
	if err != nil {
		log.Fatal("fail to read file:", err)
	}
	err = yaml.Unmarshal(file, conf)
	if err != nil {
		log.Fatal("fail to yaml unmarshal:", err)
	}
	return conf
}
