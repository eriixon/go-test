package sb_json
	
import "encoding/json"

type Coords struct {
	Lon float32
	Lat float32
}
type Weather struct{
	Id int32
	Main string
	Description string
	Icon string
}
type Main struct{
	Temp float32
	Pressure int32
	Humidity int32
	Temp_min float32
	Temp_max float32
}
type Wind struct{
	Speed float32
	Deg int32
}
type Clouds struct{
	All int32
}
type Sys struct{
	Type int32
	Id int32
	Message int32
	Country string
	Sunrise int64
	Sunset int64
}

type Forecast struct {
    Coords Coords
    Weather Weather
    Base string
    Main Main
    Visibility int32
    Wind Wind
    Clouds Clouds
    Dt int32
    Sys Sys
    Id int32
    Name string
    Cod int32
}


func ParseJSON(data string) (string, int) {
	var forecast Forecast
	json.Unmarshal([]byte(data), &forecast)
	return forecast.Name,  int(9/5 * (forecast.Main.Temp - 273) + 32)
}