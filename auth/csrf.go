package auth

import (
	"bytes"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
)

func GetCSRFToken(client *resty.Client, loginURL string) (string, error) {
	resp, err := client.R().
		SetHeader("Referer", loginURL).
		Get(loginURL)
	if err != nil {
		return "", fmt.Errorf("error occurred: %v", err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body()))
	if err != nil {
		return "", fmt.Errorf("error parsing HTML: %v", err)
	}

	csrfToken, exists := doc.Find("input[name='csrfmiddlewaretoken']").Attr("value")
	if !exists {
		return "", fmt.Errorf("CSRF token not found")
	}

	return csrfToken, nil
}
