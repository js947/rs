package api

import (
	"io/ioutil"
	"log"
	"fmt"
	"net/http"
	"github.com/spf13/viper"
)

func Get(addr string) []byte {
	client := http.Client{}

	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", viper.GetString("token")))

	r, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}
