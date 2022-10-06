package services

import (
	"encoding/json"
	"golang-hacktive/server/controllers/views"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
)

func GenerateData() *views.Response {
	water := rand.Intn(100-1) + 1
	wind := rand.Intn(100-1) + 1

	data := map[string]interface{}{
		"water": water,
		"wind":  wind,
	}

	file, _ := json.MarshalIndent(data, "", " ")

	err := ioutil.WriteFile("docs/json/alarm.json", file, 0775)

	if err != nil {
		return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
	}

	return views.SuccessResponse(http.StatusCreated, views.M_CREATED, data)
}

func ReadData() *views.Response {
	jsonFile, err := os.Open("docs/json/alarm.json")
	if err != nil {
		return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var alarm *views.Alarm

	json.Unmarshal(byteValue, &alarm)

	return views.SuccessResponse(http.StatusOK, views.M_OK, alarm)
}
