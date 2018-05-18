package ControlGas

import (

    "fmt"
    "encoding/json"

)

type DataRequest struct {

    Vehicle_id int          `json:"vehicle_id"`

}

type DataResponse struct {

    Validation_bool bool
    Vehicle_id int
    Codigo string// ClientNumber
    Nombre string// OfficialName
    Rfc string// RFC
    Tarjeta string// PublicID
    Grupo string// Nickname
    Placas string// Vehicle plate

}

func createDataRequest(validation_res *ValidationResponse) []byte {

    fmt.Println(`{"src":"ControlGas","method":"createDataRequest","type":"info","message":"init"}`)

    var data_req DataRequest
    data_req.Vehicle_id = validation_res.Vehicle_id
    data_req_json, err := json.Marshal(data_req)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(data_req_json))

    return data_req_json

}

func parseDataResponse(response []byte) *DataResponse {

    fmt.Println(`{"src":"ControlGas","method":"parseDataResponse","type":"info","message":"init"}`)

    var response_json DataResponse
    if err := json.Unmarshal(response, &response_json); err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(response))

    return &response_json

}