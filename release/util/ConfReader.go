package util

import (
	"fmt"
	"github.com/ali-zohrevand/gf-recapcha/release/models"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func GetConf(path string) (c models.Conf, err error) {

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return

	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		fmt.Println("Unmarshal: ", err)
		return
	}

	return
}
