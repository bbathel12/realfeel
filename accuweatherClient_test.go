package realfeel_test

import (
	"github.com/bbathel12/realfeel"
)

var testString string = `{
  "response": {
  "version":"0.1",
  "termsofService":"http://www.wunderground.com/weather/api/d/terms.html",
  "features": {
  "conditions": 1
  }
    }
  , "current_observation": {
        "image": {
        "url":"http://icons.wxug.com/graphics/wu2/logo_130x80.png",
        "title":"Weather Underground",
        "link":"http://www.wunderground.com"
        },
        "display_location": {
        "full":"New Orleans, LA",
        "city":"New Orleans",
        "state":"LA",
        "state_name":"Louisiana",
        "country":"US",
        "country_iso3166":"US",
        "zip":"70112",
        "magic":"1",
        "wmo":"99999",
        "latitude":"29.95999908",
        "longitude":"-90.08000183",
        "elevation":"0.9"
        },
        "observation_location": {
        "full":"City Hall, New Orleans, Louisiana",
        "city":"City Hall, New Orleans",
        "state":"Louisiana",
        "country":"US",
        "country_iso3166":"US",
        "latitude":"29.951977",
        "longitude":"-90.076790",
        "elevation":"75 ft"
        },
        "estimated": {
        },
        "station_id":"KLANEWOR118",
        "observation_time":"Last Updated on March 6, 3:49 PM CST",
        "observation_time_rfc822":"Tue, 06 Mar 2018 15:49:26 -0600",
        "observation_epoch":"1520372966",
        "local_time_rfc822":"Tue, 06 Mar 2018 15:50:25 -0600",
        "local_epoch":"1520373025",
        "local_tz_short":"CST",
        "local_tz_long":"America/Chicago",
        "local_tz_offset":"-0600",
        "weather":"Clear",
		"temperature_string":"66.2 F (19.0 C)",
        "temp_f":66.2,
        "temp_c":19.0,
        "relative_humidity":"37%",
        "wind_string":"From the WNW at 7.6 MPH Gusting to 10.1 MPH",
        "wind_dir":"WNW",
        "wind_degrees":283,
        "wind_mph":7.6,
        "wind_gust_mph":"10.1",
        "wind_kph":12.2,
        "wind_gust_kph":"16.3",
        "pressure_mb":"1013",
        "pressure_in":"29.93",
        "pressure_trend":"-",
        "dewpoint_string":"39 F (4 C)",
        "dewpoint_f":39,
        "dewpoint_c":4,
        "heat_index_string":"NA",
        "heat_index_f":"NA",
        "heat_index_c":"NA",
        "windchill_string":"NA",
        "windchill_f":"NA",
        "windchill_c":"NA",
        "feelslike_string":"66.2 F (19.0 C)",
        "feelslike_f":"66.2",
        "feelslike_c":"19.0",
        "visibility_mi":"10.0",
        "visibility_km":"16.1",
        "solarradiation":"115",
        "UV":"2.0","precip_1hr_string":"0.00 in ( 0 mm)",
        "precip_1hr_in":"0.00",
        "precip_1hr_metric":" 0",
        "precip_today_string":"1.04 in (26 mm)",
        "precip_today_in":"1.04",
        "precip_today_metric":"26",
        "icon":"clear",
        "icon_url":"http://icons.wxug.com/i/c/k/clear.gif",
        "forecast_url":"http://www.wunderground.com/US/LA/New_Orleans.html",
        "history_url":"http://www.wunderground.com/weatherstation/WXDailyHistory.asp?ID=KLANEWOR118",
        "ob_url":"http://www.wunderground.com/cgi-bin/findweather/getForecast?query=29.951977,-90.076790",
        "nowcast":""
    }
}`

func ExampleOutput() {
	//	getData()
	useCache, cache := realfeel.GetCache()
	if useCache {
		current_weather := realfeel.Unmarshal([]byte(cache))
		current_weather.Output()
	} else {
		current_weather := realfeel.Unmarshal([]byte(testString))
		// ^ once you have api key here you would use
		// current_weather := realfeel.Unmarshal(realfeel.GetData)
		current_weather.Output()
		realfeel.WriteCache(current_weather)
	}
	// Output: {"attachments":[{"image_url":"http://amberandbrice.com/realfeel/realfeeltm.gif"}],"response_type":"in_channel","text":"Weather Data for: City Hall, New Orleans, Louisiana \nLatitude:  29.951977, Longitude: -90.076790\nElevation:      75 ft\nTemperature         :  66.20° \nReal Feel™          :      66.2 \nHumidity            :       37% \nWind Speed          :  7.60 mph \nUV Index            :       2.0 \n","unfurl_links":true,"unfurl_media":true}

}
