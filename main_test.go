package main

import (
	"testing"
	"github.com/spf13/viper"
	"encoding/json"
	"fmt"
)

func TestGetConfig(t *testing.T) {
	filename := "config_example"
	getConfig(filename)

	accessToken := viper.GetString("access_token")
	if accessToken != "abcdefg" {
		t.Fatal("failed test")
	}
}

func TestFetchData(t *testing.T) {
	getConfig("config")

	token := viper.GetString("access_token")
	data := fetchData(token,"20180825")

	spireData := new(Spires)
	err := json.Unmarshal(data,spireData)
	if err != nil {
		fmt.Errorf("%s",err)
	}

	for key, value := range *spireData {
		fmt.Println(key)
		fmt.Println(value)
	}

}
