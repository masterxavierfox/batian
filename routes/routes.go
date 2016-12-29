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
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	return router
}

func BuildRoutes(db *models.DbManager) Routes {
	return Routes{
			Route{
				"HandleWebSocket",
				"GET",
				"/ws",
				handlers.HandleWebSocket(db),
			},
			Route{
				"AppDashboard",
				"GET",
				"/",
				handlers.AppDashboard,
			},
			Route{
				"NewEvent",
				"POST",
				"/log",
				handlers.NewEvent(db),
			},
			Route{
				"ListApps",
				"GET",
				"/api/v1/app",
				handlers.AllApps(db),
			},
			Route{
				"AppDetails",
				"GET",
				"/api/v1/app/{appID:[a-z0-9]+}",
				handlers.AppDetails(db),
			},
			Route{
				"UpdateApp",
				"POST",
				"/api/v1/app/{appID:[a-z0-9]+}",
				handlers.UpdateApp(db),
			},
			Route{
				"UpdateApp",
				"DELETE",
				"/api/v1/app/{appID:[a-z0-9]+}",
				handlers.DeleteApp(db),
			},
			Route{
				"AppAnalysis",
				"GET",
				"/api/v1/app/{appID:[a-z0-9]+}/analysis/{duration:[0-9]+}",
				handlers.AppAnalysis(db),
			},
			Route{
				"NewApp",
				"POST",
				"/api/v1/app",
				handlers.NewApp(db),
			},
		}
}