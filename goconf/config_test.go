package goconf

import (
	"testing"
)

func TestConstructConfig(t *testing.T) {
	config, err := ConstructConfig("../configs")
	if err != nil {
		t.Fatalf("Failed parse config, %v", err)
	}
	use := config.Use("configs")
	assertNotExist(t, use, "test")
	assert(t, use, "test.strKey", "string")
	assert(t, use, "test.intKey", "int")
	assertNotExist(t, use, "test.strArr")
	assert(t, use, "test.strArr.0", "str1")
	assert(t, use, "test.strArr.1", "str2")
	assertNotExist(t, use, "test.intArr")
	assert(t, use, "test.intArr.0", 1)
	assert(t, use, "test.intArr.1", 2)
}

func assert(t *testing.T, c *SubConf, propertyName string, expected interface{}) {
	property, err := c.GetProperty(propertyName)
	if err != nil {
		t.Fatalf("Failed get property %s. %v", propertyName, err)
	}
	if property != expected {
		t.Fatalf("Property %s shoule be %s but %s", propertyName, expected, property)
	}
}

func assertNotExist(t *testing.T, c *SubConf, propertyName string) {
	property, err := c.GetProperty(propertyName)
	if err == nil {
		t.Fatalf("Property %s should not exist but %s", propertyName, property)
	}
}
