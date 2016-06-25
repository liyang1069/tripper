package apis

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/showntop/tripper/stores"
)

var (
	store *stores.Store
)

func Setup() {
	//init the store
	store = stores.NewStore()
	log.Println("the store started...")

	router := httprouter.New()
	router.POST("/Users", CreateUsersHandler)
	router.POST("/Sessions", CreateSessionsHandler)
	router.GET("/Spots", ListSpotsHandler)

	m := &Middleware{router}
	log.Fatal(http.ListenAndServe(":9000", m))
}