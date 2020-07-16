package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/joho/godotenv"
)

func main() {
	mood, city := Info()
	// Load the .env file in the current directory
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// api key
	key := os.Getenv("KEY")
	// base url
	base := "https://api.openweathermap.org/data/2.5/weather?q=%s&units=imperial&appid=%s"
	// api url
	api := fmt.Sprintf(base, city, key)
	// request to api
	resp, err := http.Get(api)
	if err != nil {
		log.Fatalln(err)
	}
	// response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// weather structure
	var t weather
	er := json.Unmarshal(data, &t)
	if er != nil {
		panic(er)
	}
	feelsLike := int(t.Temp.FeelsLike)

	s := "It doesn't matter that you're %s because it's %d degrees Fahrenheit in %s!"
	reply := fmt.Sprintf(s, mood, feelsLike, city)
	fmt.Println(reply)
}

// Info gets mood and city from user
func Info() (string, string) {
	qs := []*survey.Question{
		{
			Name:     "mood",
			Prompt:   &survey.Input{Message: "How are you feeling?"},
			Validate: survey.Required,
		},
		{
			Name:     "city",
			Prompt:   &survey.Input{Message: "What city would you like to go to?"},
			Validate: survey.Required,
		},
	}

	answers := struct {
		Mood string
		City string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		log.Fatal(err)
	}

	return answers.Mood, answers.City
}

// weather structure
type weather struct {
	// City
	City string `json:"name"`
	// Temperature structure
	Temp m `json:"main"`
}
type m struct {
	// Temperature
	Temperature float32 `json:"temp"`
	// FeelsLike
	FeelsLike float32 `json:"feels_like"`
	// TempMin
	TempMin float32 `json:"temp_min"`
	// TempMax
	TempMax float32 `json:"temp_max"`
	// Pressure
	Pressure float32 `json:"pressure"`
	// Humidity
	Humidity float32 `json:"humidity"`
}
