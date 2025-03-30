package auth

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

// Login işlemi
func Login(client *resty.Client, loginURL string, csrfToken string) error {
	username := os.Getenv("username")
	password := os.Getenv("password")

	resp, err := client.R().
		Get(loginURL)

	if err != nil {
		return fmt.Errorf("error occurred during initial request: %v", err)
	}

	loginResp, err := client.R().
		SetHeader("Referer", loginURL).
		SetFormData(map[string]string{
			"username":            username,
			"password":            password,
			"csrfmiddlewaretoken": csrfToken,
		}).
		SetCookies(resp.Cookies()).
		Post(loginURL)

	if err != nil {
		return fmt.Errorf("error occurred during login: %v", err)
	}

	// Durum kodunu kontrol et
	if loginResp.StatusCode() == 200 {
		fmt.Println("Giriş başarılı")
	} else {
		return fmt.Errorf("giriş başarısız. Durum Kodu: %d", loginResp.StatusCode())
	}

	return nil
}
