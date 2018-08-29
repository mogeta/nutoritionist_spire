package main

import (
	"github.com/spf13/viper"
	"fmt"
	"net/http"
	"io/ioutil"
	"net/url"
)

type Spires []*SpireData

type SpireData struct {
	Type         string  `json:"type"`
	StartAt      int     `json:"start_at"`
	StopAt       int     `json:"stop_at"`
	Value        float64 `json:"value"`
	SubValue     float64 `json:"sub_value"`
	OriginalType string  `json:"original_type"`
	Comment      string  `json:"comment"`
	ModifiedType string  `json:"modified_type"`
	Modified     bool    `json:"modified"`
}

func main(){
	getConfig("config")
}

func fetchData(accessToken string,date string)[]byte{
	values := url.Values{}
	values.Add("access_token", accessToken)
	values.Add("date",date)

	resp, err := http.Get("https://app.spire.io/api/v2/streaks" + "?" + values.Encode())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	//println(string(body))
	return body
}

func getConfig(filename string){
	viper.SetConfigName(filename)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("設定ファイル読み込みエラー: %s \n", err))
	}
}
