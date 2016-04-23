package conf

import (
	"net/http"
	"github.com/ishuah/batian/controllers"
)

func Route() {
	http.HandleFunc("/", controllers.Index)
}