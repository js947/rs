package api

import (
	"fmt"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"net/http"
)

func Post(addr string, data io.Reader) ([]byte, error) {
	client := http.Client{}
	req, err := http.NewRequest("POST", viper.GetString("api")+addr, data)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", viper.GetString("token")))
	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if r.StatusCode < 200 || r.StatusCode >= 400 {
		return nil, fmt.Errorf("POST %s %s: %s", addr, r.Status, body)
	}

	return body, nil
}
