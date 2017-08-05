package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/adsense/v1.4"
)

var (
	port = 8080
)

func main() {
	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectUri := os.Getenv("GOOGLE_REDIRECT_URI")

	if len(clientId) == 0 {
		log.Fatal("GOOGLE_CLIENT_ID is empty.")
	}

	if len(clientSecret) == 0 {
		log.Fatal("GOOGLE_CLIENT_SECRET is empty.")
	}

	if len(redirectUri) == 0 {
		log.Fatal("GOOGLE_REDIRECT_URI is empty.")
	}

	oauth2Conf := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  redirectUri,
		Scopes: []string{
			"https://www.googleapis.com/auth/adsense.readonly",
		},
	}

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		url := oauth2Conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
		ctx.Redirect(http.StatusMovedPermanently, url)
	})

	r.GET("/auth", func(ctx *gin.Context) {
		code := ctx.Query("code")

		tok, err := oauth2Conf.Exchange(oauth2.NoContext, code)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		client := oauth2Conf.Client(oauth2.NoContext, tok)
		service, err := adsense.New(client)

		call := service.Accounts.List()
		resp, err := call.Do()
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, resp)
	})

	r.Run(fmt.Sprintf(":%d", port))
}
