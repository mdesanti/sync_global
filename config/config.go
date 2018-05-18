package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	HOST                string
	PORT                string
	ENDPOINT_VALIDATION string
	ENDPOINT_DATA       string
	ENDPOINT_TRX_SYSTEM string
	TIMEOUT             int64
}

var C Config

func Init() {

	absPath, _ := filepath.Abs("./config/config.txt")

	fmt.Println(`{"src":"config","method":"Init","type":"info","message":"Init"}`)

	file, err := ioutil.ReadFile(absPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := json.Unmarshal(file, &C); err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file))

}
