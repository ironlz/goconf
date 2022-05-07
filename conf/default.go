package conf

var defaultConf *Config

func InitDefaultConf(path string) error {
	def, err := ConstructConfig(path)
	if err != nil {
		return err
	}
	defaultConf = def
	return nil
}

func Use(subConfName string) *SubConf {
	return defaultConf.Use(subConfName)
}
