#include "internal/domain/asset.hpp"
#include "internal/domain/trading_pair.hpp"
#include "internal/domain/continuous_order.hpp"
#include "internal/pkg/app/matching_engine.hpp"
#include "internal/pkg/repository/user_order_book/linear/linear.hpp"
#include "internal/pkg/repository/order_book/linear/linear.hpp"
#include "internal/usecases/user_order_book/user_order_book.hpp"
#include "internal/usecases/order_book/order_book.hpp"
#include "internal/usecases/price_calculator/linear/linear_price_calculator.hpp"
#include "internal/usecases/engine_manager/engine_manager.hpp"
#include "internal/usecases/order_book_configuration/order_book_configuration.hpp"
#include "internal/usecases/cancel_queue/cancel_queue.hpp"
#include "internal/pkg/app/matching_engine.hpp"

#include <fstream>
#include <sstream>
#include <cmath>

#include "pkg/app/matching_engine.hpp"
#include "pkg/transport/client/exchange_connector.hpp"

std::vector<ContinuousOrder> readOrdersFromCSV(const std::string &filename) {
  std::vector<ContinuousOrder> orders;
  std::ifstream file(filename);
  if (!file.is_open()) {
    std::cerr << "Ошибка при открытии файла " << filename << std::endl;
    return orders;
  }

  std::string line;
  std::getline(file, line);
  size_t ind = 2;

  TradingPair pair(Asset("ETH"), Asset("USDT"));
  while (std::getline(file, line)) {
    std::istringstream ss(line);
    std::string orderSide, priceStr, sizeStr;

    std::getline(ss, orderSide, ';');
    std::getline(ss, priceStr, ';');
    std::getline(ss, sizeStr, ';');

    size_t price = std::floor(std::stod(priceStr) * 100);
    size_t size = std::floor(std::stod(sizeStr) * 1000);

    ContinuousOrder order{
      ind, pair, orderSide == "buy" ? OrderSide::Buy : OrderSide::Sell,
      size, price - 20, price + 20, 30
    };
    orders.push_back(order);
    ++ind;
  }

  file.close();
  return orders;
}


void test() {
  TradingPair pair(Asset("BTC"), Asset("USDT"));
  auto user_order_book_repo = std::make_shared<LinearUserOrderBookRepo>(pair);
  auto buy_order_book_repo = std::make_shared<LinearOrderBookRepo>(1000000);
  auto sell_order_book_repo = std::make_shared<LinearOrderBookRepo>(1000000);
  auto user_order_book = std::make_shared<UserOrderBook>(pair, user_order_book_repo);
  auto buy_order_book = std::make_shared<OrderBook>(buy_order_book_repo);
  auto sell_order_book = std::make_shared<OrderBook>(sell_order_book_repo);
  auto price_calc = std::make_shared<LinearPriceCalculator>(buy_order_book_repo, sell_order_book_repo, 1000000);
  auto order_queue = std::make_shared<OrderQueue>();

  EngineManager engine_manager{order_queue, user_order_book, buy_order_book, sell_order_book, price_calc};

  // ContinuousOrder order_1(1, pair, OrderSide::Buy, 15, 15, 15, 5);
  // ContinuousOrder order_2(2, pair, OrderSide::Sell, 15, 15, 15, 5);

  auto orders = readOrdersFromCSV("./orders.csv");
  for (const auto &i: orders) {
    order_queue->AddOrder(i);
  }

  for (int i = 0; i < 100; ++i) {
    engine_manager.Step();
  }
}

void run() {
  TradingPair pair(Asset("ETH"), Asset("USDT"));
  std::shared_ptr<OrderBookConfiguration> config_ = std::make_shared<OrderBookConfiguration>();
  config_->ChangeTickSize(pair, 1000);
  config_->ChangeMinPartCnt(pair, 1000000000);
  auto order_queue = std::make_shared<OrderQueue>();
  auto cancel_queue = std::make_shared<CancelQueue>();

  std::function<void(const ContinuousOrder &)> create_cb =
      [&order_queue](const ContinuousOrder &order) {
    order_queue->AddOrder(order);
  };

  std::function<void(const CancelOrder &)> cancel_cb =
      [&cancel_queue](const CancelOrder &order) {
        cancel_queue->AddOrder(order);
  };

  HTTPServer server(config_);
  server.SetCreateOrderCallback(create_cb);
  server.SetCancelOrderCallback(cancel_cb);
  server.Run();
}

int main() {
  Config::Init("./.env");
  std::shared_ptr<MatchingEngineService> service = std::make_shared<MatchingEngineService>();
  std::thread service_worker([service](){ service->Run(); });
  std::thread drogon_worker([](){
    std::cout << "directly before drogon app run" << std::endl;
    drogon::app().run();
  });
  service_worker.join();
  drogon_worker.join();

  // run();
  // ExchangeConnectorClient client;
  // std::function<void(const Success &)> success_cb =
  // [](const Success &success) {
  // std::cout << "success placing order " << success.message << "\n";
  // };
  //
  // std::function<void(const Error &)> error_cb =
  // [](const Error &error) {
  // std::cout << "error placing order " << error.message << "\n";
  // };
  // LimitExchangeOrder order{5, "market", "by", 10, 10, "ETHUSDT"};
  // client.PlaceOrder(order, success_cb, error_cb);
  // drogon::app().run();
}
