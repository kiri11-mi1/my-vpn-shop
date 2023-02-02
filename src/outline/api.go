package outline

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type AccessKey struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Port      int    `json:"port"`
	Method    string `json:"method"`
	AccessUrl string `json:"accessUrl"`
}

type AccessKeys []AccessKey

type KeysResponse struct {
	Keys AccessKeys `json:"accessKeys"`
}

type OutlineClient struct {
	ApiUrl string
}

func NewOutlineClient(apiUrl string) *OutlineClient {
	return &OutlineClient{
		ApiUrl: apiUrl,
	}
}

func NewTransport() *http.Transport {
	return &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
}

func (o *OutlineClient) GetKeys() (AccessKeys, error) {
	endpoint := "/access-keys"

	client := &http.Client{Transport: NewTransport()}
	resp, err := client.Get(o.ApiUrl + endpoint)
	if err != nil {
		return AccessKeys{}, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("Url:", o.ApiUrl+endpoint, "Status code:", resp.StatusCode)
		return AccessKeys{}, ErrInApi
	}
	bytesArray, err := io.ReadAll(resp.Body)
	if err != nil {
		return AccessKeys{}, nil
	}
	kr := KeysResponse{}
	if err := json.Unmarshal(bytesArray, &kr); err != nil {
		return AccessKeys{}, err
	}
	return kr.Keys, nil
}
