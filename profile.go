package main

import "os"

type Profile struct {
	keyID       string
	accountName string
	tritonURL   string
}

func NewProfile() Profile {
	return Profile{
		keyID:       os.Getenv("SDC_KEY_ID"),
		accountName: os.Getenv("SDC_ACCOUNT"),
		tritonURL:   os.Getenv("SDC_URL"),
	}
}
