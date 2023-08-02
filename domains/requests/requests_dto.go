package requests

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/vakil-bashi/log-system/logger"
	"github.com/vakil-bashi/log-system/utils/responses"
	"io/ioutil"
	"net/http"
	"strings"
)

type Req struct {
	URL    string `json:"url"`
	Method string `json:"method"`
}

func (r *Req) Init(url string, method string) {
	r.URL = url
	r.Method = method
}

func (r *Req) POST(jsonMap map[string]interface{}, bearerToken *string, simpleAuthentication *bool) ([]byte, *responses.Response) {

	jsonBody, _ := json.Marshal(jsonMap)
	payload := strings.NewReader(string(jsonBody))

	client := &http.Client{}
	req, err := http.NewRequest(r.Method, r.URL, payload)

	if err != nil {
		fmt.Println(err)
		logger.Error("Could not send request to es", err)
		return nil, responses.NewBadRequestError("Could not send request to es", "please try again later...", http.StatusBadRequest)
	}
	req.Header.Add("Content-Type", "application/json")

	if simpleAuthentication != nil {
		authHeader := fmt.Sprintf("%s:%s", "elastic", "2t5GR=RfKcTlWQaLgokD")
		authHeaderValue := "Basic " + base64.StdEncoding.EncodeToString([]byte(authHeader))
		req.Header.Set("Authorization", authHeaderValue)
	}

	if bearerToken != nil {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *bearerToken))
	}

	res, err := client.Do(req)
	if err != nil {
		logger.Error("Could not send request to es", err)
		return nil, responses.NewBadRequestError("Could not send request to es", "please try again later...", http.StatusBadRequest)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error("Could not send request to es", err)
		return nil, responses.NewBadRequestError("Could not send request to es", "please try again later...", http.StatusBadRequest)
	}
	return body, nil
}
