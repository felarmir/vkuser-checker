package vkclient

import "encoding/json"

type UserItem struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	ScreenName string `json:"screen_name"`
}

type ResoinseData struct {
	Count int        `json:"count"`
	Items []UserItem `json:"items"`
}

type UserSearchRequest struct {
	Response ResoinseData `json:"response"`
}

func UserByName(name string, user *AuthResponse) (*UserItem, error) {
	parameters := make(map[string]string)
	parameters["q"] = name
	parameters["count"] = "5"
	parameters["fields"] = "screen_name"
	parameters["v"] = "5.103"

	content, err := Request("users.search", parameters, user)
	if err != nil {
		return nil, err
	}
	var userResponse *UserSearchRequest
	if err := json.Unmarshal(content, &userResponse); err != nil {
		return nil, err
	}
	return &userResponse.Response.Items[0], nil
}
