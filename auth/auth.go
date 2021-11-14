package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/NickDubelman/pickup-list/db"
	"github.com/NickDubelman/pickup-list/db/user"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	// PathLogin is the path to handle logins
	PathLogin = "/auth/login"

	// PathCallback is the path to handle the callback from OAuth backend (Google)
	PathCallback = "/auth/google/callback"

	// PathError is redirected to when the user has an auth error
	PathError = "/auth/error"

	codeRedirect = http.StatusFound
)

// GoogleUserInfo represents a response from the Google userinfo API
type GoogleUserInfo struct {
	Email string `json:"email"`
}

// GoogleAuthFromConfig returns handlers that can be used for OAuth with Google
func GoogleAuthFromConfig() gin.HandlerFunc {
	config, err := getGoogleAuthConfig("oauthConfig.json")
	if err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			switch c.Request.URL.Path {
			case PathLogin:
				// Redirect the user to Google to authenticate
				http.Redirect(
					c.Writer,
					c.Request,
					config.AuthCodeURL(
						extractPath(c.Request.URL.Query().Get("next")),
						oauth2.SetAuthURLParam("prompt", "login"),
					),
					codeRedirect,
				)

			case PathCallback:
				// User succesfully authenticated with Google
				handleOAuth2Callback(config, c)

			case PathError:
				c.String(http.StatusInternalServerError, "Error logging in")

			}
		}
	}
}

// handleOAuth2Callback will be executed after the user authenticates with Google and
// consents to the scopes our app requires. We will retrieve their name, email, and
// picture from the Google userinfo endpoint. Lastly, we generate an access token and
// a refresh token for the user (both are JWTs)
func handleOAuth2Callback(cfg *oauth2.Config, ginCtx *gin.Context) {
	handleErr := func(err error) {
		log.Println(err)
		ginCtx.Redirect(http.StatusFound, PathError)
	}

	ctx := ginCtx.Request.Context()

	code := ginCtx.Request.URL.Query().Get("code")
	t, err := cfg.Exchange(ctx, code)
	if err != nil {
		handleErr(err)
		return
	}

	client := cfg.Client(ctx, t)
	userinfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		handleErr(err)
		return
	}
	defer userinfo.Body.Close()

	data, err := ioutil.ReadAll(userinfo.Body)
	if err != nil {
		handleErr(err)
		return
	}

	userInfo := GoogleUserInfo{}
	if err := json.Unmarshal(data, &userInfo); err != nil {
		handleErr(err)
		return
	}

	dbClient := db.FromContext(ctx)
	if dbClient == nil {
		handleErr(fmt.Errorf("could not retrieve db client from context"))
		return
	}

	user, err := dbClient.User.Query().Where(user.Email(userInfo.Email)).Only(ctx)
	if err != nil {
		if _, ok := err.(*db.NotFoundError); ok {
			// Create a user if one doesnt already exist for the given email
			user, err = dbClient.User.Create().
				SetRealName("").
				SetEmail(userInfo.Email).
				Save(ctx)
		}

		if err != nil {
			handleErr(err)
			return
		}
	}

	// Generate an access token
	accessToken, err := createAccessToken(ctx, user.ID, userInfo)
	if err != nil {
		handleErr(err)
		return
	}

	// Generate a refresh token
	refreshToken, err := createRefreshToken(user.ID)
	if err != nil {
		handleErr(err)
		return
	}

	state := ginCtx.Request.URL.Query().Get("state")

	nextURL, err := url.Parse("http://localhost:3000/login-callback")
	if err != nil {
		handleErr(err)
		return
	}

	// We need to add the user's tokens to the redirect URL so that the frontend can
	// pluck it from the URL and store the tokens in localStorage
	query := nextURL.Query()
	query.Add("accessToken", accessToken)
	query.Add("refreshToken", refreshToken)
	query.Add("state", state)
	nextURL.RawQuery = query.Encode()

	http.Redirect(ginCtx.Writer, ginCtx.Request, nextURL.String(), codeRedirect)
}

// getGoogleAuthConfig attempts to read from a given oauthConfigPath and return a
// corresponding *auth2.Config. The config file should be obtained from the Google
// Developers Console's "Credentials" page
func getGoogleAuthConfig(oauthConfigPath string) (*oauth2.Config, error) {
	jsonKey, err := ioutil.ReadFile(oauthConfigPath)
	if err != nil {
		return nil, err
	}

	conf, err := google.ConfigFromJSON(jsonKey, "profile")
	if err != nil {
		return nil, err
	}

	conf.Scopes = []string{"email"} // the scopes we need for our app
	return conf, nil
}

func extractPath(next string) string {
	nextURL, err := url.Parse(next)
	if err != nil {
		return "/"
	}
	return nextURL.Path
}
