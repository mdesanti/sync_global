package ControlGas

import (

    "fmt"
    "encoding/json"
    "strconv"
    "net/http"
    "io/ioutil"
    "../../sync_global/config"

)

type StationRequest struct {

    Numref string
    Timestamp int64
    Tag string
    Reader string
    Pump_number string
    Station_id int64
    Type string

}

type Preset struct {
    // stored procedure
    // uspAuthorizationClientDetailsGetByVehicleID
    Id string               `json:"id"`
    Validation_bool bool    `json:"validation_bool"`
    // SendClient
    Tipo string             `json:"tipo"`       // always 3
    Codigo string           `json:"codigo"`     // ClientNumber
    Nombre string           `json:"nombre"`     // OfficialName
    Direccion1 string       `json:"direccion1"` // blank
    Direccion2 string       `json:"direccion2"` // blank
    Direccion3 string       `json:"direccion3"` // blank
    Rfc string              `json:"rfc"`        // RFC
    Tarjeta string          `json:"tarjeta"`    // PublicID
    Grupo string            `json:"grupo"`      // Nickname
    Placas string           `json:"placas"`     // grab from validation
    Puntos string           `json:"puntos"`     // always 0
    Cdc string              `json:"cdc"`        // ASK ALON. MAY NOT ALWAYS BE ZERO
    // SendControl
    Codprd string           `json:"codprd"`     // 0, 1, 2 or 3. ASK ALON      
    Debsdo string           `json:"debsdo"`     // 0
    Debfch string           `json:"debfch"`     // 0
    Debnro string           `json:"debnro"`     // 0
    Debcan string           `json:"debcan"`     // 0
    Ultodm string           `json:"ultodm"`     // 0
    Ultcar string           `json:"ultcar"`     // 0
    Limcar string           `json:"limcar"`     // value. ASK ALON// preset_volume
    Limdia string           `json:"limdia"`     // value. ASK ALON// ODOMETER (when there's an odometer)
    Limsem string           `json:"limsem"`     // value. ASK ALON// value (odometer value when there's an odometer)
    Limmes string           `json:"limmes"`     // 0
    Acudia string           `json:"acudia"`     // 0
    Acusem string           `json:"acusem"`     // 0
    Acumes string           `json:"acumes"`     // 0
    Ptosdo string           `json:"ptosdo"`     // 0
    Limtur string           `json:"limtur"`     // value. ASK ALON
    Ulttur string           `json:"ulttur"`     // 0
    Acutur string           `json:"acutur"`     // value. ASK ALON
    Limprd string           `json:"limprd"`     // 0
    Acuprd string           `json:"acuprd"`     // 0
    Crefch string           `json:"crefch"`     // 0
    Crenro string           `json:"crenro"`     // 0
    Crecan string           `json:"crecan"`     // 0
    Crefch2 string          `json:"crefch2"`    // 0
    Crenro2 string          `json:"crenro2"`    // 0
    Crecan2 string          `json:"crecan2"`    // 0
    Debfch2 string          `json:"debfch2"`    // 0
    Debnro2 string          `json:"debnro2"`    // 0
    Debcan2 string          `json:"debcan2"`    // 0

}

func GetPreset_Card(w http.ResponseWriter, r *http.Request) {

    fmt.Println(`{"src":"ControlGas","method":"GetPreset_Card","type":"info","message":"init"}`)
   
}

