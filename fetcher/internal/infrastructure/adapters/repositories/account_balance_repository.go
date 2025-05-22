package repositories

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fetcher/internal/domain/entities"
	"fetcher/internal/domain/ports/repositories"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// do not edit this line
var interfaceBalanceSatisfiedConstraint repositories.IAccountBalanceRepository = (*AccountBalanceRepository)(nil)

type AccountBalanceRepository struct{}

const (
	BaseURL = "https://testnet.binance.vision"
)

func (repository *AccountBalanceRepository) GetBalance() (*entities.AccountBalance, error) {
	var entity entities.AccountBalance
	params := map[string]string{
		"timestamp": fmt.Sprintf("%d", time.Now().UnixMilli()),
	}

	params["signature"] = signRequest(params)
	formattedParams := url.Values{}
	for key, value := range params {
		formattedParams.Add(key, value)
	}

	req, err := http.NewRequest("GET", BaseURL+"/api/v3/account?"+formattedParams.Encode(), nil)
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

	balancesRaw, ok := result["balances"].([]interface{})
	if !ok {
		return nil, errors.New("unexpected type for balances")
	}

	for _, item := range balancesRaw {
		balMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		asset, _ := balMap["asset"].(string)
		freeStr, _ := balMap["free"].(string)
		lockedStr, _ := balMap["locked"].(string)

		freeFloat, err := strconv.ParseFloat(freeStr, 64)
		if err != nil {
			log.Fatalf("Failed to parse float: %v", err)
		}

		lockedFloat, err := strconv.ParseFloat(lockedStr, 64)
		if err != nil {
			log.Fatalf("Failed to parse float: %v", err)
		}

		switch asset {
		case "BTC":
			entity.BtcFree = freeFloat
			entity.BtcLocked = lockedFloat
		case "ETH":
			entity.EthFree = freeFloat
			entity.EthLocked = lockedFloat
		case "USDT":
			entity.UsdtFree = freeFloat
			entity.UsdtLocked = lockedFloat
		}
	}
	return &entity, nil
}

func NewAccountBalanceRepository() *AccountBalanceRepository {
	return &AccountBalanceRepository{}
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
