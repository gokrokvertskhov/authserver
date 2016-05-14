package main

import (
	//"encoding/json"
	"fmt"
	 "log"
	//"io"
	//"io/ioutil"
	"net/http"
	"github.com/abbot/go-http-auth"
	//"github.com/gokrokvertskhov/gauth"
	//"github.com/gokrokvertskhov/gauth/provider/github"
	"github.com/gokrokvertskhov/serverlib/jwt"
	//"strconv"
	//"github.com/gorilla/mux"
)

func HandlersInit(){
	jwt.ClientInit(Conf.Secure.PublicKey)
	jwt.ServerInit(Conf.Secure.PrivateKey, 3600)
}

func Refresh(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "{\"status\": %s}", "ok")
}

func Validate(w http.ResponseWriter, r *http.Request) {
	log.Print("Trying to validate")
     _, token := jwt.IsValidTokenRequest(r)
     if token != nil {
	     if token.Valid {
	     	fmt.Fprintf(w, "{\"status\": %s}", "ok")
	     }else{
	     	w.WriteHeader(http.StatusNotAcceptable)
	     	fmt.Fprintf(w, "{\"err\": %s}", "Expired or invalid")
	     }
	     }else {
	     	w.WriteHeader(http.StatusNotAcceptable)
	     	fmt.Fprintf(w, "{\"err\": %s}", "Expired or invalid")

	     }
}


func GenToken(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	tstring, _ := jwt.CreateToken(123, 60)
	fmt.Fprintf(w, "{\"token\": %s}", tstring)
}

func AuthCallback(w http.ResponseWriter, r *http.Request) {
 //    code := r.URL.Query().Get("code")

	// // Get access token from verification code.
	// provider := GetProvider()
	// token, err := provider.Authorize(code)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// // Get the authenticating user
	// user, err := provider.User(token)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

}