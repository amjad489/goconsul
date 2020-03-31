package core

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/spf13/cast"
	"os"
)

func App(consulAddr string, consulToken string, prefix string, output string) {
	i := 0
	var data [][]string
	consulBaseUrl, keysUrl := getUrl(consulAddr, prefix)
	allKeys := getKv(keysUrl, consulToken)
	if len(allKeys) > 0 {
		for _, keyName := range allKeys {
			i++
			getKvUrl := consulBaseUrl + cast.ToString(keyName)
			rawKeyValue := getKv(getKvUrl, consulToken)
			for _, k := range rawKeyValue {
				if rec, ok := k.(map[string]interface{}); ok {
					rawValue := cast.ToString(rec["Value"])
					decodedValue, err := b64.StdEncoding.DecodeString(rawValue)
					if err != nil {
						fmt.Println(err.Error())
					}
					strSanValue := SanitizeName(string(decodedValue), 75)
					strSanKey := SanitizeName(cast.ToString(keyName), 55)
					d := []string{cast.ToString(i), strSanKey, strSanValue}
					data = append(data, d)
				} else {
					fmt.Printf("data not a map[string]interface{}: %v\n", k)
				}
			}
		}
		if output == "json" {
			jsonOutput(data)
		} else if output == "yml" || output == "yaml" {
			yamlOutput(data)
		} else if output == "xml" {
			xmlOutput(data)
		} else if output == "csv" {
			csvOutput(data)
		}
		tableOutput(data, allKeys)

	} else {
		fmt.Println(" no keys found.")
		os.Exit(1)
	}
}
