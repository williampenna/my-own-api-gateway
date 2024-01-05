package main

import (
	"fmt"
	readYaml "my-own-api-gateway/src/controllers"
)

func main() {
	var c (*readYaml.Service)
	c = c.GetConf()

	fmt.Println(c)
}
