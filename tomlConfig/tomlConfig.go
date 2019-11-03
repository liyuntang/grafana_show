package tomlConfig

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"sync"
)

var (
	conf *Config
	once sync.Once
)

// 解析配置
func TomlCofig(configFile string) *Config {
	// 判断配置文件是否存在，如果不存在则报错，如果存在则返回文件描述句柄
	file := statFile(configFile)

	// 解析
	once.Do(func() {
		_, err3 := toml.DecodeReader(file, &conf)
		if err3 != nil {
			fmt.Println("toml configFile is bad, err3 is", err3)
			os.Exit(1)
		}
	})
	return conf
}

func statFile(file string) *os.File {
	_, err := os.Stat(file)
	if err != nil {
		fmt.Println("get file stat is bad, err is", err)
		os.Exit(1)
	}
	f, err1 := os.Open(file)
	if err1 != nil {
		fmt.Println("open file is bad, err1 is", err1)
		os.Exit(1)
	}
	return f
}