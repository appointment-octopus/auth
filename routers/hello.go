package routers

import (
 "github.com/appointment-octopus/auth/controllers"
 "github.com/appointment-octopus/auth/core/authentication"
 "github.com/codegangsta/negroni"
 "github.com/gorilla/mux"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {
 router.Handle("/test/hello",
  negroni.New(
   negroni.HandlerFunc(authentication.RequireTokenAuthentication),
   negroni.HandlerFunc(controllers.HelloController),
  )).Methods("GET")
 return router
}
