package test

import (
	"errors"
	"fmt"
	"log"
	"practiceTesting/age"
	"testing"
)

func TestValidAge(t *testing.T) {
	ageTesting1 := 3
	ageTesting2 := 10
	ageTesting3 := 20

	invalidError := fmt.Errorf(age.InvalidAgeMessage)
	if errors.Is(invalidError, age.EnterAge(ageTesting1)) {
		log.Println("fail in testing 1")
	}
	log.Println("pass in testing 1")

	if age.EnterAge(ageTesting2) != nil {
		log.Println("fail in testing 2")
	}
	log.Println("pass in testing 2")

	if errors.Is(invalidError, age.EnterAge(ageTesting3)) {
		log.Println("fail in testing 3")
	}
	log.Println("pass in testing 3")
}
