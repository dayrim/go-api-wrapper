package main

import (
	cafa "cafa-go-wrapper/pkg/cafa"
	"errors"
	"fmt"
	"strconv"
)

func main() {
	const (
		app         = "cafa-cli"
		confName    = "companyLevelConfigurationValues"
		confType    = "testing"
		confLevel   = "Warehouse"
		confLevelID = "1"
		//fill your values here
		sessionKey = ""
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

	//put your credentials here
	headers := map[string]string{
		"clientCode": clientCode,
		"sessionKey": sessionKey,
	}

	resp, err := cafaCli.Save(configToSave, headers)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Message)
	fmt.Println(resp.StatusCode)

	//Update
	const updatedValue = "updated"
	configToSave.Value = updatedValue
	resp, err = cafaCli.Update(configToSave, headers)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Message)
	fmt.Println(resp.StatusCode)

	//Get
	filters := &cafa.Configuration{
		Application: app,
		Level:       confLevel,
	}
	configs, err := cafaCli.Get(filters, headers)
	if err != nil {
		panic(err)
	}
	if configs[0].Value != updatedValue {
		panic(errors.New("update did not change the value"))
	}

	filters.Name = confName
	filters.Type = confType
	filters.Level = confLevel
	resp, err = cafaCli.Delete(filters, headers)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Message)
	fmt.Println(resp.StatusCode)

	/*
		use by ID requests for convenience
	*/

	//create the config
	resp, err = cafaCli.Save(configToSave, headers)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Message)
	fmt.Println(resp.StatusCode)

	//get the config to know it's ID
	configs, err = cafaCli.Get(filters, headers)
	if err != nil {
		panic(err)
	}

	configID, err := strconv.Atoi(configs[0].ID)
	if err != nil {
		panic(err)
	}

	//move it
	moveRequest := &cafa.MoveRequest{
		LevelID: "2",
	}
	fmt.Println("moving by ID")
	resp, err = cafaCli.MoveByID(configID, headers, moveRequest)
	if err != nil {
		panic(err)
	}

	//update it
	updateRequest := &cafa.UpdateByIDRequest{
		Value: updatedValue,
	}
	fmt.Println("updating by ID")
	resp, err = cafaCli.UpdateByID(configID, headers, updateRequest)
	if err != nil {
		panic(err)
	}

	fmt.Println("getting by ID")
	conf, err := cafaCli.GetByID(configID, headers)
	if err != nil {
		panic(err)
	}

	if conf.Value != updatedValue {
		panic(err)
	}

	fmt.Println("deleting by ID")
	resp, err = cafaCli.DeleteByID(configID, headers)
	if err != nil {
		panic(err)
	}

}
