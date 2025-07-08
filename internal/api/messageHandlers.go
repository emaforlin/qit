package handlers

import (
	"net/http"

	"github.com/emaforlin/qit/pkg/message"
	"github.com/emaforlin/qit/pkg/validation"
	"github.com/gin-gonic/gin"
)

func PostMessage(c *gin.Context) {
	var reqBody message.CreateMessageDto
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format", "details": err.Error()})
		return
	}

	if err := validation.ValidateStruct(reqBody); err != nil {
		validationErrors := validation.GetValidationErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"details": validationErrors,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Message created successfully",
		"data":    reqBody,
	})
}
