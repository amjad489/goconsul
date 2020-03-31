package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func getKv(url string, token string) []interface{} {
	var resKeys []interface{}
	req, err := http.NewRequest("GET", url, nil)
	if token != "" {
		req.Header.Add("X-Consul-Token", token)
	}

	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		//fmt.Println(contents)
		err = json.Unmarshal(contents, &resKeys)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return resKeys
}
