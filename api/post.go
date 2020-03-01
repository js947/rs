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
	req, err := http.NewRequest("POST", addr, data)
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
	if r.StatusCode != 201 {
		return nil, fmt.Errorf("upload error %s: %s", r.Status, body)
	}

	return body, nil
}
