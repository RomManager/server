package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vallezw/RomManager/backend/models"
	"github.com/vallezw/RomManager/backend/utils"
	"github.com/vallezw/RomManager/backend/utils/token"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.DoError(c, http.StatusBadRequest, err)
		return
	}

	u := models.User{}

	u.Email = input.Email
	u.Password = input.Password

	token, err := models.LoginCheck(u.Email, u.Password)

	if err != nil {
		utils.DoError(c, http.StatusBadRequest, errors.New("email or password is incorrect"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {

	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.DoError(c, http.StatusBadRequest, errors.New("no email or password given"))
		return
	}

	u := models.User{}

	u.Email = input.Email
	u.Password = input.Password

	u.Prepare()

	u.BeforeSave()

	createdUser, err := u.SaveUser()

	if err != nil {
		utils.DoError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, createdUser)
}

func CurrentUser(c *gin.Context) {
	uid, err := token.ExtractTokenID(c)

	if err != nil {
		utils.DoError(c, http.StatusBadRequest, err)
		return
	}

	u, err := models.GetUserByID(uid)

	if err != nil {
		utils.DoError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
