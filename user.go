package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

func createUser(c *gin.Context) {
	db := c.MustGet("db").(*pgxpool.Pool)

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID is not a string"})
		return
	}

	println("userid: " + userIDStr)

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		return
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return
	}
	token := tokenParts[1]

	user, err := authenticateUser(userIDStr, token, db)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var existingUserID string
	err = db.QueryRow(context.Background(), "SELECT id FROM users WHERE id = $1", user.ID).Scan(&existingUserID)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	} else if err != pgx.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = db.Exec(context.Background(), `
        INSERT INTO users (id, grade, email, image, interest_group)
        VALUES ($1, $2, $3, $4, $5)
    `, user.ID, user.Grade, user.Email, user.Image, user.InterestGroup)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}

func getUser(c *gin.Context) {
	db := c.MustGet("db").(*pgxpool.Pool)

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	authHeader := c.GetHeader("Authorization")
	token := strings.Split(authHeader, " ")[1]
	user, err := authenticateUser(userID.(string), token, db)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":             user.ID,
		"grade":          user.Grade,
		"email":          user.Email,
		"image":          user.Image,
		"interest_group": user.InterestGroup,
	})
}
