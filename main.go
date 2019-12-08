package main

import (
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gyozabu/himechat-cli/generator"
)

func responseBadRequest(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(400, gin.H{
		"error": msg,
	})
}

func handleReponse(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Max-Age", "86400")
	c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	targetName := c.DefaultQuery("name", "")
	emojiNumStr := c.DefaultQuery("e", "4")
	punctuationLabelStr := c.DefaultQuery("p", "0")

	emojiNum, err := strconv.Atoi(emojiNumStr)
	if err != nil {
		responseBadRequest(c, "It can only handle numbers.")
		return
	}

	punctuationLabel, err := strconv.Atoi(punctuationLabelStr)
	if err != nil {
		responseBadRequest(c, "It can only handle numbers.")
		return
	}
	if punctuationLabel > 3 || punctuationLabel < 0 {
		responseBadRequest(c, "Possible values for punctuation label are 0 to 3")
		return
	}

	config := generator.Config{
		TargetName: targetName,
		EmojiNum:   emojiNum,
		//		PunctiuationLebel: punctuationLabel,
	}

	result, err := generator.Start(config)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, gin.H{
		"message": result,
	})
}

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081", "https://himechat-gyoza.web.app"}
	r.Use(cors.New(config))
	r.GET("/", handleReponse)
	r.Run()
}
