package config

import (
	"encoding/json"
	"github.com/felarmir/vkuser-checker/dbclient"
	"os"
)

type Configuration struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	User       string `json:"user"`
	Password   string `json:"password"`
	DBname     string `json:"dbname"`
	Vklogin    string `json:"vklogin"`
	VKPassword string `json:"vkpssword"`
	SearchName string `json:"vkSearchName"`
	Timeout    int    `json:"timeout"`
}

func LoadConfig() (*Configuration, error) {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := &Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		return nil, err
	}

	return configuration, nil
}

func (self *Configuration) MakeDBConnection() dbclient.DBConnection {
	connection := dbclient.DBConnection{}
	connection.Host = self.Host
	connection.Port = self.Port
	connection.User = self.User
	connection.Password = self.Password
	connection.DBname = self.DBname
	return connection
}
