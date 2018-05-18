package ControlGas

import (

    "fmt"
    "encoding/json"

)

type ValidationRequest struct {

    Tag string              `json:"tag"`
    Station_id int64        `json:"station_id"`

}

type ValidationResponse struct {

    Validation_bool bool
    Vehicle_id int

}

func createValidationRequest(request_json *StationRequest) []byte {

    fmt.Println(`{"src":"ControlGas","method":"createValidationRequest","type":"info","message":"init"}`)

    var validation_req ValidationRequest
    validation_req.Tag = request_json.Tag
    validation_req.Station_id = request_json.Station_id
    validation_req_json, err := json.Marshal(validation_req)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(validation_req_json))

    return validation_req_json

}

func parseValidationResponse(response []byte) *ValidationResponse {

    fmt.Println(`{"src":"ControlGas","method":"parseValidationResponse","type":"info","message":"init"}`)

    var response_json ValidationResponse
    if err := json.Unmarshal(response, &response_json); err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(response))

    return &response_json

}