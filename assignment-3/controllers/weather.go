package controllers

import (
	"assignment-3/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func AutoReloadWeb(w http.ResponseWriter, r *http.Request) {

	fileData, err := ioutil.ReadFile("repositories/data.json")

	if err != nil {
		log.Fatal(err.Error())
	}

	var statusData *models.Weather

	err = json.Unmarshal(fileData, &statusData)
	if err != nil {
		log.Fatal(err.Error())
	}

	waterVal := statusData.Status.Water
	windVal := statusData.Status.Wind

	var (
		status string
	)

	switch {
	case waterVal <= 5 && windVal <= 6:
		status = "Aman"
	case waterVal >= 6 && waterVal <= 8 && windVal >= 7 && windVal <= 15:
		status = "Siaga"
	case waterVal > 8 && windVal > 15:
		status = "Bahaya"
	default:
		status = "Not defined"
	}

	data := map[string]string{
		"status":     status,
		"waterValue": strconv.Itoa(waterVal),
		"windValue":  strconv.Itoa(waterVal),
	}

	template, err := template.ParseFiles("templates/index.html")

	if err != nil {
		log.Fatal(err.Error())
	}

	template.Execute(w, data)
}
