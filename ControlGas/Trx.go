package ControlGas

import (

    "fmt"
    "encoding/json"

)

type TrxRequest struct {

    Pump_number string      `json:"pump_number"`
    Station_id int64        `json:"station_id"`
    Type string             `json:"type"`
    Vehicle_id int          `json:"vehicle_id"`
    Validation_bool bool    `json:"validation_bool"`

}

type TrxResponse struct {

    Id string
    Product int64
    Preset_money int64
    Preset_volume int64
    Validation_bool bool

}

func createTrxRequest(request_json *StationRequest, validation_res *ValidationResponse) []byte {

    fmt.Println(`{"src":"ControlGas","method":"createTrxRequest","type":"info","message":"init"}`)

    var trx_req TrxRequest
    trx_req.Pump_number = request_json.Pump_number
    trx_req.Station_id = request_json.Station_id
    trx_req.Vehicle_id = validation_res.Vehicle_id
    trx_req.Validation_bool = validation_res.Validation_bool
    trx_req.Type = request_json.Type
    trx_req_json, err := json.Marshal(trx_req)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(trx_req_json))

    return trx_req_json

}

func parseTrxResponse(response []byte) *TrxResponse {

    fmt.Println(`{"src":"ControlGas","method":"parseTrxResponse","type":"info","message":"init"}`)

    var response_json TrxResponse
    if err := json.Unmarshal(response, &response_json); err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(response))

    return &response_json

}