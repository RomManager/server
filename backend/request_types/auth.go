package request_types

import (
	"errors"
	"fmt"

	"github.com/vallezw/RomManager/backend/utils"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *LoginRequest) ValidateForm() error {
	if req.Email == "" {
		return errors.New("you need to give an email")
	}
	if req.Password == "" {
		return errors.New("you need to give a password")
	}
	return nil
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *RegisterRequest) ValidateForm() error {
	fmt.Println("request is: ")
	if req.Email == "" {
		return errors.New("you need to give an email")
	}
	if req.Password == "" {
		return errors.New("you need to give a password")
	}
	if !utils.IsValidMailAddress(req.Email) {
		return errors.New("given email is not a proper address")
	}
	return nil
}
