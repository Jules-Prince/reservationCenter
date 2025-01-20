package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"hexagone/reservationCenter/src/models"
	"hexagone/reservationCenter/src/utils"
)

var ctx = context.Background()

// Get all availabilities
func GetAvailabilities(c *gin.Context) {
	// Retrieve all availability IDs from the hash map
	availabilityIDs, err := utils.RedisClient.HGetAll(ctx, "availabilities").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch availability IDs"})
		return
	}

	// Prepare a list to hold detailed availability information
	var availabilities []map[string]string

	// Iterate over each availability ID
	for _, key := range availabilityIDs {
		// Fetch all fields for the current availability key
		data, err := utils.RedisClient.HGetAll(ctx, key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch availability data"})
			return
		}
		availabilities = append(availabilities, data)
	}

	// Respond with the collected availability information
	c.JSON(http.StatusOK, availabilities)
}

// Create a new availability
func CreateAvailability(c *gin.Context) {
	var newAvailability models.Availability
	if err := c.ShouldBindJSON(&newAvailability); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	newAvailability.IsBooked = false

	// Use the ID as the key for the Redis hash
	key := "availability:" + newAvailability.ID

	// Store in Redis
	err := utils.RedisClient.HSet(ctx, key, map[string]interface{}{
		"date":      newAvailability.Date,
		"startTime": newAvailability.StartTime,
		"endTime":   newAvailability.EndTime,
		"isBooked":  newAvailability.IsBooked,
	}).Err()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save availability"})
		return
	}

	// Add key to a collection of all availability IDs
	utils.RedisClient.HSet(ctx, "availabilities", newAvailability.ID, key)

	c.JSON(http.StatusCreated, newAvailability)
}
