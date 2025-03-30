package utils

import "github.com/go-resty/resty/v2"

func NewClient() *resty.Client {
	return resty.New()
}
