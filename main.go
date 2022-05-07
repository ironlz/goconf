package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"reflect"
)

func main() {
	file, err := os.Open("./configs/config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, _ := ioutil.ReadAll(file)
	dat := make(map[string]interface{}, 0)
	err = yaml.Unmarshal(bytes, dat)
	for k, v := range dat {
		ktype := reflect.TypeOf(k).Kind()
		vType := reflect.TypeOf(v)
		fmt.Println(ktype, vType)
	}
}
