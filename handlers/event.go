package handlers

import (
	"net/http"
	"vespera-server/database"
	"vespera-server/models"
	"vespera-server/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	userID := c.Param("userID")

	events, err := services.GetEventsByUserID(database.DB, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve events"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"events": events})
}
func GetEventByIDHandler(c *gin.Context) {
	id := c.Param("id")

	event, err := services.GetEventByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"event": event})
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

	event.ID = uuid.New().String()

	c.JSON(http.StatusCreated, gin.H{"event": event})
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

	if err := services.DeleteEventByID(database.DB, id); err != nil {
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
		UserID   string `json:"userID" binding:"required"`
		EventoID string `json:"eventoID" binding:"required"`
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
		UserID   string `json:"userID" binding:"required"`
		EventoID string `json:"eventoID" binding:"required"`
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

	evento, err := services.GetEventByID(database.DB, eventoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch event with likes"})
		return
	}

	c.JSON(http.StatusOK, evento.EventLikes)
}
