package main

import (
	"fmt"
	"net/http"
	"context"
	"net/url"
	"io/ioutil"
	"encoding/json"
)

type Login struct {
	Success bool `json:"success"`
	Data struct {
		Id string `json:"id"`
		ApiToken string `json:"apiToken"`
		NewUser bool `json:"newUser"`
	} `json:"data"`
}

func login(ctx context.Context, user, pass string) (string, string) {
	var client http.Client

	resp, err := client.PostForm("https://habitica.com/api/v3/user/auth/local/login", 
		url.Values{"username": {user}, "password": {pass}})
	if err != nil {
		fmt.Println("error")
	}
	defer resp.Body.Close()

	var loginData Login
	if resp.StatusCode == 200 {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println("error 2")
		}
		err3 := json.Unmarshal(bodyBytes, &loginData)
		if err3 != nil {
			fmt.Println("error 3")
		}
	}
	return loginData.Data.Id, loginData.Data.ApiToken 
}


func main() {
	id, token := login(context.TODO(), "user", "pass")
	fmt.Println("id: ", id, "token: ", token)
}