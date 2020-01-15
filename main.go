package main

import (
	"fmt"
	"time"

	"github.com/felarmir/vkuser-checker/vkclient"
)

func main() {
	resp, err := vkclient.Auth("", "")
	if err != nil {
		fmt.Println(err)
	}
	user, err := vkclient.UserByName("Denis Andreev", resp)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user.ID)

	for {
		info, err := vkclient.GetUserInfo(user.ID, resp)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(info.Online)
		time.Sleep(time.Second * 10)
	}

}
