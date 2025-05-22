package services

import (
	"context"
	"encoding/json"
	"fetcher/internal/application/use_cases"
	"fetcher/internal/domain/entities"
	"fetcher/internal/domain/ports/repositories"
	"fetcher/internal/domain/services"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math"
	"strconv"
	"time"
)

// do not edit this line
var iExchangePollerSatisfiedConstraint services.IExchangePoller = (*BinancePoller)(nil)

type BinancePoller struct {
	receiveOrderBookSnapshotUseCase *use_cases.ReceiveOrderBookSnapshotUseCase

	// TODO

}

func (binancePoller *BinancePoller) PollOrderBook(context context.Context) error {
	// When you got an OrderBookSnapshot, create appropriate UseCase and call ReceiveOrderBookSnapshotUseCase.Execute(OrderBookSnapshot)
	go binancePoller.startWebSocket(context)
	// TODO
	return nil
}

func NewBinancePoller(orderBookRepository repositories.IOrderBookRepository) (*BinancePoller, error) {

	// TODO

	receiveOrderBookSnapshotUseCase, errUseCase := use_cases.NewReceiveOrderBookSnapshotUseCase(orderBookRepository)
	if errUseCase != nil {
		return nil, errUseCase
	}
	return &BinancePoller{receiveOrderBookSnapshotUseCase: receiveOrderBookSnapshotUseCase}, nil
}

func toFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func sumVolume(levels [][]string, depth int) float64 {
	var total float64
	for i := 0; i < depth && i < len(levels); i++ {
		total += toFloat(levels[i][1])
	}
	return total
}

func calculateImbalance(bids, asks [][]string, depth int) float64 {
	bidVol := sumVolume(bids, depth)
	askVol := sumVolume(asks, depth)
	if bidVol+askVol == 0 {
		return 0
	}
	return (bidVol - askVol) / (bidVol + askVol)
}

func calculateVWAP(levels [][]string, depth int) float64 {
	var volSum, volPriceSum float64
	for i := 0; i < depth && i < len(levels); i++ {
		price := toFloat(levels[i][0])
		vol := toFloat(levels[i][1])
		volSum += vol
		volPriceSum += price * vol
	}
	if volSum == 0 {
		return 0
	}
	return volPriceSum / volSum
}

func calculateLiquidity(levels [][]string, mid float64, pct float64) float64 {
	var total float64
	for _, level := range levels {
		price := toFloat(level[0])
		vol := toFloat(level[1])
		if math.Abs(price-mid)/mid <= pct/100 {
			total += vol
		}
	}
	return total
}

func (binancePoller *BinancePoller) startWebSocket(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Остановка WebSocket по контексту")
			return
		default:
			conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
			if err != nil {
				log.Println("Ошибка подключения, повтор через 5 секунд:", err)
				time.Sleep(5 * time.Second)
				continue
			}
			log.Println("Подключено к Binance WebSocket")

			conn.SetPongHandler(func(appData string) error {
				log.Println("Pong от сервера")
				return nil
			})

			go func(c *websocket.Conn) {
				ticker := time.NewTicker(pingDelay)
				defer ticker.Stop()
				for {
					select {
					case <-ticker.C:
						err := c.WriteMessage(websocket.PingMessage, []byte("ping"))
						if err != nil {
							log.Println("Ошибка при отправке ping:", err)
							_ = c.Close()
							return
						}
						log.Println("Ping отправлен")
					case <-ctx.Done():
						log.Println("Остановка пинговой горутины")
						return
					}
				}
			}(conn)

			for {
				_, message, err := conn.ReadMessage()
				if err != nil {
					log.Println("🔌 Потеря соединения:", err)
					break
				}

				var update DepthUpdate
				err = json.Unmarshal(message, &update)
				if err != nil {
					log.Println("Ошибка парсинга:", err)
					continue
				}

				snapshot := binancePoller.processOrderBook(update)
				errWithExecute := binancePoller.receiveOrderBookSnapshotUseCase.Execute(&snapshot)
				if errWithExecute != nil {
					_ = fmt.Errorf("error when sending data")
				}
			}

			_ = conn.Close()
			log.Println("Переподключение через 5 секунд...")
			time.Sleep(5 * time.Second)
		}
	}
}

func (binancePoller *BinancePoller) processOrderBook(update DepthUpdate) entities.OrderBookSnapshot {
	if len(update.Bids) == 0 || len(update.Asks) == 0 {
		return entities.OrderBookSnapshot{}
	}

	bestBid := toFloat(update.Bids[0][0])
	bestAsk := toFloat(update.Asks[0][0])
	mid := (bestBid + bestAsk) / 2
	spread := bestAsk - bestBid

	bidVol := sumVolume(update.Bids, depth)
	askVol := sumVolume(update.Asks, depth)
	imbalance := calculateImbalance(update.Bids, update.Asks, depth)
	vwapBid := calculateVWAP(update.Bids, depth)
	vwapAsk := calculateVWAP(update.Asks, depth)
	liqBid := calculateLiquidity(update.Bids, mid, 1.0)
	liqAsk := calculateLiquidity(update.Asks, mid, 1.0)

	return entities.OrderBookSnapshot{
		Timestamp:  time.Now().Unix(),
		Depth:      depth,
		Mid:        mid,
		Spread:     spread,
		BidVolume:  bidVol,
		AskVolume:  askVol,
		BestBid:    bestBid,
		BestAsk:    bestAsk,
		VWAPBid:    vwapBid,
		VWAPAsk:    vwapAsk,
		Imbalance:  imbalance,
		LiqBid1Pct: liqBid,
		LiqAsk1Pct: liqAsk,
	}
}

type DepthUpdate struct {
	EventType string     `json:"e"`
	Time      int64      `json:"E"`
	Symbol    string     `json:"s"`
	Bids      [][]string `json:"b"`
	Asks      [][]string `json:"a"`
}

const (
	symbol    = "btcusdt"
	wsURL     = "wss://stream.binance.com:9443/ws/" + symbol + "@depth"
	depth     = 10
	pingDelay = 3 * time.Minute
)
