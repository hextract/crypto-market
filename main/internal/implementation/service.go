package implementation

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/h4x4d/crypto-market/main/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strings"
)

type MatchingEngineService interface {
	PlaceOrder(bid models.Bid) error
	CancelOrder(order_id string) error
}

type MatchingEngineStub struct {
	HTTPClient *http.Client
	BaseURL    string
}

func NewMatcingEngineStub() *MatchingEngineStub {
	_ = os.Getenv("MAIN_HOST")
	port := os.Getenv("MARKET_MAKER_PORT")
	baseURL := fmt.Sprintf("http://matching_engine:%s", port)
	return &MatchingEngineStub{
		HTTPClient: &http.Client{},
		BaseURL:    baseURL,
	}
}

func (me *MatchingEngineStub) PlaceOrder(bid models.Bid) error {
	var orderID int
	if bid.ID != nil {
		_, err := fmt.Sscanf(*bid.ID, "bid_%d", &orderID)
		if err != nil {
			return fmt.Errorf("invalid bid ID format: %v", err)
		}
	}
	var indicator = true

	if *bid.FromCurrency == "USDT" {
		indicator = false
	}

	requestBody := map[string]interface{}{
		"order_id":           orderID,
		"price_low":          *bid.MinPrice,
		"price_high":         *bid.MaxPrice,
		"amount":             *bid.AmountToBuy,
		"speed":              *bid.BuySpeed,
		"buy_sell_indicator": indicator,
		"pair":               []string{"ETH", "USDT"},
	}

	jsonBody, err := json.Marshal(requestBody)
	println(string(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %v", err)
	}
	fmt.Println(me.BaseURL + "/create-order")
	resp, err := me.HTTPClient.Post(
		me.BaseURL+"/create-order",
		"application/json",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		buf := new(strings.Builder)
		_, _ = io.Copy(buf, resp.Body)
		fmt.Println("response status:", resp.Status, "body:", buf.String())
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	logrus.Infof("Order placed successfully: ID=%d", orderID)
	return nil
}

func (me *MatchingEngineStub) CancelOrder(orderId string) error {
	var orderID int
	_, err := fmt.Sscanf(orderId, "bid_%d", &orderID)
	if err != nil {
		return fmt.Errorf("invalid bid ID format: %v", err)
	}

	requestBody := map[string]interface{}{
		"order_id": orderID,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %v", err)
	}

	resp, err := me.HTTPClient.Post(
		me.BaseURL+"/cancel-order",
		"application/json",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	logrus.Infof("Order cancelled successfully: ID=%d", orderID)
	return nil
}

type CryptoService interface {
	Encrypt(data []byte) (string, error)
	Decrypt(encrypted string) ([]byte, error)
}

type DatabaseService struct {
	pool             *pgxpool.Pool
	blockchainClient *BlockchainClient
	cryptoService    CryptoService
	logger           *logrus.Logger
}

func NewDatabaseService(connStr string, blockchainClient *BlockchainClient) (*DatabaseService, error) {
	result := new(DatabaseService)
	newPool, errPool := pgxpool.New(context.Background(), connStr)
	if (errPool != nil) || (newPool == nil) {
		return nil, errPool
	}

	cryptoService, crErr := NewLocalCryptoService()
	if crErr != nil {
		return nil, fmt.Errorf("failed to initialize crypto service: %w", crErr)
	}

	result.pool = newPool
	result.logger = logrus.New()
	result.logger.SetFormatter(&logrus.JSONFormatter{})

	result.blockchainClient = blockchainClient
	result.cryptoService = cryptoService

	return result, nil
}
