package weather

import (
    "bytes"
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"

    "github.com/paulrosania/go-charset/charset"
    _ "github.com/paulrosania/go-charset/data"
)

type Report struct {
    Weather     string  `xml:"weather"`
    Temperature string  `xml:"temperature_string"`
    Humidity    string  `xml:"relative_humidity"`
    Wind        string  `xml:"wind_string"`
    Pressure    string  `xml:"pressure_string"`
    Dewpoint    string  `xml:"dewpoint_string"`
    Visibility  string  `xml:"visibility_mi"`
}

func GetWeatherFromWeb(weatherUrl string) (*Report, error) {
    xmlData, err := FetchWeatherXML(weatherUrl)
    if err != nil { return nil, err }
    return ParseWeather(xmlData)
}

func GetWeatherFromFile(weatherPath string) (*Report, error) {
    xmlData, err := LoadWeatherXMLFile(weatherPath)
    if err != nil { return nil, err }
    return ParseWeather(xmlData)
}

func FetchWeatherXML(weatherUrl string) ([]byte, error) {
    resp, err := http.Get(weatherUrl)
    if err != nil { return nil, err }

    defer resp.Body.Close();
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil { return nil, err }

    return body, nil
}

func LoadWeatherXMLFile(weatherPath string) ([]byte, error) {
    file, err := os.Open(weatherPath)
    if err != nil { return nil, err }

    data_bytes, err := ioutil.ReadAll(file)
    if err != nil { return nil, err }

    return data_bytes, nil
}

func ParseWeather(weatherXml []byte) (*Report, error) {
    // Praise be to StackOverflow 6002619
    decoder := xml.NewDecoder(bytes.NewBuffer(weatherXml))
    decoder.CharsetReader = charset.NewReader
    var result Report
    err := decoder.Decode(&result)
    return &result, err
}
