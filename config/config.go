package config

import (
	"encoding/json"
	"os"
	"log"
)

// Config for client
type Config struct {
	FollowNameList []FollowEntity `json:"followNameList"`
	RefreshInterval int `json:"refreshInterval"`
	FastRefreshInterval int `json:"fastRefreshInterval"`
	FastRefreshCount int `json:"fastRefreshCount"`
	FixedRefreshTimeList []string `json:"fixedRefreshTimeList"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	UserList []User `json:"userList"`
	LimitLikeCount int `json:"limitLikeCount"`
	LoginRefreshTime int `json:"loginRefreshTime"`
	NeedRefreshInterval bool `json:"needRefreshInterval"`
}

type FollowEntity struct {
	Name string `json:"name"`
	LimitLikeCount int `json:"limitLikeCount"`
}

type User struct {
	Phone string `json:"phone"`
	Password string `json:"password"`
}

func ParseJSONConfig(config *Config, path string) error {
	file, err := os.Open(path) // For read access.
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(config)
}

func GetConfig() *Config {
	var configPath string
	for index,cmd := range os.Args {
		//log.Print("cmd:",cmd)
		if (cmd == "-c") {
			configPath = os.Args[index + 1]
		}
	}
	log.Print("configPath:",configPath)
	if (configPath == "") {
		configPath = "config/config.json"
	}

	configEntity := Config{}
	ParseJSONConfig(&configEntity, configPath)
	return &configEntity
}
