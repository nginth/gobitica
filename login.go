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
	"github.com/nginth/gobitica/models"
)

func login(ctx context.Context, user, pass string) (string, string, error) {
	resp, err := postToLogin(ctx, user, pass)
	if err != nil {
		return "", "", err
	}
	loginData, err := readLoginBody(resp)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	return loginData.Data.Id, loginData.Data.ApiToken, err
}

func postToLogin(ctx context.Context, user, pass string) (*http.Response, error) {
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
		return resp, err
	}

	return resp, nil
}

func readLoginBody(resp *http.Response) (models.LoginResponse, error) {
	var loginData models.LoginResponse

	if resp.StatusCode == 200 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return loginData, err
		}
		err2 := json.Unmarshal(bodyBytes, &loginData)
		if err2 != nil {
			return loginData, err2
		}
	} else {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return loginData, err
		}
		return loginData, errors.New(fmt.Sprintf("StatusCode %d, Body %s \n", resp.StatusCode, bodyBytes))
	}
	
	return loginData, nil
}


func main() {
	id, token, err := login(context.TODO(), "user", "pass")
	if err != nil {
		fmt.Println("err: ", err.Error())
	} else {
		fmt.Println("id: ", id, "token: ", token)
	}
}