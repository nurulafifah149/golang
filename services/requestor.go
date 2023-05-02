package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type Level struct {
	Wind  string
	Water string
}

func StoreData(data DataWindWater) []byte {
	//convert to JSON
	rData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(rData))
	req.Header.Set("Content-Type", "Application/json")
	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return body
}

func StatusDetect(dataIn []byte) (levelOut Level) {

	var data DataWindWater
	err := json.Unmarshal(dataIn, &data)
	if err != nil {
		panic(err)
	}

	if data.Water <= 5 {
		levelOut.Water = "aman"
	} else if data.Water > 5 && data.Water <= 8 {
		levelOut.Water = "siaga"
	} else if data.Water > 8 {
		levelOut.Water = "bahaya"
	}

	if data.Wind <= 6 {
		levelOut.Wind = "aman"
	} else if data.Wind > 6 && data.Wind <= 15 {
		levelOut.Wind = "siaga"
	} else if data.Wind > 15 {
		levelOut.Wind = "bahaya"
	}

	return
}

func Execute(perulangan int) {

	var data DataWindWater
	var resp []byte
	for i := 0; i < perulangan; i++ {
		data.Wind = rand.Intn(100)
		data.Water = rand.Intn(100)
		resp = StoreData(data)

		err := json.Unmarshal(resp, &data)
		if err != nil {
			panic(err)
		}

		respLevel := StatusDetect(resp)
		fmt.Println(string(resp))
		fmt.Printf("Status Wind : %v\n", respLevel.Wind)
		fmt.Printf("Status Water : %v\n", respLevel.Water)
		time.Sleep(time.Second * 15)
	}
}
