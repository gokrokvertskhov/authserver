package main

import (
  "fmt"
  "github.com/BurntSushi/toml"
)

type TomlConfig struct {
	Default defaultInfo
	Log logInfo
	Secure secureInfo
}

type defaultInfo struct {
	Bind string
	Auth bool
}

type logInfo struct {
	File string
	Level string
	Console bool
}

type authInfo struct {
	Client_id string
	Client_key string
}

type secureInfo struct {
	PrivateKey string
	PublicKey string
}

var Conf TomlConfig

func LoadConfig(file string) {
	if _,err := toml.DecodeFile(file, &Conf); err != nil {
		fmt.Print(err)
		//PanicOnError(err)
	}
}