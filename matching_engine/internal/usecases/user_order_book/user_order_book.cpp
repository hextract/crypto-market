#include "user_order_book.hpp"
#include "domain/continuous_order.hpp"
#include "domain/matched_details.hpp"


MatchedDetails UserOrderBook::MatchOrders(size_t price) {
  std::unordered_set<ContinuousOrder> orders = storage_->GetOrders(price);
  std::cout << "orders count - " << orders.size() <<  std::endl;
  std::vector<ContinuousOrder> buy_orders;
  std::vector<ContinuousOrder> sell_orders;
  MatchedDetails matched_details;
  for (const auto& order : orders) {
    if (order.GetSide() == OrderSide::Buy) {
      buy_orders.push_back(order);
    } else {
      sell_orders.push_back(order);
    }
  }
  int64_t imbalance = 0;
  for (const auto& order: buy_orders) {
    size_t cur_speed = order.GetRoundedSpeed(price);
    size_t need_buy = cur_speed * delta_;
    size_t remains = order.GetVolume() - account_manager_.GetBaseCoinFilled(order);
    if (need_buy >= remains) {
      imbalance -= static_cast<int64_t>(need_buy - remains);
      if (order.GetPair().GetBase().GetName() == "") {
        throw std::runtime_error("");
      }
      matched_details.AddBuyFilled(order);
      storage_->RemoveOrder(order);
    }
    account_manager_.AddFilledQuantity(order, price, std::min(need_buy, remains));
    backend_client_->SendUpdate(order, account_manager_.GetFillDetails(order));
  }
  for (const auto& order: sell_orders) {
    size_t cur_speed = order.GetRoundedSpeed(price);
    size_t need_sell = cur_speed * delta_;
    size_t remains = order.GetVolume() - account_manager_.GetBaseCoinFilled(order);
    if (need_sell >= remains) {
      imbalance += static_cast<int64_t>(need_sell - remains);
      account_manager_.AddFilledQuantity(order, price, remains);
      if (order.GetPair().GetBase().GetName() == "") {
        throw std::runtime_error("");
      }
      matched_details.AddSellFilled(order);
      storage_->RemoveOrder(order);
    } else {
      account_manager_.AddFilledQuantity(order, price, need_sell);
    }
    backend_client_->SendUpdate(order, account_manager_.GetFillDetails(order));
  }
  matched_details.AddImbalance(imbalance);
  return matched_details;
}

void UserOrderBook::SetBackendClient(const std::shared_ptr<IBackendClient>& client) {
  backend_client_ = client;
}

void UserOrderBook::AddOrder(const ContinuousOrder &order) {
  storage_->AddOrder(order);
}

void UserOrderBook::RemoveOrder(const ContinuousOrder &order) {
  storage_->RemoveOrder(order);
}

