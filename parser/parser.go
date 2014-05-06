package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
	err := filepath.Walk(".", WalkFunc)
	if err != nil {
		fmt.Println(err)
	}
	err = filepath.Walk(".", InterpolateCSV)
	if err != nil {
		fmt.Println(err)
	}
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

		table := make(map[string]map[string]map[int]string)
		minYear := 1000000
		maxYear := 0

		for _, series := range parsed.DataSet.Series {
			row := make(map[int]string)

			var country string
			var variable string
			var unit string
			for _, val := range series.SeriesKey.Value {
				switch {
				case val.Concept == "COU":
					country = val.Value
				case val.Concept == "VAR":
					variable = val.Value
				case val.Concept == "UNIT":
					unit = val.Value
				}
			}

			entries := 0

			for _, cell := range series.Obs {
				year, err := strconv.ParseInt(cell.Time, 10, 32)
				if err != nil {
					fmt.Println(err)
				}

				val, err := strconv.ParseFloat(cell.ObsValue.Value, 32)
				if err != nil {
					fmt.Println("Parse: ", err)
				} else {
					entries++
				}

				row[int(year)] = strconv.FormatFloat(val, 'f', 2, 32)

				if int(year) < minYear {
					minYear = int(year)
				}
				if int(year) > maxYear {
					maxYear = int(year)
				}
			}

			if entries > 0 {
				if table[variable+"-"+unit] == nil {
					table[variable+"-"+unit] = make(map[string]map[int]string)
				}
				table[variable+"-"+unit][country] = row
			}
		}

		for fname, data := range table {
			csv, err := os.Create(info.Name() + "-" + fname + ".csv")
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

			for cou, row := range data {
				fmt.Fprint(csv, cou)
				for y := minYear; y < maxYear; y++ {
					fmt.Fprint(csv, ",", row[y])
				}
				fmt.Fprint(csv, "\n")
			}
		}
	}
	return err
}

type DataRow struct {
	country  string
	cells    map[int]float32
	avgSlope float32
}

func NewDataRow(country string) *DataRow {
	dr := new(DataRow)
	dr.country = country
	dr.cells = make(map[int]float32)
	return dr
}

func (d *DataRow) CalcSpan() int {
	min := 12
	max := -1
	for y := range d.cells {
		if y < min {
			min = y
		}
		if y > max {
			max = y
		}
	}
	return max - min
}

func (d *DataRow) CalcSlope() {
	first := -1
	last := -1
	for i := 0; i < 12; i++ {
		if _, ok := d.cells[i]; ok {
			if first < 0 {
				first = i
			}
			last = i
		}
	}

	if last == first {
		d.avgSlope = 0
		return
	}

	fv := d.cells[first]
	lv := d.cells[last]
	d.avgSlope = (lv - fv) / float32(last-first)
}

func (d *DataRow) GetCell(i int) float32 {

	if v, ok := d.cells[i]; ok {
		return v
	}

	before := -1
	for b := i - 1; b >= 0; b-- {
		if _, ok := d.cells[b]; ok {
			before = b
			break
		}
	}

	after := -1
	for a := i + 1; a < 12; a++ {
		if _, ok := d.cells[a]; ok {
			after = a
			break
		}
	}

	if before >= 0 && after >= 0 {
		b := d.cells[before]
		a := d.cells[after]
		inc := (a - b) / float32(after-before)
		dist := float32(i - before)

		return b + inc*dist
	}

	if before >= 0 {
		b := d.cells[before]
		dist := float32(i - before)
		return b + d.avgSlope*dist
	}

	if after >= 0 {
		a := d.cells[after]
		dist := float32(after - i)
		return a - d.avgSlope*dist
	}

	return 0
}

func InterpolateCSV(path string, info os.FileInfo, err error) error {
	if filepath.Ext(path) == ".csv" {
		file, err := os.Open(path)
		defer file.Close()

		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(file)

		rows := make(map[*DataRow]bool)

		scanner.Scan()
		header := scanner.Text()

		for scanner.Scan() {
			row := strings.Split(scanner.Text(), ",")

			data := NewDataRow(row[0])
			for i := 1; i < len(row); i++ {
				if f, err := strconv.ParseFloat(row[i], 32); err == nil {
					data.cells[i-1] = float32(f)
				}
			}
			rows[data] = true
			data.CalcSlope()
		}

		fmt.Println(filepath.Dir(path) + "/interpolated/" + filepath.Base(path))

		output, err := os.Create(filepath.Dir(path) + "/interpolated/" + filepath.Base(path))
		defer output.Close()

		if err != nil {
			return err
		}

		fmt.Fprintln(output, header)

		for r := range rows {
			if r.CalcSpan() > 3 {
				fmt.Fprint(output, r.country)
				for i := 0; i < 12; i++ {
					fmt.Fprint(output, ","+strconv.FormatFloat(float64(r.GetCell(i)), 'f', 2, 32))
				}
				fmt.Fprint(output, "\n")
			}
		}
	}

	return nil
}
