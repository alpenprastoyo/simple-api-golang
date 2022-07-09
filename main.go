package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)

	router.GET("/surveys", surveysHandler)

	router.GET("/questions", questionsHandler)

	router.GET("/survey/:id/:name_survey", surveyHandler)

	router.GET("/query", queryHandler)

	router.POST("/store_survey", storeSurveyHandler)

	router.Run(":8888")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}

func questionsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":       "1",
		"question": "Nama anda",
	})
}

func surveysHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":          "1",
		"name_survey": "Rekrutmen UNS",
	})
}

func surveyHandler(c *gin.Context) {
	id := c.Param("id")
	name_survey := c.Param("name_survey")

	c.JSON(http.StatusOK, gin.H{
		"id":          id,
		"name_survey": name_survey,
	})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
	})
}

func storeSurveyHandler(c *gin.Context){
	
}
