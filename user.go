package main

import (
	"fmt"
	"github.com/nginth/gobitica/models"
	"net/http"
	"errors"
	"encoding/json"
	"io/ioutil"
	"context"
)

func getProfile(ctx context.Context, id, apiToken string) (models.UserResponse, error) {
	var user models.UserResponse

	resp, err := getUserProfile(ctx, id, apiToken)
	if err != nil {
		return user, err
	}
	user, err = readUserBody(resp)
	if err != nil {
		return user, err
	}
	return user, nil
}

func getUserProfile(ctx context.Context, id, apiToken string) (*http.Response, error) {
	var client http.Client

	request, err := http.NewRequest("GET", "https://habitica.com/api/v3/user", nil)
    request.Header["x-api-user"] = []string{id}
    request.Header["x-api-key"] = []string{apiToken}
	request = request.WithContext(ctx)
	resp, err := client.Do(request)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func readUserBody(resp *http.Response) (models.UserResponse, error) {
	var userData models.UserResponse

	if resp.StatusCode == 200 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return userData, err
		}
		err2 := json.Unmarshal(bodyBytes, &userData)
		if err2 != nil {
			return userData, err2
		}
	} else {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return userData, err
		}
		return userData, errors.New(fmt.Sprintf("StatusCode %d, Body %s \n", resp.StatusCode, bodyBytes))
	}
	
	return userData, nil
}

func main() {
	id, token, err := login(context.TODO(), "user", "pass")
	if err != nil {
		fmt.Println(err)
		return
	}
	user, err := getProfile(context.TODO(), id, token)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user)
}