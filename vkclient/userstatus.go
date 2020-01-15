package vkclient

import (
	"encoding/json"
	"strconv"
)

type User struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	IsClosed     bool   `json:"is_closed"`
	AccessClosed bool   `json:"can_access_closed"`
	Online       int    `json:"online"`
}

type UserInfoResponse struct {
	Response []User `json:"response"`
}

/*
{
"response": [{
"id": 33629963,
"first_name": "Denis",
"last_name": "Andreev",
"is_closed": false,
"can_access_closed": true,
"online": 1
}]
}
*/

func GetUserInfo(userID int, user *AuthResponse) (*User, error) {
	parameters := make(map[string]string)
	parameters["user_ids"] = strconv.Itoa(userID)
	parameters["fields"] = "onlinde"
	parameters["v"] = "5.103"

	content, err := Request("users.get", parameters, user)
	if err != nil {
		return nil, err
	}
	var userResponse *UserInfoResponse
	if err := json.Unmarshal(content, &userResponse); err != nil {
		return nil, err
	}
	return &userResponse.Response[0], nil
}
