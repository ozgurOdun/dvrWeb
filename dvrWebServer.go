package main

import (
	"fmt"
	"github.com/drone/routes"
	"github.com/ozgur/dvrDbOps"
	"github.com/ozgur/dvrRestService"
	"net/http"
)

const VERSION = "v0.1.0"

func main() {
	fmt.Println("Active dvr database server. Version:", VERSION)
	dvrDbOps.NewDb()
	mux := routes.New()
	mux.Del("/dvr/:name/delete", dvrRestService.DeleteDvr)
	mux.Post("/dvr/:name/:newstatus/update", dvrRestService.UpdateDvrStatus)
	mux.Get("/dvr/:query/query", dvrRestService.QueryActiveDvr)
	mux.Put("/dvr/:name/:ipstring/:version/:status/add", dvrRestService.AddNewDvr)
	http.Handle("/", mux)
	http.ListenAndServe(":8088", nil)

}
