package outline

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
)

type AccessKey struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Port      string `json:"port"`
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

func (o *OutlineClient) GetKeys() (AccessKeys, error) {
	endpoint := "/access-keys"

	req := &fasthttp.Request{}
	res := &fasthttp.Response{}

	req.Header.SetMethod(http.MethodGet)
	req.SetRequestURI(o.ApiUrl + endpoint)

	err := fasthttp.Do(req, res)
	if err != nil {
		return AccessKeys{}, err
	}
	keys := AccessKeys{}
	if res.StatusCode() != http.StatusOK {
		return AccessKeys{}, errors.New(fmt.Sprintf(API_ERROR_MESSAGE, res.StatusCode(), res.Body()))
	}
	if err := json.Unmarshal(res.Body(), &keys); err != nil {
		return AccessKeys{}, err
	}
	return keys, nil
}
