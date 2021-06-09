package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost string
	DBUser string
	DBName string
	DBPwd  string
}

func ReadConfig() (conf *Config, err error) {
	conf = new(Config)
	v := viper.New()
	//设置读取的文件路径
	v.AddConfigPath("conf")
	v.AddConfigPath("./")
	v.SetConfigType("yaml")
	v.SetConfigName("conf")
	err = v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	// DB
	if v.IsSet("DBConfig") {
		conf.DBHost = v.Get("DBConfig.Host").(string)
		conf.DBName = v.Get("DBConfig.Name").(string)
		conf.DBUser = v.Get("DBConfig.User").(string)
		conf.DBPwd = v.Get("DBConfig.Password").(string)
	}
	return conf, nil
}
