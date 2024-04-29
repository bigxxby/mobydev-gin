package auth

import (
	"database/sql"
	"log"
	"net/http"
	"net/mail"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// @Tags passwordChange
// POST_ChangePassword changes current users password
// @Summary Change user password
// @Description Changes the password of the authenticated user
// @Accept json
// @Produce json
// @Param data body routes.ChangePasswordRequest true "Password change data"
// @Security ApiKeyAuth
// @Success 200 {object} routes.DefaultMessageResponse "Password updated"
// @Failure 400 {object} routes.DefaultMessageResponse "Bad request"
// @Failure 401 {object} routes.DefaultMessageResponse "Unauthorized"
// @Failure 500 {object} routes.DefaultMessageResponse "Internal server error"
// @Router /api/change-password [post]
func (m *AuthRoute) POST_ChangePassword(c *gin.Context) {
	userId := c.GetInt("userId")
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return

	}
	data := struct {
		OldPassword     string `json:"oldPassword" binding:"required"`
		NewPassword     string `json:"newPassword" binding:"required"`
		ConfirmPassword string `json:"confirmPassword" binding:"required"`
	}{}

	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	err = m.DB.UserRepository.CheckUserPassoword(userId, data.OldPassword)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Old password is wrong",
		})
		return
	}
	if data.ConfirmPassword != data.NewPassword {
		c.JSON(400, gin.H{
			"message": "Passwords dont match",
		})
		return
	}
	err = utils.CheckPasswordIsValid(data.NewPassword)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "New password is not Valid",
		})
		return
	}
	newPasswordHash, err := utils.HashPassword(data.NewPassword)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	err = m.DB.UserRepository.UpdatePasswordById(userId, newPasswordHash)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Password updated",
	})
}

// @Tags passwordChange
// POST_SendCode sends verification code
// @Summary Send verification code
// @Description Sends a verification code to the provided email address
// @Accept json
// @Produce json
// @Param data body routes.SendCodeRequest true "Email data"
// @Success 200 {object} routes.DefaultMessageResponse "Code sent"
// @Failure 400 {object} routes.DefaultMessageResponse "Bad request"
// @Failure 404 {object} routes.DefaultMessageResponse "User not found"
// @Failure 500 {object} routes.DefaultMessageResponse "Internal server error"
// @Router /api/send-code [post]
func (m *AuthRoute) POST_SendCode(c *gin.Context) {
	data := struct {
		Email string `json:"email" binding:"required"`
	}{}

	err := c.BindJSON(&data)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	_, err = mail.ParseAddress(data.Email) //is Really email
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	// err = m.DB.UserRepository.CheckUserExistsByEmail(data.Email)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		c.JSON(404, gin.H{
	// 			"message": "User not found",
	// 		})
	// 		return
	// 	}
	// 	log.Println(err.Error())
	// 	c.JSON(500, gin.H{
	// 		"message": "Internal server error",
	// 	})
	// 	return
	// }
	code := utils.GenerateVerificationCode()
	err = m.DB.UserRepository.AddVerificationCode(data.Email, code)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	err = utils.SendCodeEmail(data.Email, code)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	log.Println(code)
	c.JSON(200, gin.H{
		"message": "Code sent to the email (check spam too)",
	})
}

// @Tags passwordChange
// POST_VerifyCode verifies verification code
// @Summary Verify verification code
// @Description Verifies the verification code for the provided email address and generates a temporary password
// @Accept json
// @Produce json
// @Param data body routes.VerifyCodeRequest true "Verification code data"
// @Success 200 {object} routes.DefaultMessageResponse "Temporary password sent"
// @Failure 400 {object} routes.DefaultMessageResponse "Bad request"
// @Failure 500 {object} routes.DefaultMessageResponse "Internal server error"
// @Router /api/verify-code [post]
func (m *AuthRoute) POST_VerifyCode(c *gin.Context) {

	data := struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}{}

	err := c.BindJSON(&data)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	_, err = mail.ParseAddress(data.Email) //is Really email
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	err = m.DB.UserRepository.ValidateCode(data.Email, data.Code)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(400, gin.H{
				"message": "Code is not valid",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	newTemporaryPassword := utils.GenerateTempPassword(14)
	hashPassword, err := utils.HashPassword(newTemporaryPassword)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	err = m.DB.UserRepository.UpdatePassword(data.Email, hashPassword)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	err = utils.SendTempPasswordEmail(data.Email, newTemporaryPassword)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Temporary password sent to the email",
	})
}
