package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Handler holds all dependencies required for our API handlers (DI pattern).
type Handler struct {
	db *gorm.DB
}

func (h *Handler) getHandler(c echo.Context) error {
	var messages []Message
	if err := h.db.Find(&messages).Error; err != nil {
		// Return 500 on any DB error
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "Error",
			Message: "Could not find the messages",
		})
	}

	return c.JSON(http.StatusOK, &messages)
}

func (h *Handler) postHandler(c echo.Context) error {
	var message Message
	if err := c.Bind(&message); err != nil {
		// 400 for bad JSON/malformed request
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not add the message",
		})
	}

	// Basic validation
	if message.Text == "" {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not add empty message",
		})
	}

	if err := h.db.Create(&message).Error; err != nil {
		// 500 for DB creation failure
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "Error",
			Message: "Could not create the message",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message was successfully created",
	})
}

func (h *Handler) patchHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// 400 for non-integer ID
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Bad ID",
		})
	}

	// 1. Find record first to ensure it exists (for 404)
	var message Message
	if err := h.db.First(&message, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 404 if record doesn't exist
			return c.JSON(http.StatusNotFound, Response{
				Status:  "Error",
				Message: "Message was not found",
			})
		}
		// 500 for other DB errors
		log.Printf("DB First error: %v", err)
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "Error",
			Message: "Server error",
		})
	}

	// 2. Bind the new data
	var updatedMessage Message
	if err := c.Bind(&updatedMessage); err != nil {
		return c.JSON(http.StatusBadRequest, Response{ // 400 (Bad JSON)
			Status:  "Error",
			Message: "Invalid input",
		})
	}

	// 3. Validate new data
	if updatedMessage.Text == "" {
		return c.JSON(http.StatusBadRequest, Response{ // 400 (Bad Data)
			Status:  "Error",
			Message: "Text cannot be empty",
		})
	}

	// 4. Update the found record
	if err := h.db.Model(&message).Update("text", updatedMessage.Text).Error; err != nil {
		log.Printf("DB Update error: %v", err) // 500
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "Error",
			Message: "Could not update the message",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message was updated",
	})
}

func (h *Handler) deleteHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Bad ID",
		})
	}

	// 1. Find record to ensure it exists (for 404)
	var message Message
	if err := h.db.First(&message, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, Response{
				Status:  "Error",
				Message: "Message was not found",
			})
		}
		log.Printf("DB First error: %v", err)
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "Error",
			Message: "Server error",
		})
	}

	// 2. Delete the record
	if err := h.db.Delete(&Message{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "Error",
			Message: "Could not delete the message",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message deleted",
	})
}
