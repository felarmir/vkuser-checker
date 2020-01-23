package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/felarmir/vkuser-checker/apiservice"
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
		panic(err)
	}
	user, err := vkclient.UserByName(loadedConfig.SearchName, resp)
	if err != nil {
		fmt.Println(err)
	}

	credentials := loadedConfig.MakeDBConnection()
	db, err := credentials.PGConnect()
	if err != nil {
		fmt.Println("Info:", err)
		panic("Exit from app")
	}

	go func() {
		updateData(db, user, loadedConfig, resp)
	}()

	err = apiservice.StartServer(db)
	if err != nil {
		fmt.Println("Info:", err)
	}

}

func updateData(db *sql.DB, user *vkclient.UserItem, loadedConfig *config.Configuration, resp *vkclient.AuthResponse) {
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
