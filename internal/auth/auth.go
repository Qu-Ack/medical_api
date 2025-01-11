package auth

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

const (
	maxAge = 86400 * 30
	isProd = false
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(".env files isn't loaded properly")
	}
	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	store := sessions.NewCookieStore([]byte("tryandbruteforcethisbitch"))

	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(google.New(googleClientId, googleClientSecret, "http://localhost:8080/auth/google/callback", "email", "profile"))

}

func BeginGoogleAuth(c *gin.Context) {
	q := c.Request.URL.Query()
	q.Add("provider", "google")
	c.Request.URL.RawQuery = q.Encode()
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func OAuthCallback(c *gin.Context) {

	q := c.Request.URL.Query()
	q.Add("provider", "google")
	c.Request.URL.RawQuery = q.Encode()
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	session, _ := gothic.Store.Get(c.Request, "gothic_session")
	session.Values["user"] = user
	session.Save(c.Request, c.Writer)

	http.Redirect(c.Writer, c.Request, "http://localhost:3000/", http.StatusTemporaryRedirect)
}
