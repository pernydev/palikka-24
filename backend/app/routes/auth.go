package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(c *gin.Context) {
	// Discord oAuth flow. This request is made by the client and contains the code from Discord in the body.
	// The code is then exchanged for an access token§ and the user's information is fetched.

	bodyStr, _ := c.GetRawData()
	code := string(bodyStr)

	if code == "" {
		c.JSON(400, gin.H{"error": "No code provided"})
		return
	}

	fmt.Println(code)

	body := gin.H{
		"grant_type":   "authorization_code",
		"code":         code,
		"redirect_uri": os.Getenv("DISCORD_REDIRECT_URI"),
	}

	bodyString := []byte("")
	for key, value := range body {
		bodyString = append(bodyString, []byte(key+"="+value.(string)+"&")...)
	}
	bodyString = bodyString[:len(bodyString)-1]

	req, err := http.NewRequest("POST", "https://discord.com/api/v10/oauth2/token", bytes.NewBuffer(bodyString))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(os.Getenv("DISCORD_CLIENT_ID"), os.Getenv("DISCORD_CLIENT_SECRET"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to send request"})
		return
	}

	defer resp.Body.Close()

	var data gin.H
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to decode response"})
		return
	}

	fmt.Println(data)

	if data["access_token"] == nil {
		c.JSON(400, gin.H{"error": "Invalid code"})
		return
	}

	req, err = http.NewRequest("GET", "https://discord.com/api/v10/users/@me", nil)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("Authorization", "Bearer "+data["access_token"].(string))

	resp, err = client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to send request"})
		return
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to decode response"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": data["id"],
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to sign token"})
		return
	}

	c.JSON(200, gin.H{"token": tokenString})
}