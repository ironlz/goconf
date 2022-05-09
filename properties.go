package goconf

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// Properties 子配置，代表一个配置文件
type Properties map[string]interface{}

func (c Properties) String() string {
	datas, err := json.Marshal(c)
	if err != nil {
		return "{}"
	}
	return string(datas)
}

func (c Properties) Size() int {
	return len(c)
}

func (c Properties) AddProperty(propertyName string, value interface{}) {
	c[propertyName] = value
}

func (c Properties) GetProperty(propertyName string) (interface{}, error) {
	v, exist := c[propertyName]
	if exist {
		return v, nil
	}
	return v, fmt.Errorf("property %s not exist", propertyName)
}

func (c Properties) GetPropertyWithDefault(propertyName string, def interface{}) interface{} {
	property, err := c.GetProperty(propertyName)
	if err != nil {
		return def
	}
	return property
}

func (c Properties) AddStringProperty(propertyName string, value string) {
	c.AddProperty(propertyName, value)
}

func (c Properties) GetStringProperty(propertyName string) (string, error) {
	v, err := c.GetProperty(propertyName)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v), err
}

func (c Properties) GetStringPropertyWithDefault(propertyName, defaultVaule string) string {
	v, err := c.GetStringProperty(propertyName)
	if err != nil {
		return defaultVaule
	}
	return v
}

func (c Properties) AddIntProperty(propertyName string, value int) {
	c.AddProperty(propertyName, value)
}

func (c Properties) GetIntProperty(propertyName string) (int, error) {
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
		return strconv.Atoi(fmt.Sprintf("%v", v))
	}
}

func (c Properties) GetIntPropertyWithDefault(propertyName string, def int) int {
	property, err := c.GetIntProperty(propertyName)
	if err != nil {
		return def
	}
	return property
}

func (c Properties) AddArrayProperty(propertyName string, value []interface{}) {
	c.AddProperty(propertyName, value)
}

func (c Properties) GetArrayProperty(propertyName string) ([]interface{}, error) {
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

func (c Properties) GetArrayPropertyWithDefault(propertyName string, def []interface{}) []interface{} {
	property, err := c.GetArrayProperty(propertyName)
	if err != nil {
		return def
	}
	return property
}

func (c Properties) AddStringArrayProperty(propertyName string, value []string) {
	c.AddProperty(propertyName, value)
}

func (c Properties) GetStringArrayProperty(propertyName string) ([]string, error) {
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

func (c Properties) GetStringArrayPropertyWithDefault(propertyName string, def []string) []string {
	property, err := c.GetStringArrayProperty(propertyName)
	if err != nil {
		return def
	}
	return property
}

func (c Properties) AddIntArrayProperty(propertyName string, value []int) {
	c.AddProperty(propertyName, value)
}

func (c Properties) GetIntArrayProperty(propertyName string) ([]int, error) {
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

func (c Properties) GetIntArrayPropertyWithDefault(propertyName string, def []int) []int {
	property, err := c.GetIntArrayProperty(propertyName)
	if err != nil {
		return def
	}
	return property
}
