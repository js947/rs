package api

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

func Get(addr string) ([]byte, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", viper.GetString("token")))

	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
