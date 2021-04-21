package tools

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type LoginInfo struct {
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

func ReadYaml() LoginInfo {
	data := openFile()
	var loginInfo LoginInfo
	err := yaml.Unmarshal(data, &loginInfo)
	checkError(err)
	log.Printf("loginID : %s\n", loginInfo.ID)
	log.Printf("loginPWD : %s\n", loginInfo.PASSWORD)

	return loginInfo
}
