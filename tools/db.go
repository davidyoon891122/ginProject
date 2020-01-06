package tools

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type UserInfo struct {
	ID       string `yaml:"id"`
	PASSWORD string `yaml:"pwd"`
}

var yamlPath string = "tempDB.yaml"

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func openFile() []byte {
	data, err := ioutil.ReadFile(yamlPath)
	checkError(err)
	return data
}

func ReadYaml() UserInfo {
	data := openFile()
	var userInfo UserInfo
	err := yaml.Unmarshal(data, &userInfo)
	checkError(err)
	log.Printf("userID : %s\n", userInfo.ID)
	log.Printf("userPWD : %s\n", userInfo.PASSWORD)

	return userInfo
}
