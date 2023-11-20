package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/HarshKakran/groceries.go/models"
	"github.com/gin-gonic/gin"
)

func StoreInfo(c *gin.Context) {
	userId, err := strconv.Atoi(c.GetHeader(HEADER_USER_ID))
	if err != nil {
		log.Fatalf("error while converting string to int. %v", err)
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
		log.Fatalf("error while converting string to int. %v", err)
	}

	// extracting request body
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		log.Panicf("error while extracting item data. %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	for _, user := range models.Users {
		if user.ID == userId && user.Store != nil {
			user.Store.Items = append(user.Store.Items, item)
			c.JSON(http.StatusOK, gin.H{"message": "item added"})
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
		log.Fatalf("error while converting string to int. %v", err)
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalf("error while converting string to int. %v", err)
	}

	for _, user := range models.Users {
		if user.ID == userId && user.Store != nil {
			for index, item := range user.Store.Items {
				if index == itemId {
					if err := c.ShouldBindJSON(&item); err != nil {
						log.Panicf("error while extracting item data. %v", err)
						c.JSON(http.StatusBadRequest, gin.H{
							"message": err.Error(),
						})
						return
					}
					user.Store.Items[index] = item
					c.JSON(http.StatusOK, gin.H{
						"message": "item added",
						"store":   user.Store,
					})
					return
				}
			}
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "item id not found",
	})

}

func DeleteItem(c *gin.Context) {
	userId, err := strconv.Atoi(c.GetHeader(HEADER_USER_ID))
	if err != nil {
		log.Fatalf("error while converting string to int. %v", err)
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalf("error while converting string to int. %v", err)
	}

	for _, user := range models.Users {
		if user.ID == userId && user.Store != nil {
			for index, _ := range user.Store.Items {
				if index == itemId {
					user.Store.Items = append(user.Store.Items[:index], user.Store.Items[index+1:]...)
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

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "store not found",
	})
}
