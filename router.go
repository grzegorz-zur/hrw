package main

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	tokenPath       = "/api/webserver/token"
	loginPath       = "/api/user/login"
	dataSwitchPath  = "/api/dialup/mobile-dataswitch"
	tokenHeader     = "__RequestVerificationToken"
	defaultUsername = "admin"
)

// Request to router.
type Request struct {
	xml.Name   `xml:"request"`
	Username   *string `xml:"Username"`
	Password   *string `xml:"Password"`
	DataSwitch *int    `xml:"dataswitch"`
}

// Response from router.
type Response struct {
	xml.Name `xml:"response"`
	Token    *string `xml:"token"`
}

var client = &http.Client{}

// Execute requests.
func Execute(url, method string, token *string, request *Request, response *Response) (int, error) {

	var reqBody io.Reader
	if request != nil {
		reqContent, err := xml.Marshal(request)
		if err != nil {
			return 0, err
		}
		reqBody = bytes.NewReader(reqContent)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return 0, err
	}
	if reqBody != nil {
		req.Header.Add("Content-Type", "text/xml; charset=UTF-8")
	}
	if token != nil {
		req.Header.Add("__RequestVerificationToken", *token)
	}

	resp, err := client.Do(req)
	if err != nil {
		if resp != nil {
			return resp.StatusCode, err
		}
		return 0, err
	}
	defer resp.Body.Close()

	respContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}

	if response != nil {
		err = xml.Unmarshal(respContent, response)
		if err != nil {
			return resp.StatusCode, err
		}
	}
	return resp.StatusCode, nil

}

// Token retrievs token from router.
func Token(router Router) (string, error) {
	url := "http://" + router.Address + tokenPath
	response := &Response{}
	code, err := Execute(url, "GET", nil, nil, response)
	if err != nil {
		return "", err
	}
	if code != http.StatusOK {
		return "", err
	}
	if response.Token == nil {
		return "", nil
	}
	return *response.Token, nil
}

// Login into router.
func Login(router Router, token string) (bool, error) {
	url := "http://" + router.Address + loginPath
	request := &Request{}
	username := defaultUsername
	request.Username = &username
	encPassword := base64.StdEncoding.EncodeToString([]byte(router.Password))
	request.Password = &encPassword
	code, err := Execute(url, "POST", &token, request, nil)
	if err != nil {
		return false, err
	}
	return code == http.StatusOK, nil
}

// SwitchData switches data on/off.
func SwitchData(router Router, token string, on bool) (bool, error) {
	url := "http://" + router.Address + dataSwitchPath
	request := &Request{}
	value := 0
	if on {
		value = 1
	}
	request.DataSwitch = &value
	code, err := Execute(url, "POST", &token, request, nil)
	if err != nil {
		return false, err
	}
	return code == http.StatusOK, nil
}
