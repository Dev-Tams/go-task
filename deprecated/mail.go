package deprecated

import (
	"fmt"
	"unicode/utf8"
)

type EmailValidator struct {
	Length int
	Email  string
}

func getLength(email string) EmailValidator {

	e := EmailValidator{
		Length: 0,
		Email:  email,
	}
	return e
}

func countLength(e *EmailValidator) int {
	// count the length of the email string
	count := 0
	for range e.Email {
		count++
	}
	return count
}

func (e *EmailValidator) checkMail() {

	Length := utf8.RuneCountInString(e.Email)
	fmt.Println("Length of the email is", Length)

	if len(e.Email) <= 5 {
		fmt.Println("Email is invalid")
	} else if len(e.Email) >= 6 && len(e.Email) <= 255 {
		fmt.Println("Email is valid")
	} else {
		fmt.Println("Email is too long")
	}
}
