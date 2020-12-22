package services

import (
 "github.com/appointment-octopus/auth/core/authentication"
 "github.com/appointment-octopus/auth/services/models"
 "encoding/json"
 jwt "github.com/dgrijalva/jwt-go"
 "github.com/dgrijalva/jwt-go/request"
 "net/http"
 "log"
)

type TokenAuthentication struct {
	Token string `json:"token" form:"token"`
}

func SignUp(requestUser *models.User) (int) {
  _, err := requestUser.CreateUser(); if err != nil {
    log.Fatal(err)
    return http.StatusInternalServerError
  }
  return http.StatusOK
}

func Login(requestUser *models.User) (int, []byte) {
 authBackend := authentication.InitJWTAuthenticationBackend()
if authBackend.Authenticate(requestUser) {
  token, err := authBackend.GenerateToken(requestUser.UUID)
  if err != nil {
   return http.StatusInternalServerError, []byte("")
  } else {
   response, _ := json.Marshal(TokenAuthentication{token})
   return http.StatusOK, response
  }
 }
return http.StatusUnauthorized, []byte("")
}

func RefreshToken(requestUser *models.User) []byte {
 authBackend := authentication.InitJWTAuthenticationBackend()
 token, err := authBackend.GenerateToken(requestUser.UUID)
 if err != nil {
  panic(err)
 }
 response, err := json.Marshal(TokenAuthentication{token})
 if err != nil {
  panic(err)
 }
 return response
}

func Logout(req *http.Request) error {
 authBackend := authentication.InitJWTAuthenticationBackend()
 tokenRequest, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
  return authBackend.PublicKey, nil
 })
 if err != nil {
  return err
 }
 tokenString := req.Header.Get("Authorization")
 return authBackend.Logout(tokenString, tokenRequest)
}
