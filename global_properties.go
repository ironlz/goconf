package goconf

import "fmt"

var gProperties Properties

func InitGlobalPropertiesFromFile(file string) error {
	fromFile, err := PropertiesFromFile(file)
	if err != nil {
		return err
	}
	return setGPropertiesOnce(fromFile)
}

func InitGlobalPropertiesFromDir(filePath string) error {
	propertiesFromDir, err := PropertiesFromDir(filePath, true)
	if err != nil {
		return err
	}
	return setGPropertiesOnce(propertiesFromDir)
}

func setGPropertiesOnce(p Properties) error {
	if gProperties == nil {
		gProperties = p
	} else {
		return fmt.Errorf("duplicate init")
	}
	return nil
}
