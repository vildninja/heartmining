package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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
		defer file.Close()
		if err != nil {
			fmt.Println(err)
			return err
		}

		data := xml.NewDecoder(file)
		parsed := XMLMessage{}
		data.Decode(&parsed)

		fmt.Println(len(parsed.DataSet.Series), len(parsed.DataSet.Series[0].SeriesKey.Value))
		fmt.Println(parsed.DataSet.Series[0].Obs[0])

		table := make(map[string]map[int]string)
		minYear := 1000000
		maxYear := 0

		for _, series := range parsed.DataSet.Series {
			row := make(map[int]string)
			country := ""

			for _, val := range series.SeriesKey.Value {
				if val.Concept == "COU" {
					country = val.Value
					break
				}
			}

			for _, cell := range series.Obs {
				year, err := strconv.ParseInt(cell.Time, 10, 32)
				if err != nil {
					fmt.Println(err)
				}

				row[int(year)] = cell.ObsValue.Value

				if int(year) < minYear {
					minYear = int(year)
				}
				if int(year) > maxYear {
					maxYear = int(year)
				}
			}
			table[country] = row
		}

		csv, err := os.Create(info.Name() + ".csv")
		defer csv.Close()
		if err != nil {
			fmt.Println(err)
		}

		//header
		fmt.Fprint(csv, "Country")
		for y := minYear; y < maxYear; y++ {
			fmt.Fprint(csv, ",", y)
		}
		fmt.Fprint(csv, "\n")

		for cou, row := range table {
			fmt.Fprint(csv, cou)
			for y := minYear; y < maxYear; y++ {
				fmt.Fprint(csv, ",", row[y])
			}
			fmt.Fprint(csv, "\n")
		}
	}
	return err
}
