package orders_executor

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"executor/internal/binance"
	"executor/internal/models"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ResultCallback func(result *models.Result) error

func (oe *OrdersExecutor) ExecuteOrder(order *models.Order, resultCallback ResultCallback) error {
	if order == nil {
		return errors.New("order is nil")
	}

	symbol := *order.Ticker
	side := "BUY"
	if *order.Indicator == "sell" {
		side = "SELL"
	}
	orderType := "MARKET"
	if *order.Type == "limit" {
		orderType = "LIMIT"
	}
	quantity := strconv.FormatFloat(*order.Volume, 'f', -1, 64)

	params := map[string]string{
		"symbol":     symbol,
		"side":       side,
		"type":       orderType,
		"quantity":   quantity,
		"recvWindow": binance.RecvWindow,
		"timestamp":  fmt.Sprintf("%d", time.Now().Unix()*1000),
	}

	//signature := signRequest(params)
	//params["signature"] = signature
	//
	//formattedParams := url.Values{}
	//for key, value := range params {
	//	formattedParams.Add(key, value)
	//}

	responseResult, err := oe.Client.ExecuteOrder(params)
	if err != nil {
		return fmt.Errorf("error while executing order: %s", err)
	}

	if code, exists := responseResult["code"]; exists {
		price := 0.0
		volume := 0.0
		result := &models.Result{
			Order:  order,
			Price:  &price,
			Volume: &volume,
		}
		_ = fmt.Errorf("binance error: code %v, message: %v", code, responseResult["msg"])
		return resultCallback(result)
	}

	//arrayFills := binance.ParseFills(responseResult)
	//fmt.Println(arrayFills)

	//req, err := http.NewRequest("POST", baseURL+"/api/v3/order", strings.NewReader(formattedParams.Encode()))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//req.Header.Set("X-MBX-APIKEY", os.Getenv("API_KEY"))
	//
	//client := &http.Client{}
	//resp, err := client.Do(req)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer resp.Body.Close()
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	_ = fmt.Errorf("error when receiving the response")
	//}
	//var responseResult map[string]interface{}
	//err = json.Unmarshal(body, &responseResult)
	//if err != nil {
	//	_ = fmt.Errorf("error processing the response")
	//}

	executedQty := responseResult["executedQty"].(string)
	if executedQty == "" {
		_ = fmt.Errorf("empty executed quantity")
	}

	fills, ok := responseResult["fills"].([]interface{})
	if !ok || len(fills) == 0 {
		fmt.Println("array fills is empty or missing")
	}

	firstFill := fills[0].(map[string]interface{})
	price := firstFill["price"].(string)

	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Println("Error during conversion price:", err)
	}

	executedQtyFloat, err := strconv.ParseFloat(executedQty, 64)
	if err != nil {
		fmt.Println("Error during conversion executedQty:", err)
	}

	result := &models.Result{
		Order:  order,
		Price:  &priceFloat,
		Volume: &executedQtyFloat,
	}

	// realization started
	//const maxVolumeProceeded = 30.0
	//// simulating work
	//volumeProceeded := min(*order.Volume, maxVolumeProceeded)
	//result := &models.Result{
	//	Order:  order,
	//	Price:  order.Price,
	//	Volume: &volumeProceeded,
	//}
	//time.Sleep(time.Second * 2)
	// realization ended

	return resultCallback(result)
}

func signRequest(params map[string]string) string {
	var keys []string
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var queryString string
	for _, key := range keys {
		queryString += key + "=" + params[key] + "&"
	}
	queryString = strings.TrimRight(queryString, "&")

	mac := hmac.New(sha256.New, []byte(os.Getenv("API_SECRET")))
	mac.Write([]byte(queryString))
	return hex.EncodeToString(mac.Sum(nil))
}
