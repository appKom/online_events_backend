package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config := LoadConfig()

	dbpool, err := pgxpool.New(context.Background(), config.GetDSN())
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbpool.Close()

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("db", dbpool)
		c.Next()
	})

	router.POST("/users", createUser)
	router.GET("/users/:id", getUser)

	router.Run(":8080")
}
