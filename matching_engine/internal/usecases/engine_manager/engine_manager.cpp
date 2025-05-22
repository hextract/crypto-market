#include "engine_manager.hpp"

#include <vector>
#include <thread>
#include <chrono>

#include "domain/continuous_order.hpp"
#include "domain/matched_details.hpp"


void EngineManager::CreateOrdersStep() {
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
    order_storage_.insert(std::make_pair(order.GetOrderId(), order));
  }
  auto clearing_price_opt = price_calculator_->CalculatePrice();
  if (!clearing_price_opt.has_value()) {
    return;
  }
  size_t clearing_price = clearing_price_opt.value();
  MatchedDetails matched_details = user_order_book_->MatchOrders(clearing_price);
  price_calculator_->ChangeImbalance(matched_details.GetImbalance());

  std::vector<ContinuousOrder> buy_filled = std::move(matched_details.GetBuyFilled());
  std::vector<ContinuousOrder> sell_filled = std::move(matched_details.GetSellFilled());

  for (const auto& order : buy_filled) {
    order_storage_.erase(order.GetOrderId());
  }

  for (const auto& order : sell_filled) {
    order_storage_.erase(order.GetOrderId());
  }

  for (const auto& buy_order : buy_filled) {
    buy_order_book_->DeleteOrder(buy_order);
  }
  for (const auto& sell_order : sell_filled) {
    sell_order_book_->DeleteOrder(sell_order);
  }
}

void EngineManager::CancelOrdersStep() {
  auto start_time = std::chrono::steady_clock::now();
  auto total_dt = std::chrono::milliseconds(time_delta_);
  while (std::chrono::steady_clock::now() - start_time < total_dt) {
    auto order_opt = cancel_queue_->Pop();
    if (!order_opt.has_value()) {
      break;
    }
    CancelOrder cancel_order = order_opt.value();
    if (!order_storage_.contains(cancel_order.GetOrderId())) {
      std::cout << "There is no order with id " << cancel_order.GetOrderId() << " to cancel" << std::endl;
      return;
    }
    ContinuousOrder original_order = order_storage_.at(cancel_order.GetOrderId());
    if (original_order.GetSide() == OrderSide::Buy) {
      buy_order_book_->DeleteOrder(original_order);
    } else {
      sell_order_book_->DeleteOrder(original_order);
    }
    user_order_book_->CancellOrder(original_order);
  }
}

void EngineManager::RunCreateOrders() {
  while (true) {
    CreateOrdersStep();
  }
}

void EngineManager::RunCancelOrders() {
  while (true) {
    CancelOrdersStep();
  }
}

void EngineManager::Run() {
  std::thread worker_create([this]() {
    RunCreateOrders();
  });
  std::thread worker_cancel([this]() {
    RunCancelOrders();
  });
  worker_create.join();
  worker_cancel.join();
}



