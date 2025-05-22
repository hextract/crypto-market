#include "orders_repository.hpp"
#include <numeric>

// public

namespace Constants {
const size_t kInfId = std::numeric_limits<size_t>::max();
} // namespace Constants

OrdersRepository::OrdersRepository() :
    bid_orders_([](const ContinuousOrder& lhs, const ContinuousOrder& rhs) {
      if (lhs.GetHighPrice() == rhs.GetHighPrice()) {
        return lhs.GetOrderId() < rhs.GetOrderId();
      }
      return lhs.GetHighPrice() > rhs.GetHighPrice();
    }),
    ask_orders_([](const ContinuousOrder& lhs, const ContinuousOrder& rhs) {
      if (lhs.GetLowPrice() == rhs.GetLowPrice()) {
        return lhs.GetOrderId() < rhs.GetOrderId();
      }
      return lhs.GetLowPrice() < rhs.GetLowPrice();
    }) {}

void OrdersRepository::AddOrder(const ContinuousOrder& order) {
  switch (order.GetSide()) {
    case OrderSide::Buy:
      bid_orders_.Insert(order);
      break;
    case OrderSide::Sell:
      ask_orders_.Insert(order);
      break;
  }
}

void OrdersRepository::RemoveOrder(const ContinuousOrder& order) {
  switch (order.GetSide()) {
    case OrderSide::Buy:
      bid_orders_.Delete(order);
      break;
    case OrderSide::Sell:
      ask_orders_.Delete(order);
      break;
  }
}

std::unordered_set<ContinuousOrder> OrdersRepository::GetOrders(size_t price) {

  ContinuousOrder
      fake_bid_order = ContinuousOrder(Constants::kInfId, TradingPair(Asset(""), Asset("")), OrderSide::Buy, 0, 0, price + 1, 0);
  auto bid_orders = bid_orders_.GetTail(fake_bid_order);

  ContinuousOrder
      fake_ask_order = ContinuousOrder(Constants::kInfId, TradingPair(Asset(""), Asset("")), OrderSide::Sell, 0, price - 1, 0, 0);
  auto ask_orders = ask_orders_.GetTail(fake_ask_order);

  std::unordered_set<ContinuousOrder> orders;
  for (auto order : bid_orders) {
    if (order.GetHighPrice() >= price) {
      orders.emplace(order);
    }
  }
  for (auto order : ask_orders) {
    if (order.GetLowPrice() <= price) {
      orders.emplace(order);
    }
  }
  return orders;
}