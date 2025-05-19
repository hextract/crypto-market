#include <memory>

#include "matching_engine.hpp"
#include "usecases/cancel_queue/cancel_queue.hpp"
#include "usecases/order_queue/order_queue.hpp"
#include "usecases/order_book/order_book.hpp"
#include "usecases/user_order_book/user_order_book.hpp"
#include "usecases/price_calculator/segment_treap/segment_treap_price_calculator.hpp"
#include "pkg/repository/order_book/linear/linear.hpp"
#include "pkg/repository/user_order_book/linear/linear.hpp"

#include "pkg/repository/order_book/optimal/ask_curve_repository.hpp"
#include "pkg/repository/order_book/optimal/bid_curve_repository.hpp"
#include "pkg/repository/user_order_book/optimal/orders_repository.hpp"

MatchingEngineService::MatchingEngineService() {
  TradingPair pair(Asset("ETH"), Asset("USDT"));
  order_book_config_ = std::make_shared<OrderBookConfiguration>();
  order_book_config_->ChangeTickSize(pair, 1000);
  order_book_config_->ChangeMinPartCnt(pair, 1000000000);

  auto user_order_book_repo = std::make_shared<OrdersRepository>();
  auto buy_order_book_repo = std::make_shared<BidCurveRepository>();
  auto sell_order_book_repo = std::make_shared<AskCurveRepository>();
  auto user_order_book = std::make_shared<UserOrderBook>(pair, user_order_book_repo);
  auto buy_order_book = std::make_shared<OrderBook>(buy_order_book_repo);
  auto sell_order_book = std::make_shared<OrderBook>(sell_order_book_repo);
  auto price_calc = std::make_shared<SegmentTreapPriceCalculator>(buy_order_book_repo, sell_order_book_repo);

  auto order_queue = std::make_shared<OrderQueue>();
  auto cancel_queue = std::make_shared<CancelQueue>();
  std::function<void(const ContinuousOrder&)> create_cb =
      [order_queue](const ContinuousOrder& order) {
        order_queue->AddOrder(order);
      };

  std::function<void(const CancelOrder&)> cancel_cb =
      [cancel_queue](const CancelOrder& order) {
        cancel_queue->AddOrder(order);
      };

  std::function<std::optional<size_t>()> get_clearing_price_handler = [price_calc]() {
    return price_calc->CalculatePrice();
  };

  std::function< std::vector<CurveFracture>(size_t, size_t)>
  get_ask_curve_handler = [sell_order_book_repo](size_t left_boundary_price, size_t right_boundary_price) {
    return sell_order_book_repo->GetCurve(left_boundary_price, right_boundary_price);
  };

  std::function< std::vector<CurveFracture>(size_t, size_t)>
      get_bid_curve_handler = [buy_order_book_repo](size_t left_boundary_price, size_t right_boundary_price) {
    return buy_order_book_repo->GetCurve(left_boundary_price, right_boundary_price);
  };

  http_server_ = std::make_shared<HTTPServer>(order_book_config_);

  http_server_->SetCreateOrderCallback(create_cb);
  http_server_->SetCancelOrderCallback(cancel_cb);
  http_server_->SetGetAskCurveHandler(get_ask_curve_handler);
  http_server_->SetGetBidCurveHandler(get_bid_curve_handler);
  http_server_->SetGetClearingPriceHandler(get_clearing_price_handler);
  http_server_->SetPair(pair);

  engine_manager_ =
      std::make_shared<EngineManager>(order_queue, user_order_book, buy_order_book, sell_order_book, price_calc);
}

void MatchingEngineService::Run() {
  std::thread th_server([this]() { http_server_->Run(); });
  std::thread th_manager([this]() { engine_manager_->Run(); });
  th_server.join();
  th_manager.join();
}
