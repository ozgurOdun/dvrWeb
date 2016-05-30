package dvrRestService

import (
	"encoding/json"
	"fmt"
	"github.com/ozgurOdun/dvrWeb/dvrDbOps"
	"net/http"
	"strings"
)

func init() {

}

func QueryActiveDvr(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	query := params.Get(":query")
	var dvrs []*dvrDbOps.Dvr
	var count int64
	if state := strings.Compare(query, "all"); state == 0 {
		dvrs, count = dvrDbOps.GetAllDvr()
		if count <= 0 {
			fmt.Println("Getting all dvrs failed")
			fmt.Fprintf(w, "Failed")
		} else {
			fmt.Println("Getting all dvrs successful")
			answer, err := json.Marshal(dvrs)
			if err == nil {
				fmt.Println(string(answer))
				fmt.Fprintf(w, string(answer))
			} else {
				fmt.Fprintf(w, "Failed")
			}
		}
	} else if state = strings.Compare(query, "alive"); state == 0 {
		dvrs, count = dvrDbOps.GetAliveDvr()
		if count <= 0 {
			fmt.Println("Getting alive dvrs failed")
			fmt.Fprintf(w, "Failed")
		} else {
			fmt.Println("Getting alive dvrs successful")
			answer, err := json.Marshal(dvrs)
			if err == nil {
				fmt.Println(string(answer))
				fmt.Fprintf(w, string(answer))
			} else {
				fmt.Fprintf(w, "Failed")
			}
		}
	}
}

func AddNewDvr(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	dvr := new(dvrDbOps.Dvr)
	dvr.Name = params.Get(":name")
	dvr.IpAddress = params.Get(":ipstring")
	dvr.Version = params.Get(":version")
	dvr.Status = params.Get(":status")
	done := dvrDbOps.AddNewDvr(dvr)
	if done == true {
		fmt.Println("Adding new dvr successful. Name:", dvr.Name)
		answer, err := json.Marshal(dvr)
		if err == nil {
			fmt.Fprintf(w, string(answer))
		} else {
			fmt.Fprintf(w, "Failed")
		}
	} else {
		fmt.Println("Adding new dvr failed. Name:", dvr.Name)
		fmt.Fprintf(w, "Failed")
	}
}

func UpdateDvrStatus(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	name := params.Get(":name")
	newStatus := params.Get((":newstatus"))
	done := dvrDbOps.UpdateDvrStatusByName(name, newStatus)
	if done == true {
		fmt.Println("Updating dvr status by name successful. New status for", name, newStatus)
		fmt.Fprintf(w, "Success")
	} else {
		fmt.Println("Updating dvr status by name failed.", name, newStatus)
		fmt.Fprintf(w, "Failed")
	}
}

func DeleteDvr(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	name := params.Get(":name")
	done := dvrDbOps.DeleteDvrByName(name)
	if done == true {
		fmt.Println("Deleting dvr by name successful", name)
		fmt.Fprintf(w, "Success")
	} else {
		fmt.Println("Deleting dvr by name failed", name)
		fmt.Fprintf(w, "Failed")
	}
}
