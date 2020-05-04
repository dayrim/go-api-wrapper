package main

import (
	"cafa-go-wrapper/pkg/cafa"
	"fmt"
	"strconv"
	"time"
)

func main() {
	const (
		app         = "cafa-cli"
		confName    = "companyLevelConfigurationValues"
		confType    = "testing"
		confLevel   = "Warehouse"
		confLevelID = "1"
		//fill your values here
		secret     = ""
		clientCode = ""
		baseUrl    = ""
	)

	cafaCli := cafa.NewClient(baseUrl, nil)

	//compose a config to save
	configToSave := &cafa.Configuration{
		Application: app,
		Level:       confLevel,
		LevelID:     confLevelID,
		Type:        confType,
		Name:        confName,
		Value: struct {
			Field1 string
			Field2 string
		}{
			"oneConfValue",
			"secondConfValue",
		},
	}

	timestamp := strconv.Itoa(int(time.Now().Unix()))
	hmac, err := cafa.NewHMAC(clientCode, "/configuration", secret, timestamp)
	if err != nil {
		panic(err)
		return
	}
	//put your credentials here
	headers := map[string]string{
		"clientCode": clientCode,
		"timestamp":  timestamp,
		"hmac-auth":  hmac,
	}

	fmt.Println("saving")
	resp, err := cafaCli.Save(configToSave, headers)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Message)
	fmt.Println(resp.StatusCode)

	fmt.Println("deleting")
	resp, err = cafaCli.Delete(configToSave, headers)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Message)
	fmt.Println(resp.StatusCode)
}
