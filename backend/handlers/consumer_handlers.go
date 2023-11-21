package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/HarshKakran/groceries.go/models"
	"github.com/gin-gonic/gin"
)

func ConsumerPage(c *gin.Context) {
	err := models.Load(USERS_DATA_FILE, &models.Users)
	if err != nil {
		log.Panicf("error while refreshing users data. %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong...",
		})
	}

	var availableItems []models.Item

	for _, user := range models.Users {
		if user.Store == nil {
			continue
		}
		for _, item := range user.Store.Items {
			if item.Quantity > 0 {
				availableItems = append(availableItems, item)
			}
		}
	}

	c.JSON(http.StatusOK, availableItems)
}

func CreateOrder(c *gin.Context) {
	err := models.Load(USERS_DATA_FILE, &models.Users)
	if err != nil {
		log.Panicf("error while refreshing users data. %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong...",
		})
	}

	userId, err := strconv.Atoi(c.GetHeader(HEADER_USER_ID))
	if err != nil {
		log.Panicf("error while converting string to int. %v", err)
	}

	var orderItems models.Orders
	if err := c.ShouldBindJSON(&orderItems); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	// update item's quantity
	for _, orderItem := range orderItems {
		_, username, err := models.GetIndexAndUsername(orderItem.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
		// fmt.Println("USERNAME: ", username)

		// updateQuantities[username] = append(updateQuantities[username], itemIdx)
		for _, user := range models.Users {
			if user.Username == username {
				for i, item := range user.Store.Items {
					if item.ID == orderItem.ID {
						user.Store.Items[i].Quantity -= item.Quantity
						// fmt.Println(i.Quantity)
					}
				}
			}
		}
	}
	if err := models.Save(USERS_DATA_FILE, models.Users); err != nil {
		log.Panicf("error while saving user data. %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong... try again",
		})
		return
	}

	// creating the order history for the consumer
	for _, user := range models.Users {
		if user.ID == userId {
			orders := models.ConsumerOrdersMapping[userId]
			orders = append(orders, orderItems...)
			models.ConsumerOrdersMapping[userId] = orders
			if err := models.Save(ORDERS_DATA_FILE, &models.ConsumerOrdersMapping); err != nil {
				log.Panicf("error while saving order data. %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "something went wrong... try again",
				})
				return
			}
			c.JSON(http.StatusCreated, gin.H{
				"meesage": "order created",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "user not found",
	})
}
