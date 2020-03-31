package core

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cast"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Result struct {
	Key   string
	Value string
}
type Output struct {
	Results []Result
}

var msg = "data exported to file: data."

func yamlOutput(data [][]string) {
	var yamlOutput Output
	for _, v := range data {
		yamlOutput.Results = append(yamlOutput.Results, Result{
			Key:   v[1],
			Value: v[2],
		})
	}
	file, _ := yaml.Marshal(yamlOutput)
	_ = ioutil.WriteFile("data.yaml", file, 0644)
	fmt.Println(msg + "yaml")
	os.Exit(0)
}
func jsonOutput(data [][]string) {
	var jsonOutput Output
	for _, v := range data {
		jsonOutput.Results = append(jsonOutput.Results, Result{
			Key:   v[1],
			Value: v[2],
		})
	}
	file, _ := json.MarshalIndent(jsonOutput, "", " ")
	_ = ioutil.WriteFile("data.json", file, 0644)
	fmt.Println(msg + "json")
	os.Exit(0)
}

func xmlOutput(data [][]string) {
	var xmlOutput Output
	for _, v := range data {
		xmlOutput.Results = append(xmlOutput.Results, Result{
			Key:   v[1],
			Value: v[2],
		})
	}
	file, _ := xml.Marshal(xmlOutput)
	_ = ioutil.WriteFile("data.xml", file, 0644)
	fmt.Println(msg + "xml")
	os.Exit(0)
}

func csvOutput(data [][]string) {
	csvFile, err := os.Create("data.csv")
	if err != nil {
		fmt.Println(err.Error())
	}
	csvWriter := csv.NewWriter(csvFile)
	_ = csvWriter.Write([]string{"id", "key", "value"})
	for _, Row := range data {
		_ = csvWriter.Write(Row)
	}

	csvWriter.Flush()
	fmt.Println(msg + "csv")
	os.Exit(0)
}

func tableOutput(data [][]string, allKeys []interface{}) {
	// display table output
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetCaption(true, "Total keys: "+cast.ToString(len(allKeys)))
	table.SetHeader([]string{"Id", "Key", "Value"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
	fmt.Println(tableString.String())
	// use less `command` to display for easy viewing
	cmd := exec.Command("/usr/bin/less")
	cmd.Stdin = strings.NewReader(tableString.String())
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
