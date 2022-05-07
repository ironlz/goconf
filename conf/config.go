package conf

import (
	"io/fs"
	"path/filepath"
)

// Config 代表来自某个路径的配置文件
type Config struct {
	content  map[string]SubConf // 文件名称 配置
	confPath string             // 初始化时指定的配置文件路径
}

// Use 声明使用哪一个配置文件
func (c *Config) Use(conf string) *SubConf {
	fileName, _ := spiltFileNameAndType(conf)
	subConf, exist := c.content[fileName]
	if exist {
		return &subConf
	} else {
		return &SubConf{}
	}
}

func ConstructConfig(path string) (*Config, error) {
	c := Config{
		content:  make(map[string]SubConf),
		confPath: path,
	}
	err := filepath.Walk(path, func(currentPath string, info fs.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		subConf, fn, err := ConstructSubConf(currentPath)
		if err != nil {
			return err
		}
		c.content[fn] = subConf
		return nil
	})
	return &c, err
}
