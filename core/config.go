package core

import (
	"bufio"
	"log"
	"os"
	"sniping/config"

	"gopkg.in/yaml.v3"
)

func InitConfig(configPath string) config.Config {
    data, err := os.ReadFile(configPath)
    if err != nil {
        panic(err)
    }

    var raw config.Config
    err = yaml.Unmarshal(data, &raw)
    if err != nil {
        panic(err)
    }

    return raw
}

func LoadInputs(cfg config.Config) []string {
    file, err := os.Open(cfg.Input.Path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	var users []string
    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
		users = append(users, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return users
}

