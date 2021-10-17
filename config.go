package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Debug bool `yaml:"debug"`
	Port  int  `yaml:"port"`

	BlockType BlockTypeEnum `yaml:"blockType"`

	Limiter struct {
		IpLimit int `yaml:"ipLimit"`
		IpBurst int `yaml:"ipBurst"`
	} `yaml:"limiter"`

	CustomSearch struct {
		Data string `yaml:"data"`
	} `yaml:"customSearch"`

	CustomSubtitle struct {
		ApiUrl   string `yaml:"apiUrl"`
		TeamName string `yaml:"teamName"`
	} `yaml:"customSubtitle"`

	Cache struct {
		AccessKey  time.Duration `yaml:"accessKey"`
		User       time.Duration `yaml:"user"`
		PlayUrl    time.Duration `yaml:"playUrl"`
		THSeason   time.Duration `yaml:"thSeason"`
		THSubtitle time.Duration `yaml:"thSubtitle"`
	} `yaml:"cache"`

	Proxy struct {
		CN string `yaml:"cn"`
		HK string `yaml:"hk"`
		TW string `yaml:"tw"`
		TH string `yaml:"th"`
	} `yaml:"proxy"`

	Reverse struct {
		CN string `yaml:"cn"`
		HK string `yaml:"hk"`
		TW string `yaml:"tw"`
		TH string `yaml:"th"`
	} `yaml:"reverse"`

	ReverseSearch struct {
		CN string `yaml:"cn"`
		HK string `yaml:"hk"`
		TW string `yaml:"tw"`
		TH string `yaml:"th"`
	} `yaml:"reverseSearch"`

	Auth struct {
		CN bool `yaml:"cn"`
		HK bool `yaml:"hk"`
		TW bool `yaml:"tw"`
		TH bool `yaml:"th"`
	} `yaml:"auth"`

	PostgreSQL struct {
		Host         string `yaml:"host"`
		User         string `yaml:"user"`
		Password     string `yaml:"password"`
		PasswordFile string `yaml:"passwordFile"`
		DBName       string `yaml:"dbName"`
		Port         int    `yaml:"port"`
	} `yaml:"postgreSQL"`
}

func validateConfigPath(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}
	if fi.IsDir() {
		return fmt.Errorf("'%s' is a directory", path)
	}
	return nil
}

func parseFlags() (string, error) {
	var configPath string

	flag.StringVar(&configPath, "config", "./config.yml", "Path to config file")

	flag.Parse()

	if err := validateConfigPath(configPath); err != nil {
		return "", err
	}

	return configPath, nil
}

func initConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
