package models

type LoginResponse struct {
	Success bool `json:"success"`
	Data struct {
		Id string `json:"id"`
		ApiToken string `json:"apiToken"`
		NewUser bool `json:"newUser"`
	} `json:"data"`
}