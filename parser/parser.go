package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
)

type XMLMessage struct {
	Name    xml.Name `xml:"message"`
	DataSet *XMLDataSet
}

type XMLDataSet struct {
	Name   xml.Name `xml:"DataSet"`
	Series []*XMLSeries
}

type XMLSeries struct {
	Name      xml.Name `xml:"Series"`
	SeriesKey XMLSeriesKey
	Obs       []*XMLObs
}

type XMLSeriesKey struct {
	Name  xml.Name `xml:"SeriesKey"`
	Value []*XMLValue
}

type XMLValue struct {
	Name    xml.Name `xml:"Value"`
	Concept string   `xml:"concept,attr"`
	Value   string   `xml:"value,attr"`
}

type XMLObs struct {
	Name     xml.Name `xml:"Obs"`
	Time     string
	ObsValue XMLObsValue
}

type XMLObsValue struct {
	Name  xml.Name `xml:"ObsValue"`
	Value string   `xml:"value,attr"`
}

func main() {
	filepath.Walk(".", WalkFunc)
}

func WalkFunc(path string, info os.FileInfo, err error) error {
	fmt.Println(path)
	if !info.IsDir() && filepath.Ext(path) == ".xml" {
		fmt.Println(path, "is xml")
		file, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
			return err
		}

		data := xml.NewDecoder(file)
		parsed := XMLMessage{}
		data.Decode(&parsed)

		fmt.Println(len(parsed.DataSet.Series), len(parsed.DataSet.Series[0].SeriesKey.Value))
		fmt.Println(parsed.DataSet.Series[0].Obs[0])
	}
	return err
}
