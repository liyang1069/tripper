package apis

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Middleware struct {
	httprouter *httprouter.Router
}

func (m *Middleware) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Println("===========================")
	log.Println(req.Method + "------" + req.URL.Path)
	log.Printf("%v", req)
	err := ParseParams(req.Body)
	if err != nil {
		log.Printf("%v", err)
		renderError(rw, err, http.StatusInternalServerError)
		return
	}
	log.Printf("%v", params)
	m.httprouter.ServeHTTP(rw, req)
	log.Println("===========================")
}
