package main

import (
	"fmt"
	"net/http"
	"os"

	"./ControlGas"
	"./config"
	"github.com/gorilla/mux"
)

func main() {

	fmt.Println(`{"src":"sync_global","method":"main","type":"info","message":"SYNC_GLOBAL started"}`)

	config.Init()

	listenHTTP()

}

func listenHTTP() {

	port := os.Getenv("PORT")

	if port == "" {
		port = config.C.PORT
	}

	fmt.Println("[HTTP] SYNC_GLOBAL listening on " + config.C.HOST + ":" + port)

	router := mux.NewRouter()
	// ControlGas
	router.HandleFunc("/ControlGas/getPreset/tag", ControlGas.GetPreset_Tag).Methods("POST")
	router.HandleFunc("/ControlGas/getPreset/card", ControlGas.GetPreset_Card).Methods("POST")
	router.HandleFunc("/ControlGas/TerminaTRAN", ControlGas.TerminaTRAN).Methods("POST")
	router.HandleFunc("/ControlGas/CompletaTRAN", ControlGas.CompletaTRAN).Methods("POST")
	router.HandleFunc("/ControlGas/CancelTRAN", ControlGas.CancelTRAN).Methods("POST")

	if err := http.ListenAndServe(config.C.HOST+":"+port, router); err != nil {
		fmt.Println(err)
	}

}
