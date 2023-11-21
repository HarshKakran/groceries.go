package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/HarshKakran/groceries.go/models"
	"github.com/gin-gonic/gin"
)

func StoreInfo(c *gin.Context) {
	userId, err := strconv.Atoi(c.GetHeader(HEADER_USER_ID))
	if err != nil {
		log.Panicf("error while converting string to int. %v", err)
	}
	for _, user := range models.Users {
		if user.ID == userId && user.Store != nil {
			c.JSON(http.StatusFound, *user.Store)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "store not found",
	})
}

func CreateItem(c *gin.Context) {
	userId, err := strconv.Atoi(c.GetHeader(HEADER_USER_ID))
	if err != nil {
		log.Panicf("error while converting string to int. %v", err)
	}

	// extracting request body
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		log.Panicf("error while extracting item DATA_FILE. %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	for _, user := range models.Users {
		if user.ID == userId && user.Store != nil {
			item.ID = strconv.Itoa(len(user.Store.Items)) + "_" + user.Username
			user.Store.Items = append(user.Store.Items, item)
			c.JSON(http.StatusOK, gin.H{"message": "item added",
				"item": item})
			models.Save(USERS_DATA_FILE, models.Users)
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "store not found",
	})
}

func UpdateItem(c *gin.Context) {
	userId, err := strconv.Atoi(c.GetHeader(HEADER_USER_ID))
	if err != nil {
		log.Panicf("error while converting string to int. %v", err)
	}

	fmt.Println(c.Param("id"))
	itemIdx, _, err := models.GetIndexAndUsername(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}

	/*
		for _, user := range models.Users {
			if user.ID == userId && user.Store != nil {
				for index, item := range user.Store.Items {
					if index == itemId {
						if err := c.ShouldBindJSON(&item); err != nil {
							log.Panicf("error while extracting item USERS_DATA_FILE. %v", err)
							c.JSON(http.StatusBadRequest, gin.H{
								"message": err.Error(),
							})
							return
						}
						user.Store.Items[index] = item
						models.Save(USERS_DATA_FILE, models.Users)
						c.JSON(http.StatusOK, gin.H{
							"message": "item added",
							"store":   user.Store,
						})
						return
					}
				}
			}
		}
	*/
	var user *models.User
	for _, u := range models.Users {
		if u.ID == userId && u.Store != nil {
			user = &u
		}
	}
	for idx, item := range user.Store.Items {
		if idx == itemIdx {
			if err := c.ShouldBindJSON(&item); err != nil {
				log.Panicf("error while extracting item USERS_DATA_FILE. %v", err)
				c.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
				return
			}
			user.Store.Items[idx] = item
			models.Save(USERS_DATA_FILE, models.Users)
			c.JSON(http.StatusOK, gin.H{
				"message": "item added",
				"store":   user.Store,
			})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "item id not found",
	})

}

func DeleteItem(c *gin.Context) {
	userId, err := strconv.Atoi(c.GetHeader(HEADER_USER_ID))
	if err != nil {
		log.Panicf("error while converting string to int. %v", err)
	}

	itemIdx, _, err := models.GetIndexAndUsername(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}

	/*
		for _, user := range models.Users {
			if user.ID == userId && user.Store != nil {
				for index := range user.Store.Items {
					if index == itemIdx {
						user.Store.Items = append(user.Store.Items[:index], user.Store.Items[index+1:]...)
						models.Save(USERS_DATA_FILE, models.Users)
						c.JSON(http.StatusOK, gin.H{
							"message": "item deleted",
						})
						return
					}
				}
				c.JSON(http.StatusNotFound, gin.H{
					"message": "item not found",
				})
			}
		}
	*/
	var user *models.User
	for _, u := range models.Users {
		if u.ID == userId && u.Store != nil {
			user = &u
		}
	}
	for idx := range user.Store.Items {
		if idx == itemIdx {
			user.Store.Items = append(user.Store.Items[:idx], user.Store.Items[idx+1:]...)
			models.Save(USERS_DATA_FILE, models.Users)
			c.JSON(http.StatusOK, gin.H{
				"message": "item deleted",
			})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "store not found",
	})
}
