package main

import (
    "flag"
    "fmt"

    "weather"
)

var (
    pWeatherUrl = flag.String("weatherurl", "http://www.weather.gov/xml/current_obs/KSAN.xml", "URL for www.weather.gov XML")
    pWeatherFile = flag.String("weatherfile", "", "Read weather XML from this file instead of a URL")
)

func main() {
    flag.Parse()
    var wReport *weather.Report
    var err error
    if *pWeatherFile != "" {
        wReport, err = weather.GetWeatherFromFile(*pWeatherFile)
    } else {
        wReport, err = weather.GetWeatherFromWeb(*pWeatherUrl)
    }
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Temperature: %v\nWind: %v\nPressure: %v\nDewpoint: %v\nVisibility: %v\nHumidity: %v\nWeather: %v\n",
        wReport.Temperature,
        wReport.Wind,
        wReport.Pressure,
        wReport.Dewpoint,
        wReport.Visibility,
        wReport.Humidity,
        wReport.Weather)
}
