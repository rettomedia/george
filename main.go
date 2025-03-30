package main

import (
	"log"

	"george/auth"
	"george/utils"
)

func main() {
	client := utils.NewClient()

	loginURL := "https://retto.social/login/"

	csrfToken, err := auth.GetCSRFToken(client, loginURL)
	if err != nil {
		log.Fatalf("Error getting CSRF token: %v", err)
	}

	err = auth.Login(client, loginURL, csrfToken)
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}
}
