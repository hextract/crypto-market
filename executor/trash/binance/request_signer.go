package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"sort"
	"strings"
)

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
