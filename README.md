# Google Adsense Example with Golang
This project shows that how Google Adsense API works with Golang.

Installation
============
    $ go get -u github.com/odg0318/google-adsense-example
    
Get Started
===========
1. Make sure `Google Adsense API` is activated on Google API Console.
2. Create `OAuth2 Client ID` on Google API Console and set `redirect_uri` to `http://127.0..1:8080/auth`
3. cd $GOPATH/github.com/odg0318/google-adsense-example
4. Set following environments; `GOOGLE_CLIENT_ID`, `GOOGLE_CLIENT_SECRET`, `GOOGLE_REDIRECT_URI`
5. Execute `go build && ./google-adsense-example` and a web server will run.
6. Connect to http://127.0.0.1:8080/ on web browser and select your account to access.
7. Some result of API will be shown on web browser.

Environment Variables
=====================
Environment | Description
----------- | -----------
GOOGLE_CLIENT_ID | Google OAuth2 Client ID
GOOGLE_CLIENT_SECRET | Google OAuth2 Client Secret
GOOGLE_REDIRECT_URI | Redirect URI is used to receive an authorization code from Google.

Other Examples
======================
You can check [here](https://godoc.org/google.golang.org/api/adsense/v1.4) available API list.

### Accounts:list
```golang
client := oauth2Conf.Client(oauth2.NoContext, tok)
service, err := adsense.New(client)

call := service.Accounts.List()
resp, err := call.Do()
```

### Accounts.reports:generate
```golang
client := oauth2Conf.Client(oauth2.NoContext, tok)
service, err := adsense.New(client)

call := service.Accounts.Reports()
call := report.Generate(accountId, fromDate, toDate)
call.Dimension("COUNTRY_CODE","DATE")
call.Metric("EARNINGS")
resp, err := call.Do()
```

With Refresh Token
==================
Refresh Token is described in [documentation](https://developers.google.com/identity/protocols/OAuth2WebServer#offline).
```golang
url := oauth2Conf.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
```

References
==========
* [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
* [https://godoc.org/golang.org/x/oauth2](https://godoc.org/golang.org/x/oauth2)
* [https://godoc.org/google.golang.org/api/adsense/v1.4](https://godoc.org/google.golang.org/api/adsense/v1.4)
