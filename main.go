package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"

	"app/db"
)

type Response struct {
	ID interface{} `json:"_id"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()

	client := db.ConnectMongoDB()
	db.CreateCollection(client)
	collection := db.GetConnectionOfCollection(client)

	e.POST("/user", func(c echo.Context) error {
		jsonBody := echo.Map{}
		if err := c.Bind(&jsonBody); err != nil {
			errorResponse := ErrorResponse{Message: "Bad Request"}
			return c.JSON(http.StatusBadRequest, errorResponse)
		}
		log.Print(jsonBody)
		result, _ := collection.InsertOne(context.TODO(), jsonBody)
		response := Response{ID: result.InsertedID}
		return c.JSON(http.StatusOK, response)
	})

	e.GET("/user/:_id", func(c echo.Context) error {
		raw_id := c.Param("_id")
		log.Print(raw_id)
		_id, _ := primitive.ObjectIDFromHex(raw_id)
		result := echo.Map{}
		if err := collection.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&result); err != nil {
			errorResponse := ErrorResponse{Message: "Not Found"}
			return c.JSON(http.StatusNotFound, errorResponse)
		}
		return c.JSON(http.StatusOK, result)
	})

	e.DELETE("/user/:_id", func(c echo.Context) error {
		raw_id := c.Param("_id")
		log.Print(raw_id)
		_id, _ := primitive.ObjectIDFromHex(raw_id)
		result, _ := collection.DeleteOne(context.TODO(), bson.M{"_id": _id})
		return c.JSON(http.StatusOK, result)
	})

	defer client.Disconnect(context.TODO())

	e.Logger.Fatal(e.Start(":1334"))
}
