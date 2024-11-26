package handlers

import (
	"log"
	"net/http"
	"strconv"
	"vespera-server/database"
	"vespera-server/models"
	"vespera-server/services"

	"github.com/gin-gonic/gin"
)

func GetAllEventsHandler(c *gin.Context) {
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not initialized"})
		return
	}
	events, err := services.GetAllEvents(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve events"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"events": events})
}
func GetEventsByUserIDHandler(c *gin.Context) {


	userID, exists := c.GetQuery("userID")
	
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userID is required"})
		return
	}
	log.Printf("userID: %s", userID)
	events, err := services.GetEventsByUserID(database.DB, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve events"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": events, "message": "Events found successfully"})
}

func GetEventByIDHandler(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "invalid request payload"})
	}

	event, err := services.GetEventByID(database.DB, uint(idInt))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": event, "message": "Event found successfully"})
}
func CreateEventHandler(c *gin.Context) {
	var event models.Evento
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	event.UserID = c.Param("userID")
	if err := services.CreateEvent(database.DB, &event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}


	c.JSON(http.StatusCreated, gin.H{"message":"event successfully created", "data": event})
}

func UpdateEventHandler(c *gin.Context) {
	id := c.Param("id")

	var updatedEvent models.Evento

	if err := services.UpdateEvent(database.DB, id, updatedEvent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func DeleteEventHandler(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
	}

	if err := services.DeleteEventByID(database.DB, uint(idInt)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}

func DeleteEventsByUserIDHandler(c *gin.Context) {
	userID := c.Param("userID")

	if err := services.DeleteEventsByUserID(database.DB, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete events"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "All events deleted successfully"})
}

// AddLikeHandler handles the request to like an event
func AddLikeHandler(c *gin.Context) {
	var input struct {
		UserID   uint `json:"userID" binding:"required"`
		EventoID uint `json:"eventoID" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.AddLike(database.DB, input.UserID, input.EventoID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event liked successfully"})
}

// RemoveLikeHandler handles the request to remove a like from an event
func RemoveLikeHandler(c *gin.Context) {
	var input struct {
		UserID   uint `json:"userID" binding:"required"`
		EventoID uint `json:"eventoID" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.RemoveLike(database.DB, input.UserID, input.EventoID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove like"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Like removed successfully"})
}

// GetLikesHandler handles the request to get all likes for an event
func GetLikesHandler(c *gin.Context) {
	eventoID := c.Param("eventoID")

	eventoIDInt, err := strconv.Atoi(eventoID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request"})
	}

	evento, err := services.GetEventByID(database.DB, uint(eventoIDInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch event with likes"})
		return
	}

	c.JSON(http.StatusOK, evento.EventLikes)
}
