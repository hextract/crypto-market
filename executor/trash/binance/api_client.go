package binance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

type BinanceClient struct {
	BaseURL string
}

func NewBinanceClient() *BinanceClient {
	return &BinanceClient{
		BaseURL: baseURL,
	}
}

func (bc *BinanceClient) ExecuteOrder(params map[string]string) (map[string]interface{}, error) {
	params["timestamp"] = fmt.Sprintf("%d", time.Now().Unix()*1000)
	params["signature"] = signRequest(params)

	formattedParams := url.Values{}
	for key, value := range params {
		formattedParams.Add(key, value)
	}

	req, err := http.NewRequest("POST", bc.BaseURL+"/api/v3/order", bytes.NewReader([]byte(formattedParams.Encode())))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("X-MBX-APIKEY", os.Getenv("API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var responseResult map[string]interface{}
	err = json.Unmarshal(body, &responseResult)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return responseResult, nil
}

func (bc *BinanceClient) GetBalance() (map[string]interface{}, error) {
	params := map[string]string{
		"timestamp": fmt.Sprintf("%d", time.Now().UnixMilli()),
	}

	params["signature"] = signRequest(params)
	formattedParams := url.Values{}
	for key, value := range params {
		formattedParams.Add(key, value)
	}

	req, err := http.NewRequest("GET", bc.BaseURL+"/api/v3/account?"+formattedParams.Encode(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("X-MBX-APIKEY", os.Getenv("API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return result, nil
}
