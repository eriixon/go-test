package main

import ("fmt"
		"gopkg.in/resty.v1"
		"sb_json"
		"errors"
)

func main(){
	
	// request data from API for Boise
	data, err := owmapi("Boise")
	
	if data != "" {
		city, temp := sb_json.ParseJSON(data)
		fmt.Printf("\nThe forecast in %v for tomorow is %v F", city, temp)
	} else {
		fmt.Printf("%v\n", err)
	}
}
	
func owmapi(city string) (string, error) {

	// GET request
	link := "http://api.openweathermap.org/data/2.5/weather?q=" + city + "&&APPID=bd82255fd0a21fa1238699b9eda2ee35"
	resp, err := resty.R().Get(link)

	// explore response object
	fmt.Printf("\nResponse from Openweathermap.org\n")
	if (err == nil) {		
		fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
		fmt.Printf("\nResponse Time: %v\n", resp.Time())
		return resp.String(), nil
	} else {
		fmt.Printf("\nError: %v\n", err)
		return "", errors.New("\nThe forecast is not available now")
	}
}