func GetPreset_Tag(w http.ResponseWriter, r *http.Request) {

    fmt.Println(`{"src":"ControlGas","method":"GetPreset_Tag","GetPreset_Tag":"info","message":"init"}`)

    if r.Body == nil {
        http.Error(w, "Please send a request body", 400)
    }

    request_body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println(err)
    }
    //fmt.Println("step1")
    req := parseStationRequest(request_body)// returns struct
    //fmt.Println("step2")
    validation_req := createValidationRequest(req)// returns json
    //fmt.Println("step3")
    validation_b := send(config.C.ENDPOINT_VALIDATION, validation_req)
    //fmt.Println("step4")
    validation_res := parseValidationResponse(validation_b)// returns struct
    //fmt.Println("step5")
    trx_req := createTrxRequest(req, validation_res)// returns json
    //fmt.Println("step6")
    trx_b := send(config.C.ENDPOINT_TRX_SYSTEM, trx_req)
    //fmt.Println("step7")
    trx_res := parseTrxResponse(trx_b)// returns struct
    //fmt.Println("step8")
    data_req := createDataRequest(validation_res)// returns json
    //fmt.Println("step9")
    data_b := send(config.C.ENDPOINT_DATA, data_req)
    //fmt.Println("step10")
    data_res := parseDataResponse(data_b)// returns struct
    //fmt.Println("step11")
    preset := createPreset(trx_res, data_res)// returns json

    w.WriteHeader(http.StatusOK)
    w.Write(preset)

}

func parseStationRequest(request []byte) *StationRequest {

    fmt.Println(`{"src":"ControlGas","method":"parseStationRequest","type":"info","message":"init"}`)

    var request_json StationRequest
    if err := json.Unmarshal(request, &request_json); err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(request))

    return &request_json

}

func createPreset(trx_res *TrxResponse, data_res *DataResponse) []byte {

    fmt.Println(`{"src":"ControlGas","method":"createPreset","type":"info","message":"init"}`)

    var preset Preset

    preset.Id = trx_res.Id              // id (grab from trx)
    preset.Validation_bool = trx_res.Validation_bool

    preset.Tipo = "3"                   // always 3
    preset.Codigo = data_res.Codigo     // (grab from data)
    preset.Nombre = data_res.Nombre     // (grab from data)
    preset.Direccion1 = " "
    preset.Direccion2 = " "
    preset.Direccion3 = " "
    preset.Rfc = data_res.Rfc           // (grab from data)
    preset.Tarjeta = data_res.Tarjeta   // (grab from data)
    preset.Grupo = data_res.Grupo       // (grab from data)
    preset.Placas = data_res.Placas     // (grab from data)
    preset.Puntos = "0"
    preset.Cdc = "0"                    // may have value in the case of contingency?

    preset.Codprd = strconv.FormatInt(trx_res.Product, 10)// product (grab from trx)str := strconv.FormatInt(n, 10)
    preset.Debsdo = "0"
    preset.Debfch = "0"
    preset.Debnro = "0"
    preset.Debcan = "0"
    preset.Ultodm = "0"
    preset.Ultcar = "0"
    preset.Limcar = strconv.FormatInt(trx_res.Preset_volume, 10)// preset_volume (grab from trx)
    preset.Limdia = "0"                     // ODOMETER (0 when odometer value is 0) 
    preset.Limsem = "0"                     // odometer value (0 when value is 0) (grab from data)
    preset.Limmes = "0"
    preset.Acudia = "0"
    preset.Acusem = "0"
    preset.Acumes = "0"
    preset.Ptosdo = "0"

    if trx_res.Validation_bool {

        preset.Limtur = "200"               // has value (larger than Acutur when approved)
        preset.Acutur = "100"               // has value

    } else {

        preset.Limtur = "100"               // has value (larger than Acutur when approved)
        preset.Acutur = "200"               // has value

    }
    
    preset.Ulttur = "0"
    preset.Limprd = "0"
    preset.Acuprd = "0"
    preset.Crefch = "0"
    preset.Crenro = "0"
    preset.Crecan = "0"
    preset.Crefch2 = "0"
    preset.Crenro2 = "0"
    preset.Crecan2 = "0"
    preset.Debfch2 = "0"
    preset.Debnro2 = "0"
    preset.Debcan2 = "0"

    preset_json, err := json.Marshal(preset)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(`{"src":"ControlGas","method":"createPreset","type":"info","message":"Preset created"}` + " " + string(preset_json))

    return preset_json

}