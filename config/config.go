package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

const (
	configPath string = "./config/"
	appEnv     string = "APPENV"

	envProduction  string = "production"
	envStaging     string = "staging"
	envDevelopment string = "development"
	envTest        string = "test"
)

var conf *Config
var once sync.Once

//Get initializes config (if not yet initialized) then returns Config
func Get() Config {
	Init()
	return *conf
}

//Init initialized config by reading config json
func Init() {
	var err error
	once.Do(func() {
		conf, err = readConfig()
	})
	if err != nil {
		panic(err)
	}
}

func GetEnv() string {
	env := os.Getenv(appEnv)

	switch env {
	case envProduction, envStaging, envDevelopment, envTest:
	default:
		env = envDevelopment
	}

	return env
}

//readConfig read config json
func readConfig() (*Config, error) {
	env := GetEnv()

	path := fmt.Sprintf("%s%s.json", configPath, env)
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	result := new(Config)
	err = json.Unmarshal(c, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
