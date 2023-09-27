package age

import "fmt"

const InvalidAgeMessage = "age is not valid"

func EnterAge(age int) error {
	if age <= 6 {
		return fmt.Errorf(InvalidAgeMessage)
	}

	if age >= 18 {
		return fmt.Errorf(InvalidAgeMessage)
	}

	return nil
}
