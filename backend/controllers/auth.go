package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vallezw/RomManager/backend/models"
	"github.com/vallezw/RomManager/backend/request_types"
	"github.com/vallezw/RomManager/backend/utils"
	"github.com/vallezw/RomManager/backend/utils/token"
)

func Login(c *gin.Context) {
	var req request_types.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.DoError(c, http.StatusBadRequest, err)
		return
	}
	if err := req.ValidateForm(); err != nil {
		utils.DoError(c, http.StatusBadRequest, err)
		return
	}

	u := models.User{}

	u.Email = req.Email
	u.Password = req.Password

	token, err := models.LoginCheck(u.Email, u.Password)

	if err != nil {
		utils.DoError(c, http.StatusBadRequest, errors.New("email or password is incorrect"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Register(c *gin.Context) {
	var req request_types.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.DoError(c, http.StatusBadRequest, err)
		return
	}
	if err := req.ValidateForm(); err != nil {
		utils.DoError(c, http.StatusBadRequest, err)
		return
	}

	u := models.User{}

	u.Email = req.Email
	u.Password = req.Password

	u.Prepare()

	u.BeforeSave()

	createdUser, err := u.SaveUser()

	if err != nil {
		fmt.Println(err)
		if err.Error() == "UNIQUE constraint failed: users.email" {
			utils.DoError(c, http.StatusBadRequest, errors.New("email is already in usage"))
			return
		}
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
