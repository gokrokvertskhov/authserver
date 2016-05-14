package main

import (
  "github.com/abbot/go-http-auth"
  "github.com/gokrokvertskhov/serverlib/apiserver"
)

func Secret(user, realm string) string {
        if user == "john" {
                // password is "hello"
                return "$1$dlPL2MqE$oQmn16q49SqdmhenQuNgs1"
        }
        return ""
}
var authenticator = auth.NewBasicAuthenticator("example.com", Secret)

var routes = apiserver.Routes {
    apiserver.Route {
		"Refresh",
		"GET",
		"/auth/refresh",
		Refresh,

	},
    apiserver.Route {
		"Validate",
		"GET",
		"/auth/validate",
		Validate,

	},
	apiserver.Route {
		"Token",
		"GET",
		"/auth/token",
		authenticator.Wrap(GenToken),
	},
}