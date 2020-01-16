package main

import (
	"fmt"
	"time"

	"github.com/felarmir/vkuser-checker/config"
	"github.com/felarmir/vkuser-checker/dbclient"
	"github.com/felarmir/vkuser-checker/vkclient"
)

func main() {
	loadedConfig, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	resp, err := vkclient.Auth(loadedConfig.Vklogin, loadedConfig.VKPassword)
	if err != nil {
		fmt.Println(err)
	}
	user, err := vkclient.UserByName("Denis Andreev", resp)
	if err != nil {
		fmt.Println(err)
	}

	credentials := loadedConfig.MakeDBConnection()
	db, err := credentials.PGConnect()
	if err != nil {
		fmt.Println("Info:", err)
		panic("Exit from app")
	}

	for {
		info, err := vkclient.GetUserInfo(user.ID, resp)
		if err != nil {
			fmt.Println(err)
		}
		err = dbclient.InsertRow(db, *info)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(info.Online)
		time.Sleep(time.Second * time.Duration(loadedConfig.Timeout))
	}

}
