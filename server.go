package main

import (
 "github.com/appointment-octopus/auth/routers"
 "github.com/appointment-octopus/auth/settings"
 "github.com/codegangsta/negroni"
 "net/http"
 "github.com/appointment-octopus/auth/core/db/migration"
)

func main() {
 settings.Init()
 migration.AutoMigration()
 router := routers.InitRoutes()
 n := negroni.Classic()
 n.UseHandler(router)
 http.ListenAndServe(":5000", n)
}
