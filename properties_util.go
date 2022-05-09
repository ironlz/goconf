package goconf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strconv"
)

var UnSupportConfigFileError = fmt.Errorf("unsupport config file")

func PropertiesFromFile(fileName string) (Properties, error) {
	properties, err := parseYamlFile(fileName)
	return properties, err
}

func PropertiesFromMap(properties map[string]interface{}) Properties {
	return extractSubConf(properties)
}

func PropertiesFromDir(path string, useFileNamePrefix bool) (Properties, error) {
	result := Properties{}
	err := filepath.Walk(path, func(currentPath string, info fs.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		propertiesFromFile, err := PropertiesFromFile(currentPath)
		if err != nil {
			return err
		}
		fileName, _ := spiltFileNameAndType(info.Name())
		for k, v := range propertiesFromFile {
			if useFileNamePrefix {
				result.AddProperty(fileName+"."+k, v)
			} else {
				result.AddProperty(k, v)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func parseYamlFile(fileName string) (Properties, error) {
	_, fileType := spiltFileNameAndType(fileName)
	if fileType == "yaml" {
		fileData, err := ioutil.ReadFile(fileName)
		if err != nil {
			return nil, err
		}
		tmp := map[string]interface{}{}
		err = yaml.Unmarshal(fileData, tmp)
		if err != nil {
			return nil, err
		}
		return extractSubConf(tmp), nil
	}
	return nil, UnSupportConfigFileError
}

func extractSubConf(tmp map[string]interface{}) Properties {
	result := Properties{}
	for currentKey, v := range tmp {
		kind := reflect.TypeOf(v).Kind()
		if kind == reflect.Map {
			parseMap(currentKey, v.(map[string]interface{}), result)
		} else if kind == reflect.Slice || kind == reflect.Array {
			parseSlice(currentKey, v.([]interface{}), result)
		} else {
			result[currentKey] = v
		}
	}
	return result
}

func parseSlice(prefix string, array []interface{}, result Properties) {
	result[prefix] = array
	for index, v := range array {
		currentKey := prefix + "." + strconv.Itoa(index)
		kind := reflect.TypeOf(v).Kind()
		if kind == reflect.Map {
			parseMap(currentKey, v.(map[string]interface{}), result)
		} else if kind == reflect.Slice || kind == reflect.Array {
			parseSlice(currentKey, v.([]interface{}), result)
		} else {
			result[currentKey] = v
		}
	}
}

func parseMap(prefix string, m map[string]interface{}, result Properties) {
	result[prefix] = m
	for k, v := range m {
		currentKey := prefix + "." + k
		kind := reflect.TypeOf(v).Kind()
		if kind == reflect.Map {
			parseMap(currentKey, v.(map[string]interface{}), result)
		} else if kind == reflect.Slice || kind == reflect.Array {
			parseSlice(currentKey, v.([]interface{}), result)
		} else {
			result[currentKey] = v
		}
	}
}
