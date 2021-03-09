package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"path/filepath"
	"sync"
	"time"
)

var (
	cfg  *tomlConfig
	once sync.Once
)

func test1() {
	var config = Config()
	fmt.Println(config)
}

func test2() {
	var config = Config()
	fmt.Println(config)
}

func main() {
	test1()
	test2()
}

func Config() *tomlConfig {

	//var once sync.Once

	once.Do(func() {
		filePath, err := filepath.Abs("./Library/Open_Source_Library/Config/toml/test.toml")
		if err != nil {
			panic(err)
		}
		fmt.Printf("parse toml file once. filePath: %s\n", filePath)
		if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
			panic(err)
		}
	})
	return cfg
}

type tomlConfig struct {
	Title   string
	Owner   ownerInfo
	DB      database `toml:"database"`
	Servers map[string]server
	Clients clients
}

type ownerInfo struct {
	Name string
	Org  string `toml:"organization"`
	Bio  string
	DOB  time.Time
}

type database struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type server struct {
	IP string
	DC string
}

type clients struct {
	Data  [][]interface{} `toml:"interface{}接收可以是任意类型"`
	Hosts []string
}
