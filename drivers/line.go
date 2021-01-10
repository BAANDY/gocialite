package drivers

import (
	"io/ioutil"
	"net/http"

	"github.com/BAANDY/gocialite/structs"
	"golang.org/x/oauth2"
)

const lineDriverName = "line"

func init() {
	lineEndpoint := oauth2.Endpoint{
		AuthURL:   "https://api.line.me/v2/oauth/accessToken",
		TokenURL:  "https://api.line.me/oauth2/v2.1/verify",
		AuthStyle: oauth2.AuthStyleInParams,
	}
	registerDriver(lineDriverName, LineDefaultScopes, LineUserFn, lineEndpoint, LineAPIMap, LineUserMap)
}

// LineUserMap is the map to create the User struct
var LineUserMap = map[string]string{
	"userId":      "ID",
	"email":       "Email",
	"displayName": "FullName",
	"given_name":  "FirstName",
	"family_name": "LastName",
	"pictureUrl":  "Avatar",
}

// LineAPIMap is the map for API endpoints
var LineAPIMap = map[string]string{
	"endpoint":     "https://api.line.me",
	"userEndpoint": "/v2/profile",
}

// LineUserFn is a callback to parse additional fields for User
var LineUserFn = func(client *http.Client, u *structs.User) {
	// Get user ID
	req, err := http.NewRequest("GET", LineAPIMap["endpoint"]+LineAPIMap["authEndpoint"], nil)
	// ...
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	data, err := jsonDecode(res)
	if err != nil {
		return
	}

	u.ID = data["userId"].(string)

	// Fetch other user information
	if v, ok := data["displayName"]; ok {
		u.FullName = v.(string)
	}
	if v, ok := data["pictureUrl"]; ok {
		u.Avatar = v.(string)
	}
}

// LineDefaultScopes contains the default scopes
var LineDefaultScopes = []string{"profile", "email"}
