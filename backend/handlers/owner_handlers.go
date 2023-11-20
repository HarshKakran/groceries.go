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
	for _, store := range models.Stores {
		if store.Owner.ID == userId {
			c.JSON(http.StatusFound, store)
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "store not found",
	})
}

func CreateItem(c *gin.Context) {}
func UpdateItem(c *gin.Context) {}
func DeleteItem(c *gin.Context) {}
