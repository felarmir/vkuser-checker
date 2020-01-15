package main

import (
	"fmt"

	"github.com/felarmir/vkuser-checker/vkclient"
)

func main() {
	resp, err := vkclient.Auth("", "")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.AccessToken)

}
