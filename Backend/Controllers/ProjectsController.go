package Controllers

import (
	"Backend/Models"
	"Backend/Utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ProjectsController struct{}

func singleProjectHandler(res http.ResponseWriter, req *http.Request) {
	var err error

	// CHECK WE HAVE A VALID SLUG
	vars := mux.Vars(req)
	slug := vars["slug"]
	if slug == "" {
		res.WriteHeader(http.StatusBadRequest)
		_, err = res.Write(CreateErrorResponse("slug is missing. It should be in the format: protocol://host:port/projects/slug-slug"))
		if err != nil {
			log.Println(err.Error())
			return
		}
		return
	}

	// CHECK WE HAVE A VALID LOCALE
	locale := req.URL.Query().Get("locale")
	if locale == "" {
		res.WriteHeader(http.StatusBadRequest)
		_, err = res.Write(CreateErrorResponse("locale is missing"))
		if err != nil {
			log.Println(err.Error())
			return
		}
		return
	}

	// GET PROJECT BY SLUG
	project, err := Models.GetProjectBySlug(slug, locale)

	if err != nil {
		log.Println(err.Error())
		res.WriteHeader(http.StatusNotFound)
		_, err = res.Write(make([]byte, 0))
		if err != nil {
			log.Println(err.Error())
			return
		}
		return
	}

	// CONVERT TO A JSON STRING
	marhsalledProject, err := json.Marshal(project)
	if err != nil {
		log.Println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
	}

	// WRITE THE DATA
	res.WriteHeader(http.StatusOK)
	_, err = res.Write(marhsalledProject)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func multipleProjectsHandler(res http.ResponseWriter, req *http.Request) {
	var err error

	// CHECK WE HAVE A VALID LOCALE
	locale := req.URL.Query().Get("locale")

	if locale == "" {
		res.WriteHeader(http.StatusBadRequest)

		_, err = res.Write(CreateErrorResponse("locale is missing"))
		if err != nil {
			log.Println(err.Error())
			return
		}
		return
	}

	// GET PROJECTS
	projects, err := Models.GetProjects(locale)
	if err != nil {
		log.Println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		_, err = res.Write(make([]byte, 0))
		if err != nil {
			log.Println(err.Error())
			return
		}
		return
	}

	// CONVERT TO A JSON STRING
	marshalledData, err := json.Marshal(projects)
	if err != nil {
		log.Println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		_, err = res.Write(make([]byte, 0))
		if err != nil {
			log.Println(err.Error())
			return
		}
		return
	}

	// WRITE THE DATA
	res.WriteHeader(http.StatusOK)
	_, err = res.Write(marshalledData)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func favoriteProjectsHandler(res http.ResponseWriter, req *http.Request) {
	var err error

	// CHECK WE HAVE A VALID LOCALE
	locale := req.URL.Query().Get("locale")

	if locale == "" {
		res.WriteHeader(http.StatusBadRequest)

		_, err = res.Write(CreateErrorResponse("locale is missing"))
		if err != nil {
			log.Println(err.Error())
			return
		}
		return
	}

	// GET PROJECTS
	projects, err := Models.GetProjects(locale)
	if err != nil {
		log.Println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		_, err = res.Write(make([]byte, 0))
		if err != nil {
			log.Println(err.Error())
			return
		}
		return
	}

	projects = Utils.Filter(projects, func(project Models.ProjectModel) bool {
		return project.IsFavorite
	})

	// CONVERT TO A JSON STRING
	marshalledData, err := json.Marshal(projects)
	if err != nil {
		log.Println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		_, err = res.Write(make([]byte, 0))
		if err != nil {
			log.Println(err.Error())
			return
		}
		return
	}

	// WRITE THE DATA
	res.WriteHeader(http.StatusOK)
	_, err = res.Write(marshalledData)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func (ProjectsController) Setup(r *mux.Router) {
	r.HandleFunc("/api/v1/projects", multipleProjectsHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/projects/favorites", favoriteProjectsHandler).Methods(http.MethodGet)
	// HAS TO BE LAST OTHERWISE IT CATCHES ALL OTHER REQUESTS
	r.HandleFunc("/api/v1/projects/{slug:[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*}", singleProjectHandler).Methods(http.MethodGet)
}
