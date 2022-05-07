package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
	"reflect"
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

func (c *SubConf) GetStringProperty(propertyName string) (string, error) {
	v, err := c.GetProperty(propertyName)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", v), err
}

func (c *SubConf) GetStringPropertyWithDefault(propertyName, defaultVaule string) string {
	v, err := c.GetProperty(propertyName)
	if err != nil {
		return defaultVaule
	}
	return fmt.Sprintf("%s", v)
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

	}
}

func (c *SubConf) GetIntPropertyWithDefault(propertyName string, defaultVaule int) int {
	v, err := c.GetProperty(propertyName)
	if err != nil {
		return defaultVaule
	}
	return fmt.Sprintf("%s", v)
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
	for key, value := range tmp {
		switch reflect.TypeOf(value).Kind() {
		case reflect.Map:
			parseMap(key, value.(map[string]interface{}), result)
			break
		default:
			result[key] = value
		}
	}
	return result
}

func parseMap(prefix string, m map[string]interface{}, result SubConf) {
	for k, v := range m {
		currentKey := prefix + "." + k
		switch reflect.TypeOf(v).Kind() {
		case reflect.Map:
			parseMap(prefix+"."+k, v.(map[string]interface{}), result)
			break
		default:
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
