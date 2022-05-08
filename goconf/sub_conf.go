package goconf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strconv"
)

// SubConf 子配置，代表一个配置文件
type SubConf map[string]interface{}

func (c *SubConf) GetProperty(propertyName string) (interface{}, error) {
	v, exist := (*c)[propertyName]
	if exist {
		return v, nil
	}
	return v, fmt.Errorf("property %s not exist", propertyName)
}

func (c *SubConf) GetPropertyWithDefault(propertyName string, def interface{}) interface{} {
	property, err := c.GetProperty(propertyName)
	if err != nil {
		return def
	}
	return property
}

func (c *SubConf) GetStringProperty(propertyName string) (string, error) {
	v, err := c.GetProperty(propertyName)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", v), err
}

func (c *SubConf) GetStringPropertyWithDefault(propertyName, defaultVaule string) string {
	v, err := c.GetStringProperty(propertyName)
	if err != nil {
		return defaultVaule
	}
	return v
}

func (c *SubConf) GetIntProperty(propertyName string) (int, error) {
	v, err := c.GetProperty(propertyName)
	if err != nil {
		return 0, err
	}
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int:
		return v.(int), nil
	case reflect.Int8:
		return int(v.(int8)), nil
	case reflect.Int16:
		return int(v.(int16)), nil
	case reflect.Int32:
		return int(v.(int32)), nil
	case reflect.Int64:
		return int(v.(int64)), nil
	case reflect.Float32:
		return int(v.(float32)), nil
	case reflect.Float64:
		return int(v.(float64)), nil
	case reflect.Uint:
		return int(v.(uint)), nil
	case reflect.Uint8:
		return int(v.(uint8)), nil
	case reflect.Uint16:
		return int(v.(uint16)), nil
	case reflect.Uint32:
		return int(v.(uint32)), nil
	case reflect.Uint64:
		return int(v.(uint64)), nil
	default:
		return strconv.Atoi(fmt.Sprintf("%s", v))
	}
}

func (c *SubConf) GetIntPropertyWithDefault(propertyName string, def int) int {
	property, err := c.GetIntProperty(propertyName)
	if err != nil {
		return def
	}
	return property
}

func (c *SubConf) GetArrayProperty(propertyName string) ([]interface{}, error) {
	property, err := c.GetProperty(propertyName)
	if err != nil {
		return nil, err
	}
	kind := reflect.TypeOf(property).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		return property.([]interface{}), err
	} else {
		return nil, fmt.Errorf("%s not a slice or array", propertyName)
	}
}

func (c *SubConf) GetArrayPropertyWithDefault(propertyName string, def []interface{}) []interface{} {
	property, err := c.GetArrayProperty(propertyName)
	if err != nil {
		return def
	}
	return property
}

func (c *SubConf) GetStringArrayProperty(propertyName string) ([]string, error) {
	property, err := c.GetArrayProperty(propertyName)
	if err != nil {
		return nil, err
	}
	size := len(property)
	result := make([]string, 0)
	for i := 0; i < size; i++ {
		stringProperty, err := c.GetStringProperty(propertyName + "." + strconv.Itoa(i))
		if err != nil {
			return nil, err
		}
		result = append(result, stringProperty)
	}
	return result, nil
}

func (c *SubConf) GetStringArrayPropertyWithDefault(propertyName string, def []string) []string {
	property, err := c.GetStringArrayProperty(propertyName)
	if err != nil {
		return def
	}
	return property
}

func (c *SubConf) GetIntArrayProperty(propertyName string) ([]int, error) {
	property, err := c.GetArrayProperty(propertyName)
	if err != nil {
		return nil, err
	}
	size := len(property)
	result := make([]int, 0)
	for i := 0; i < size; i++ {
		intProperty, err := c.GetIntProperty(propertyName + "." + strconv.Itoa(i))
		if err != nil {
			return nil, err
		}
		result = append(result, intProperty)
	}
	return result, err
}

func (c *SubConf) GetIntArrayPropertyWithDefault(propertyName string, def []int) []int {
	property, err := c.GetIntArrayProperty(propertyName)
	if err != nil {
		return def
	}
	return property
}

func ConstructSubConf(subConfPath string) (SubConf, string, error) {
	startIndex := 0
	for i, ch := range subConfPath {
		if ch == filepath.Separator {
			startIndex = i + 1
		}
	}
	fileName := subConfPath[startIndex:]
	subConf, fn, err := parseYamlFile(subConfPath, fileName)
	return subConf, fn, err
}

func parseYamlFile(path, fileName string) (SubConf, string, error) {
	fileName, fileType := spiltFileNameAndType(fileName)
	if fileType == "yaml" {
		fileData, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, fileName, err
		}
		tmp := map[string]interface{}{}
		err = yaml.Unmarshal(fileData, tmp)
		if err != nil {
			return nil, fileName, err
		}
		return extractSubConf(tmp), fileName, nil
	}
	return SubConf{}, fileName, nil
}

func extractSubConf(tmp map[string]interface{}) SubConf {
	result := SubConf{}
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

func parseSlice(prefix string, array []interface{}, result SubConf) {
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

func parseMap(prefix string, m map[string]interface{}, result SubConf) {
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

func spiltFileNameAndType(fileName string) (string, string) {
	lastDot := 0
	for i, ch := range fileName {
		if ch == '.' {
			lastDot = i
		}
	}
	if lastDot == 0 {
		return fileName, ""
	}
	return fileName[:lastDot], fileName[lastDot+1:]
}
