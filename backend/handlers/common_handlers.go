package handlers

import (
	"log"
	"net/http"

	"github.com/HarshKakran/groceries.go/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signin(c *gin.Context) {
	var credentials models.User
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
	}

	// hashedPassword, err := hashPassword(credentials.Password)
	// if err != nil {
	// 	log.Panicf("error while hashing the password %v", err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "intetnal server error",
	// 	})
	// }

	var user *models.User

	for _, u := range models.Users {
		if credentials.Username == u.Username {
			user = &u
			break
		}
	}

	// if user != nil {
	// 	fmt.Println(user, checkPasswordHash(credentials.Password, user.Password))
	// }

	if user != nil && checkPasswordHash(credentials.Password, user.Password) {
		c.JSON(http.StatusFound, gin.H{
			"message":      "user logged in",
			HEADER_USER_ID: user.ID,
		})
		return
	} else if user == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "username not found",
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "invalid credentials",
	})
}

func Signup(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
	}

	for _, u := range models.Users {
		if u.Username == newUser.Username {
			c.JSON(http.StatusConflict, gin.H{
				"message": "username already exist",
			})
			return
		}
	}

	unhashedPassword := newUser.Password
	hashedPassword, err := hashPassword(unhashedPassword)
	if err != nil {
		log.Panicf("error hashing password %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "try again later",
		})
	}
	newUser.Password = hashedPassword
	newUser.ID = len(models.Users)

	models.Users = append(models.Users, newUser)
	if err = models.Save(USERS_DATA_FILE, models.Users); err != nil {
		log.Panicf("error while saving. %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
	})
}

func Signout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully logged out",
	})
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
