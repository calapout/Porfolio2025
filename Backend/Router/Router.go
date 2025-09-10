package Router

import (
	"Backend/Controllers"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Setup() {
	r := mux.NewRouter()
	r.StrictSlash(true)

	Controllers.ProjectsController{}.Setup(r)
	Controllers.UploadsController{}.Setup(r)

	http.Handle("/", handlers.CORS()(r))
}
