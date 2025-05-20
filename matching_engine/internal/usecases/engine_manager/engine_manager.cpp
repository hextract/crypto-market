#include "engine_manager.hpp"

#include <vector>

#include "domain/continuous_order.hpp"
#include "domain/matched_details.hpp"


void EngineManager::Step() {
  auto start_time = std::chrono::steady_clock::now();
  auto total_dt = std::chrono::milliseconds(time_delta_);
  auto matching_margin = std::chrono::milliseconds(time_delta_ / 2);

  while (std::chrono::steady_clock::now() - start_time < total_dt - matching_margin) {
    auto order_opt = order_queue_->Pop();
    if (!order_opt.has_value()) {
      break;
    }
    ContinuousOrder order = order_opt.value();
    if (order.GetSide() == OrderSide::Buy) {
      buy_order_book_->AddOrder(order);
    } else {
      sell_order_book_->AddOrder(order);
    }
    user_order_book_->AddOrder(order);
  }
  auto clearing_price_opt = price_calculator_->CalculatePrice();
  if (!clearing_price_opt.has_value()) {
    return;
  }
  size_t clearing_price = clearing_price_opt.value();
//  std::cout << "Clearing price - " << clearing_price << std::endl;
  MatchedDetails matched_details = user_order_book_->MatchOrders(clearing_price);
  price_calculator_->ChangeImbalance(matched_details.GetImbalance());

  std::vector<ContinuousOrder> buy_filled = std::move(matched_details.GetBuyFilled());
  std::vector<ContinuousOrder> sell_filled = std::move(matched_details.GetSellFilled());

//  std::cout << "step was done!\n";
//  std::cout << "filled buy orders:\n";
  if (buy_filled.empty()) {
//    std::cout << "no buy orders filled!";
  } else {
    for (const auto& buy_order : buy_filled) {
//      std::cout << buy_order.GetOrderId() << " ";
    }
  }
//  std::cout << "\nfilled sell orders:\n";
  if (sell_filled.empty()) {
//    std::cout << "no sell orders filled!";
  } else {
    for (const auto& sell_order : sell_filled) {
//      std::cout << sell_order.GetOrderId() << " ";
    }
  }
//  std::cout << '\n';

  for (const auto& buy_order : buy_filled) {
    buy_order_book_->DeleteOrder(buy_order);
  }
  for (const auto& sell_order : sell_filled) {
    sell_order_book_->DeleteOrder(sell_order);
  }
}

void EngineManager::Run() {
  while (true) {
    Step();
  }
}



