package main

import (

    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "../sync_global/ControlGas"
    "../sync_global/config"

)

func main() {

    fmt.Println(`{"src":"sync_global","method":"main","type":"info","message":"SYNC_GLOBAL started"}`)

    config.Init()

    listenHTTP()

}

func listenHTTP() {

    fmt.Println("[HTTP] SYNC_GLOBAL listening on " + config.C.HOST + ":" + config.C.PORT)  

    router := mux.NewRouter()
    // ControlGas
    router.HandleFunc("/ControlGas/getPreset/tag", ControlGas.GetPreset_Tag).Methods("POST")
    router.HandleFunc("/ControlGas/getPreset/card", ControlGas.GetPreset_Card).Methods("POST")
    router.HandleFunc("/ControlGas/TerminaTRAN", ControlGas.TerminaTRAN).Methods("POST")
    router.HandleFunc("/ControlGas/CompletaTRAN", ControlGas.CompletaTRAN).Methods("POST")
    router.HandleFunc("/ControlGas/CancelTRAN", ControlGas.CancelTRAN).Methods("POST")

    if err := http.ListenAndServe(config.C.HOST + ":" +config.C.PORT, router); err != nil {
        fmt.Println(err)
    }

}