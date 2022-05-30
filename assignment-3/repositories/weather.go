package repositories

import (
	"assignment-3/models"
	"assignment-3/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

func GetData(data *models.Weather) {
	for {
		data.Status.Wind = utils.RandomNumber(1, 100, 2)
		data.Status.Water = utils.RandomNumber(1, 100, 3)

		jsonData, err := json.Marshal(data)
		log.Printf("%+v\n", data.Status)

		if err != nil {
			log.Fatal(err.Error())
		}
		err = ioutil.WriteFile("repositories/data.json", jsonData, 0644)

		if err != nil {
			log.Fatal(err.Error())
		}
		time.Sleep(15 * time.Second)
	}
}
