package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type FileInfo struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func UploadFile(name string, data *bytes.Buffer) (*FileInfo, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, err := w.CreateFormFile("file", name)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fw, data)
	if err != nil {
		return nil, err
	}
	w.Close()

	client := http.Client{}
	req, err := http.NewRequest("POST", "https://platform.rescale.com/api/v2/files/contents/", &b)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", viper.GetString("token")))
	req.Header.Add("Content-Type", w.FormDataContentType())

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
		var errorresponse struct {
			Detail string `json:"detail"`
		}
		json.Unmarshal(body, &errorresponse)
		return nil, fmt.Errorf("upload error %s: %s", r.Status, errorresponse.Detail)
	}

	var fileinfo FileInfo
	json.Unmarshal(body, &fileinfo)
	return &fileinfo, nil
}
