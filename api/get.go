package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

func do(addr string, verb string) ([]byte, error) {
	client := http.Client{}

	req, err := http.NewRequest(verb, addr, nil)
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

func Get(addr string) ([]byte, error) {
	return do(addr, "GET")
}
func Delete(addr string) ([]byte, error) {
	return do(addr, "DELETE")
}
