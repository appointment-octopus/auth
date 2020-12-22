package controllers

import (
 "github.com/appointment-octopus/auth/services"
 "github.com/appointment-octopus/auth/services/models"
 "encoding/json"
 "net/http"
)

func decodeUser(r *http.Request) *models.User {
  requestUser := new(models.User)
  decoder := json.NewDecoder(r.Body)
  decoder.Decode(&requestUser)
  return requestUser
}

func SignUp(w http.ResponseWriter, r *http.Request) {
  requestUser := decodeUser(r)
  responseStatus := services.SignUp(requestUser)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(responseStatus)
}

func Login(w http.ResponseWriter, r *http.Request) {
  requestUser := decodeUser(r)
  responseStatus, token := services.Login(requestUser)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(responseStatus)
  w.Write(token)
}

func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
  requestUser := decodeUser(r)
  w.Header().Set("Content-Type", "application/json")
  w.Write(services.RefreshToken(requestUser))
}

func Logout(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
 err := services.Logout(r)
 w.Header().Set("Content-Type", "application/json")
 if err != nil {
  w.WriteHeader(http.StatusInternalServerError)
 } else {
  w.WriteHeader(http.StatusOK)
 }
}

