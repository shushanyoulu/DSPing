package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Version     string
	Name        string
	Ip          string
	Db          string
	Alertsound  string
	Thdchecksec string
	Thdoccnum   string
	Thdavgdelay string
	Thdloss     string
	Tline       string
	Tsymbolsize string
	Targets     []Target
}

type configInfo struct {
	EmailConfig emailConfig
	LogConfig   logConfig
}

type emailConfig struct {
	EmailServerHost string
	EmailServerPort int
	SendFromUser    string
	SenderPasswd    string
	SendToUser      string
	CarbonCopyUser  string
}
type logConfig struct {
	LogPath string
}

const configPath = "../conf/dsPing.toml"

var configToml configInfo

// Opening (or creating) config file in JSON format
func readConfig(filename string) Config {
	config := Config{}
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		logger.Fatal("Config File Not Found!")
	} else {
		err = json.NewDecoder(file).Decode(&config)
		if err != nil {
			logger.Fatal(err)
		}
	}
	return config
}
func logPath() string {
	if _, err := toml.DecodeFile(configPath, &configToml); err != nil {
		fmt.Println(err)
	}
	return configToml.LogConfig.LogPath
}
func sendServerHost() string {
	if s, err := toml.DecodeFile(configPath, &configToml); err != nil {
		fmt.Println(err)
		fmt.Println(s)
	}
	return configToml.EmailConfig.EmailServerHost
}

func sendServerPort() int {
	if s, err := toml.DecodeFile(configPath, &configToml); err != nil {
		fmt.Println(err)
		fmt.Println(s)
	}
	return configToml.EmailConfig.EmailServerPort
}

func sendFromUser() string {
	if s, err := toml.DecodeFile(configPath, &configToml); err != nil {
		fmt.Println(err)
		fmt.Println(s)
	}
	return configToml.EmailConfig.SendFromUser
}

func senderPasswd() string {
	if s, err := toml.DecodeFile(configPath, &configToml); err != nil {
		fmt.Println(err)
		fmt.Println(s)
	}
	return configToml.EmailConfig.SenderPasswd
}

func senderToUser() string {
	if s, err := toml.DecodeFile(configPath, &configToml); err != nil {
		fmt.Println(err)
		fmt.Println(s)
	}
	return configToml.EmailConfig.SendToUser
}

func carbonCopyUser() string {
	if s, err := toml.DecodeFile(configPath, &configToml); err != nil {
		fmt.Println(err)
		fmt.Println(s)
	}
	return configToml.EmailConfig.CarbonCopyUser
}

// func main() {
// 	fmt.Println(sendServerHost())
// 	fmt.Println(sendServerPort())
// 	fmt.Println(sendFromUser())
// 	fmt.Println(senderPasswd())
// 	fmt.Println(senderToUser())
// 	fmt.Println(carbonCopyUser())
// 	fmt.Println(logPath())
// }
