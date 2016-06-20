package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/ishuah/batian/utils"
	"github.com/ishuah/batian/models"
	"github.com/ishuah/batian/handlers"
)

type Route struct {
	Name		string
	Method		string
	Pattern 	string
	HandlerFunc	http.HandlerFunc
}

type Routes []Route

func NewRouter(routes Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler

        handler = route.HandlerFunc
        handler = utils.Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	//Serve static files
	//router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	return router
}

func BuildRoutes(db *models.DbManager) Routes {
	return Routes{

			Route{
				"NewEvent",
				"POST",
				"/api/v1/event",
				handlers.NewEvent(db),
			},

			Route{
				"AllEvents",
				"GET",
				"/api/v1/event",
				handlers.AllEvents(db),
			},
		}
}