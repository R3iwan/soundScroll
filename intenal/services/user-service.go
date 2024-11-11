package services

import (
	"log"

	"github.com/R3iwan/soundScroll/intenal/models"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func RegisterUserService(req models.RegisterUser) error {
	if err := validate.Struct(req); err != nil {
		log.Fatalf("Incorrect Register Info")
	}

	if req.Age <= 0 || req.Age > 150 {
		log.Fatalf("Incorrect Age")
	}

	if req.FirstName == "" || req.Email == "" || req.LastName == "" || req.Password == "" || req.Username == "" {
		log.Fatalf("Please fill all the blanks")
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return err
	}

	req = models.RegisterUser{
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  hashedPassword,
		Age:       req.Age,
		Email:     req.FirstName,
	}

	log.Println("User registered successfully!")
	return nil
}

// func LoginUserService(req models.LoginUser) error {
// 	if err := validate.Struct(req); err != nil {
// 		log.Fatalf("Incorrect login info")
// 	}
// }

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost+5)
	if err != nil {
		return " ", err
	}
	return string(bytes), nil
}

// func checkHashPassword(password, hashedPassword string) bool {
// 	bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// 	return true
// }
