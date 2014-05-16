package main

import (
	"bufio"
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/vildninja/heartmining/worldmap/painter"
	"os"
	"path/filepath"
	"strings"
)

var codes map[string]string
var clusterColors []colorful.Color

func main() {
	codes = painter.CountryCodes()
	clusterColors = []colorful.Color{
		colorful.Color{1, 0, 0},
		colorful.Color{0, 1, 0},
		colorful.Color{0, 0, 1},
		colorful.Color{1, 1, 0},
		colorful.Color{1, 0, 1},
		colorful.Color{0, 1, 1},
		colorful.Color{1, 1, 1},
		colorful.Color{0, 0, 0},
	}

	err := filepath.Walk(".", WalkFunc)
	if err != nil {
		fmt.Println(err)
	}
}

type Country struct {
	name     string
	clusters map[string]int
	max      string
}

func NewCountry(name string) *Country {
	c := new(Country)
	c.clusters = make(map[string]int)
	c.name = codes[name]
	return c
}

func WalkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() && filepath.Ext(path) == ".csv" {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		reader := bufio.NewReader(file)

		line, err := reader.ReadString('\n')
		header := strings.TrimSpace(strings.Split(line, ";"))

		fmt.Printf("Header: %v \n", len(header))

		country := -1
		cluster := -1

		clusters := make(map[string]int)
		countries := make(map[string]*Country)

		for i := range header {
			fmt.Println(strings.ToLower(header[i]))
			switch strings.ToLower(header[i]) {
			case "country":
				country = i
			case "cluster":
				cluster = i
			}
		}

		fmt.Printf("cou: %v, clu: %v \n", country, cluster)

		for {
			line, err = reader.ReadString('\n')
			row := strings.Split(line, ";")
			if len(row) != len(header) {
				break
			}

			if _, ok := clusters[row[cluster]]; !ok {
				clusters[row[cluster]] = len(clusters)
			}

			c, ok := countries[row[country]]
			if !ok {
				c = NewCountry(row[country])
				countries[row[country]] = c
			}

			c.clusters[row[cluster]]++

			fmt.Printf(c.name+" added to "+row[cluster]+" len(clusters) : %v \n", len(clusters))
		}

		svg := painter.NewSVG()

		for _, co := range countries {
			max := -1
			for cl, count := range co.clusters {
				if count > max {
					co.max = cl
					max = count
				}
			}
			svg.Colors[co.name] = clusterColors[clusters[co.max]].Hex()
		}

		svg.Generate(path + ".svg")
	}

	return nil
}
