package routers
import (
    "github.com/appointment-octopus/auth/controllers"
    "github.com/appointment-octopus/auth/core/authentication"
    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
    router.HandleFunc(
        "/signup",
        controllers.SignUp,
    ).Methods("POST")

    router.HandleFunc(
        "/token-auth",
        controllers.Login,
    ).Methods("POST")

    router.Handle(
        "/refresh-token-auth",
        negroni.New(
            negroni.HandlerFunc(controllers.RefreshToken),
        )).Methods("GET")

    router.Handle(
        "/logout",
        negroni.New(
            negroni.HandlerFunc(
                authentication.RequireTokenAuthentication,
            ),
            negroni.HandlerFunc(controllers.Logout),
        )).Methods("GET")
    return router
}

