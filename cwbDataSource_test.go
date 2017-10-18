package twweather

import (
	"encoding/xml"
	"fmt"
	"os"
	"testing"
)

var (
	sampleXML       []byte
	locationXML     []byte
	exampleElements = make(map[string]float64)
)

func init() {
	exampleElements["ELEV"] = 227
	exampleElements["WDIR"] = 56
	exampleElements["WDSD"] = 1.9
	exampleElements["TEMP"] = 26.6
	exampleElements["HUMD"] = 0.79
	exampleElements["PRES"] = 989.1
	exampleElements["SUN"] = -99
	exampleElements["H_24R"] = 0.0
	exampleElements["H_FX"] = -99
	exampleElements["H_XD"] = -99
	exampleElements["H_FXT"] = -99
}

func createTestError(format string, params ...interface{}) error {
	return fmt.Errorf(fmt.Sprintf(format, params))
}

func matchExampleElements(t *testing.T, station *StationStatus) error {
	for name, expected := range exampleElements {
		element, ok := station.WeatherElements[name]
		if !ok {
			return createTestError("Element %s not found!", name)
		}
		if element != expected {
			return createTestError("Element %s should be %f got %v!", name, expected, element)
		}
		t.Logf("Element match %s => %f = %f", name, expected, element)
	}
	return nil
}

// Test if we can unmarshal location xml with struct stationLocation
func TestParseLocation(t *testing.T) {
	location := new(StationStatus)
	err := xml.Unmarshal(locationXML, location)
	if err != nil {
		t.Fatal(err)
	}
	if location.StationName != "橫山" {
		t.Fail()
	}
	if count := len(location.WeatherElements); count != 11 {
		t.Logf("weather element count of the sample location should be 11. Got %d", count)
		t.Fail()
	}
	matchExampleElements(t, location)
}

func TestParseList(t *testing.T) {

}

func TestLoadData(t *testing.T) {
	t.Skip()
	weather.cwbDataSource = &cwbDataSource{os.Getenv("cwbAPIKey")}
	dataSet := weather.cwbDataSource.loadDataSet(StationStatusDataID)
	t.Log(string(dataSet.RawData))
}
