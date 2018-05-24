// Gets weather data from weather underground to display for
// Slack message /realfeel. RealFeel™
package realfeel

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var zip string

func init() {
	godotenv.Load()
	godotenv.Load(".env.cache")
	var zipPtr *string = flag.String("zip", "70121", "Your zip code")
	fmt.Println(zip)
	zip = *zipPtr
}

// CurrentObservation holds the current weather conditions from
// Unmarshaled json data from api call
type CurrentObservation struct {
	Location            map[string]string `json:"display_location"`
	ObservationLocation map[string]string `json:"observation_location"`
	Temp                float32           `json:"temp_f"`
	RealFeel            string            `json:"feelslike_f"`
	Humidity            string            `json:"relative_humidity"`
	WindSpeed           float32           `json:"wind_mph"`
	UVindex             string            `json:"UV"`
}

// WeatherData type holds top level unmarshalled data from api call
type WeatherData struct {
	Response map[string]interface{} `json:"response"`
	Current  CurrentObservation     `json:"current_observation"`
}

type GeoLookup struct {
	Response map[string]interface{} `json:"response"`
	Location LocationResponse       `json:"location"`
}
type LocationResponse struct {
	RequestUrl string `json:"requesturl"`
}

// Outputs weather data in a format digestible by slack slash command
func (w WeatherData) Output() {

	output := map[string]interface{}{
		"response_type": "in_channel",
	}
	attachments := make([]map[string]string, 1)
	attachments[0] = map[string]string{
		"image_url": "http://amberandbrice.com/realfeel/realfeeltm.gif",
	}
	output["attachments"] = attachments

	var text string = ""

	text += fmt.Sprintf("Weather Data for: %s \n", w.Current.ObservationLocation["full"])
	text += fmt.Sprintf("Latitude: %10s, Longitude: %10s\n", w.Current.ObservationLocation["latitude"], w.Current.ObservationLocation["longitude"])
	text += fmt.Sprintf("Elevation: %10s\n", w.Current.ObservationLocation["elevation"])

	//	text += fmt.Sprint(`
	//
	//  _____                  _                          _              _        _                        _      _
	//|_      _|              |  |                      (  )          |  |    |  |                    |  |  |  |
	//    |  |  ___      __|  |  __  _  _      _|/  ___    |  |    |  |  ___    __  _|  |_|  |__      ___  _  __
	//    |  |/  _  \  /  _    |/  _    |  |  |  |  /  __|  |  |/\|  |/  _  \/  _    |  __|  '_  \  /  _  \  '__|
	//    |  |  (_)  |  (_|  |  (_|  |  |_|  |  \__  \  \    /\    /    __/  (_|  |  |_|  |  |  |    __/  |
	//    \_/\___/  \__,_|\__,_|\__,  |  |___/    \/    \/  \___|\__,_|\__|_|  |_|\___|_|
	//                                                __/  |
	//                                              |___/
	//`)

	text += fmt.Sprintf("%-20s:%7.2f° \n", "Temperature", w.Current.Temp)
	text += fmt.Sprintf("%-20s:%10s \n", "Real Feel™", w.Current.RealFeel)
	text += fmt.Sprintf("%-20s:%10s \n", "Humidity", w.Current.Humidity)
	text += fmt.Sprintf("%-20s:%6.2f mph \n", "Wind Speed", w.Current.WindSpeed)
	text += fmt.Sprintf("%-20s:%10s \n", "UV Index", w.Current.UVindex)
	output["text"] = text

	output["unfurl_media"] = true
	output["unfurl_links"] = true
	forprinting, _ := json.Marshal(output)
	fmt.Println(string(forprinting))
}

// Gets cache if it exists and if it is less than 1 hour old it will
// return true as the first value, to show that this cache is valid
func GetCache() (bool, string) {
	useCache := false
	cache := ""

	timestamp := os.Getenv("timestamp")
	if len(timestamp) != 0 {
		now := time.Now()
		timeInSeconds, _ := strconv.Atoi(timestamp)
		lastUpdate := time.Unix(int64(timeInSeconds), 0)
		diff := now.Sub(lastUpdate)
		if diff.Hours() < 1 {
			useCache = true
		}
		cache = os.Getenv("weatherData")
	}
	return useCache, cache
}

// Writes new data to cache with new timestamp
func WriteCache(data WeatherData) {
	t := time.Now()
	var strdata []byte
	var cache map[string]string = make(map[string]string)
	strdata, _ = json.Marshal(data)
	cache["timestamp"] = strconv.FormatInt(t.Unix(), 10)
	cache["weatherData"] = string(strdata)
	godotenv.Write(cache, ".env.cache")
}

// Makes api call to weather underground and returns the response
func GetData() []byte {
	var request_string string

	request_string = GeoLookupRequest(zip)

	url := os.Getenv("API_URL") + os.Getenv("API_KEY") + "/" + os.Getenv("CALL") + request_string
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	//body = append(body[1:], body[:len(body)-1]...)
	return body
}

// Specific unmarshal function for api return data
func Unmarshal(weatherJson []byte) WeatherData {
	var data WeatherData //[]map[string]interface{}

	err := json.Unmarshal(weatherJson, &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func GeoLookupRequest(zip string) (request_url string) {
	var data GeoLookup

	url := os.Getenv("API_URL") + os.Getenv("API_KEY") + "/geolookup/q/" + zip + ".json"
	fmt.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	request_url = data.Location.RequestUrl
	request_url = strings.Replace(request_url, ".html", ".json", 1)
	return
}
