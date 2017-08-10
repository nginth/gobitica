package main

import (
	"fmt"
	"net/http"
	"context"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"strconv"
	"errors"
)

type Login struct {
	Success bool `json:"success"`
	Data struct {
		Id string `json:"id"`
		ApiToken string `json:"apiToken"`
		NewUser bool `json:"newUser"`
	} `json:"data"`
}

func login(ctx context.Context, user, pass string) (string, string, error) {
	var client http.Client

	data := url.Values{"username": {user}, "password": {pass}}
	request, err := http.NewRequest(
		"POST",
		"https://habitica.com/api/v3/user/auth/local/login", 
		bytes.NewBufferString(data.Encode()))
    request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	request = request.WithContext(ctx)
	resp, err := client.Do(request)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	var loginData Login
	if resp.StatusCode == 200 {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			return "", "", err2
		}
		err3 := json.Unmarshal(bodyBytes, &loginData)
		if err3 != nil {
			return "", "", err3
		}
	} else {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return "", "", errors.New(fmt.Sprintf("StatusCode %d, Body %s \n", resp.StatusCode, bodyBytes))
	}
	return loginData.Data.Id, loginData.Data.ApiToken, err
}


func main() {
	id, token, err := login(context.TODO(), "user", "pass")
	if err != nil {
		fmt.Println("err: ", err.Error())
	} else {
		fmt.Println("id: ", id, "token: ", token)
	}
}