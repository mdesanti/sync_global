package ControlGas

import (

    "fmt"
    "net/http"
    "io/ioutil"

)

func CompletaTRAN(w http.ResponseWriter, r *http.Request) {

    fmt.Println(`{"src":"ControlGas","method":"CompletaTRAN","type":"info","message":"init"}`)

    if r.Body == nil {
        http.Error(w, "Please send a request body", 400)
    }

    request_body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(string(request_body))

}