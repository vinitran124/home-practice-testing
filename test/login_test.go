package test

import (
	"log"
	"practiceTesting/login"
	"testing"
)

func TestLogin(t *testing.T) {
	validUsername := login.ValidUsername
	validPassword := login.ValidPassword
	invalidUsername := "vinitrannn"
	invalidPassword := "vinitranxautrai"

	if login.Login(invalidUsername, invalidPassword) == true {
		log.Println("fail in testing 1")
	}
	log.Println("pass in testing 1")

	if login.Login(validUsername, invalidPassword) == true {
		log.Println("fail in testing 2")
	}
	log.Println("pass in testing 2")

	if login.Login(invalidUsername, validPassword) == true {
		log.Println("fail in testing 3")
	}
	log.Println("pass in testing 3")

	if login.Login(validUsername, validPassword) == false {
		log.Println("fail in testing 4")
	}
	log.Println("pass in testing 4")
}
