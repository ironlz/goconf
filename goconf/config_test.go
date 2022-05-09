package goconf

import (
	"testing"
)

func TestFromFile(t *testing.T) {
	properties, err := PropertiesFromFile("../configs/config.yaml")
	if err != nil {
		t.Fatalf("failed convert file to properties. %v", err)
	}
	assert(t, properties, "test.intp", 1)
	assert(t, properties, "test.stringp", "string")
	assert(t, properties, "test.intArr.0", 12)
	assert(t, properties, "test.intArr.1", 13)
	assert(t, properties, "test.stringArr.0", "str1")
	assert(t, properties, "test.stringArr.1", "str2")
	assert(t, properties, "test.complexArr.0", 2)
	assert(t, properties, "test.complexArr.1", "str")
	assert(t, properties, "test.complexArr.2.eleint", 1)
	assert(t, properties, "test.complexArr.3.elestr", "elestr")
}

func TestPropertiesFromDir(t *testing.T) {
	propertiesFromDir, err := PropertiesFromDir("../configs", true)
	if err != nil {
		t.Fatalf("failed convert file to properties. %v", err)
	}
	assertNotExist(t, propertiesFromDir, "test.intp")
	assertNotExist(t, propertiesFromDir, "test.stringp")
	assertNotExist(t, propertiesFromDir, "test.intArr.0")
	assertNotExist(t, propertiesFromDir, "test.intArr.1")
	assertNotExist(t, propertiesFromDir, "test.stringArr.0")
	assertNotExist(t, propertiesFromDir, "test.stringArr.1")
	assertNotExist(t, propertiesFromDir, "test.complexArr.0")
	assertNotExist(t, propertiesFromDir, "test.complexArr.1")
	assertNotExist(t, propertiesFromDir, "test.complexArr.2.eleint")
	assertNotExist(t, propertiesFromDir, "test.complexArr.3.elestr")

	assert(t, propertiesFromDir, "config.test.intp", 1)
	assert(t, propertiesFromDir, "config.test.stringp", "string")
	assert(t, propertiesFromDir, "config.test.intArr.0", 12)
	assert(t, propertiesFromDir, "config.test.intArr.1", 13)
	assert(t, propertiesFromDir, "config.test.stringArr.0", "str1")
	assert(t, propertiesFromDir, "config.test.stringArr.1", "str2")
	assert(t, propertiesFromDir, "config.test.complexArr.0", 2)
	assert(t, propertiesFromDir, "config.test.complexArr.1", "str")
	assert(t, propertiesFromDir, "config.test.complexArr.2.eleint", 1)
	assert(t, propertiesFromDir, "config.test.complexArr.3.elestr", "elestr")

	assert(t, propertiesFromDir, "config2.test.intp", 1)
	assert(t, propertiesFromDir, "config2.test.stringp", "string")
	assert(t, propertiesFromDir, "config2.test.intArr.0", 12)
	assert(t, propertiesFromDir, "config2.test.intArr.1", 13)
	assert(t, propertiesFromDir, "config2.test.stringArr.0", "str1")
	assert(t, propertiesFromDir, "config2.test.stringArr.1", "str2")
	assert(t, propertiesFromDir, "config2.test.complexArr.0", 2)
	assert(t, propertiesFromDir, "config2.test.complexArr.1", "str")
	assert(t, propertiesFromDir, "config2.test.complexArr.2.eleint", 1)
	assert(t, propertiesFromDir, "config2.test.complexArr.3.elestr", "elestr")
}

func assert(t *testing.T, c Properties, propertyName string, expected interface{}) {
	property, err := c.GetProperty(propertyName)
	if err != nil {
		t.Fatalf("Failed get property %s. %v", propertyName, err)
	}
	if property != expected {
		t.Fatalf("Property %s shoule be %s but %s", propertyName, expected, property)
	}
}

func assertNotExist(t *testing.T, c Properties, propertyName string) {
	property, err := c.GetProperty(propertyName)
	if err == nil {
		t.Fatalf("Property %s should not exist but %s", propertyName, property)
	}
}
