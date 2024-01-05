package readYaml

import (
	"io/ioutil"
	"log"
	apiGateway "my-own-api-gateway/src/structs"

	"gopkg.in/yaml.v3"
)

type Service apiGateway.Service

func (c *Service) GetConf() *Service {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
