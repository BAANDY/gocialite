package drivers

import (
	"net/http"

	"github.com/BAANDY/gocialite/structs"
	"golang.org/x/oauth2/google"
)

const googleDriverName = "google"

func init() {
	registerDriver(googleDriverName, GoogleDefaultScopes, GoogleUserFn, google.Endpoint, GoogleAPIMap, GoogleUserMap)
}

// GoogleUserMap is the map to create the User struct
var GoogleUserMap = map[string]string{
	"id":          "sub",
	"email":       "email",
	"name":        "name",
	"given_name":  "given_name",
	"family_name": "family_name",
	"picture":     "picture",
}

// GoogleAPIMap is the map for API endpoints
var GoogleAPIMap = map[string]string{
	"endpoint":     "https://www.googleapis.com",
	"userEndpoint": "/oauth2/v3/tokeninfo",
}

// GoogleUserFn is a callback to parse additional fields for User
var GoogleUserFn = func(client *http.Client, u *structs.User) {}

// GoogleDefaultScopes contains the default scopes
var GoogleDefaultScopes = []string{"profile", "email"}
