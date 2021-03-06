package main

import (

    "fmt"
    "net/http"
    //"net/url"
    "github.com/gorilla/mux"
    "encoding/json"
    "bytes"

)

const (

    // port 5000 and name it application.go
    HTTP_PORT_AUTH = "8008"
    HTTP_PORT_SYNC = "8080"
    CONN_HOST = "localhost"

)

func main() {
    
    go listenHTTP()

    for {}
}

func listenHTTP() {

    fmt.Println("[HTTP] SYNC_GLOBAL listening for SYNC_LOCAL on " + CONN_HOST + ":" + HTTP_PORT_SYNC)  

    router := mux.NewRouter()
    router.HandleFunc("/getPreset/tag", getPreset_Tag).Methods("POST")
    router.HandleFunc("/getPreset/card", getPreset_Card).Methods("POST")
    http.ListenAndServe(CONN_HOST + ":" + HTTP_PORT_SYNC, router)

    //if err := http.ListenAndServe(CONN_HOST + ":" + FS_HTTP_PORT, nil); err != nil {
        //log.Fatal(err)
    //}

}

func getPreset_Card(w http.ResponseWriter, r *http.Request) {

    fmt.Println("getPreset_Card()")
    /*
    {   

        "pump_number": 59,
        "station_id": 2402,
        "vehicle_id": 18305,
        "validated": true,
        "validation_code": ["VALIDATION_CODE"],
        "timestamp": XXX,
        "reader_id": XXX,

    }
    */
    if r.Body == nil {

        http.Error(w, "Please send a request body", 400)
    }

    type Data struct {
        Card string// must be caps or will fail Marshal
        Numref string
        Timestamp int64
        Station string
    }

    var d Data 

    err := json.NewDecoder(r.Body).Decode(&d)

    if err != nil {
        http.Error(w, err.Error(), 400)
        return// MODIFY
    }
   
    fmt.Println(d)
    fmt.Println(d.Card)
    fmt.Println(d.Numref)

    validateCard()// send to authorization

    getPreset()// send to trx

    getData()// get other data

    // write JSON response with Preset/Denial + Client Info
    w.WriteHeader(http.StatusOK)

}

func getPreset_Tag(w http.ResponseWriter, r *http.Request) {

    fmt.Println("getPreset_Tag()")
    /*
    {   

        "pump_number": 59,
        "station_id": 2402,
        "vehicle_id": 18305,
        "validated": true,
        "validation_code": ["VALIDATION_CODE"],
        "timestamp": XXX,
        "reader_id": XXX,

    }
    */
    if r.Body == nil {

        http.Error(w, "Please send a request body", 400)
    }

    type Data struct {
        Tag string// must be caps or will fail Marshal
        Pump string
        Reader string
        Numref string
        Timestamp int64
        Station string
    }

    var d Data 

    err := json.NewDecoder(r.Body).Decode(&d)

    if err != nil {
        http.Error(w, err.Error(), 400)
        return// MODIFY
    }
   
    fmt.Println(d)
    fmt.Println(d.Tag)
    fmt.Println(d.Pump)
    fmt.Println(d.Reader)
    fmt.Println(d.Numref)

    validateTag(d.Tag, d.Station)// send to authorization

    getPreset()// send to trx

    getData()// get other data

    // write JSON response with Preset/Denial + Client Info
    w.WriteHeader(http.StatusOK)

}

func validateTag(tag string, station string) {

    fmt.Println("validateTag()")

    type Data struct {// properties must have caps or Marshal will fail

        Tag string
        Station string

    }

    var d Data
    d.Tag = tag
    d.Station = station

    d_json, _ := json.Marshal(d)// d_json is a byte holding the json data

    response, err := http.Post("http://" + CONN_HOST + ":" + HTTP_PORT_AUTH + "/validate/tag", "application/json", bytes.NewBuffer(d_json))
    if err != nil {
        fmt.Println(err)// handle error
    }

    defer response.Body.Close()

    if response.StatusCode == 200 {

        //
    } 
    
}

func validateCard() {

    fmt.Println("validateCard")
    /*
    type Data struct {// properties must have caps or Marshal will fail

        Card string
        Numref string
        Timestamp int64
        Station string

    }

    var d Data
    d.Card = card
    d.Numref = numref
    d.Timestamp = now_ms
    d.Station = "123"

    d_json, _ := json.Marshal(d)// d_json is a byte holding the json data

    response, err := http.Post("http://" + CONN_HOST + ":" + HTTP_PORT_AUTH + "/validate/card", "application/json", bytes.NewBuffer(d_json))
    if err != nil {
        fmt.Println(err)// handle error
    }

    defer response.Body.Close()

    if response.StatusCode == 200 {

        //
    } 
    */
}

func getPreset() {

    fmt.Println("getPreset()")

}

func getData() {

    fmt.Println("getData()")

}