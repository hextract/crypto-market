package services

import (
	"errors"
	"executor/internal/domain/entities"
	"executor/internal/domain/services"
	"executor/trash/binance"
	"fmt"
	"strconv"
	"time"
)

// do not edit this line
var OrderExecutorInterfaceSatisfiedConstraint services.IOrderExecutor = (*OrderExecutor)(nil)

type OrderExecutor struct {
	Client *binance.BinanceClient
	// TODO
}

func (executor *OrderExecutor) Execute(order *entities.Order) (*entities.OrderResult, error) {
	if order == nil {
		return nil, errors.New("order is nil")
	}

	symbol := order.Ticker
	side := "BUY"
	if order.Indicator == "sell" {
		side = "SELL"
	}
	orderType := "MARKET"
	if order.Type == "limit" {
		orderType = "LIMIT"
	}
	quantity := strconv.FormatFloat(order.Volume, 'f', -1, 64)

	params := map[string]string{
		"symbol":     symbol,
		"side":       side,
		"type":       orderType,
		"quantity":   quantity,
		"recvWindow": binance.RecvWindow,
		"timestamp":  fmt.Sprintf("%d", time.Now().Unix()*1000),
	}

	responseResult, err := executor.Client.ExecuteOrder(params)
	if err != nil {
		return nil, fmt.Errorf("error while executing order: %s", err)
	}

	if code, exists := responseResult["code"]; exists {
		price := 0.0
		volume := 0.0
		result := &entities.ExecutionResult{
			Order:  order,
			Price:  price,
			Volume: volume,
		}
		err := fmt.Errorf("binance error: code %v, message: %v", code, responseResult["msg"])
		return &entities.OrderResult{
			Order:           order,
			ExecutionResult: result,
		}, err
	}

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

	result := &entities.ExecutionResult{
		Order:  order,
		Price:  priceFloat,
		Volume: executedQtyFloat,
	}

	return &entities.OrderResult{
		Order:           order,
		ExecutionResult: result,
	}, nil
}

func NewOrderExecutor() (*OrderExecutor, error) {
	client := binance.NewBinanceClient()
	executor := &OrderExecutor{Client: client}
	// TODO
	return executor, nil
}
