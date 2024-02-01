package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/winhandsome309/gobook/docs"
	"github.com/winhandsome309/gobook/pkg/models"
	"github.com/winhandsome309/gobook/pkg/routes"
)

//	@title			Gobook Application
//	@description	This is a book management application
//	@version		1.0
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	WinHandsome
//	@contact.url	https://web.facebook.com/winhandsomee/
//	@contact.email	xuanthangnguyen2002@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:9010
//	@BasePath		/book/
func main() {
	r := mux.NewRouter()
	models.Init()
	routes.RegisterBookStoreRoutes(r)
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":9010", r))
}
