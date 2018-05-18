package ControlGas

import (

    "fmt"
    "net/http"
    "bytes"
    "io/ioutil"
    "time"
    "../../sync_global/config"

)

func send(url string, request_json []byte) []byte {

    fmt.Println(`{"src":"ControlGas","method":"send","type":"info","message":"init"}`)

    getToken()

    request, err := http.NewRequest("POST", url, bytes.NewBuffer(request_json))

    fmt.Println(request)

    request.Header.Set("Content-Type", "application/json")
    request.Header.Set("Accept", "application/json")

    client := &http.Client{
        Timeout: time.Duration(config.C.TIMEOUT) * time.Second,
    }
    response, err := client.Do(request)

    if err != nil {
        fmt.Println(`{"src":"ControlGas","method":"send","type":"error","message":"` + err.Error() + `}`)
    }

    defer response.Body.Close()

    response_body, err := ioutil.ReadAll(response.Body)
    response_body_s := string(response_body)

    fmt.Println(`{"src":"ControlGas,"method":"send","type":"info","message":"Response received","data":` + response_body_s + `}`)
    /*
    if response.StatusCode == 200 {
    } 
    */
    return response_body

}

func getToken() {

    fmt.Println(`{"src":"ControlGas","method":"getToken","type":"info","message":"init"}`)

}