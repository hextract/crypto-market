#include "linear.hpp"
#include "unordered_set"

#include "domain/continuous_order.hpp"
#include "domain/config.hpp"


LinearUserOrderBookRepo::LinearUserOrderBookRepo(const TradingPair &pair) : pair_(pair) {
  book_ = std::vector<std::unordered_set<ContinuousOrder>>(1000000);
}

std::unordered_set<ContinuousOrder> LinearUserOrderBookRepo::GetOrders(size_t price) {
  return book_[price];
}

void LinearUserOrderBookRepo::AddOrder(const ContinuousOrder &order) {
  for (size_t i = order.GetLowPrice(); i <= order.GetHighPrice(); ++i) {
    book_[i].insert(order);
  }
  if (order.GetSide() == OrderSide::Buy) {
    for (size_t i = 0; i < order.GetLowPrice(); ++i) {
      book_[i].insert(order);
    }
  } else {
    for (size_t i = order.GetHighPrice() + 1; i < book_.size(); ++i) {
      book_[i].insert(order);
    }
  }
}

void LinearUserOrderBookRepo::RemoveOrder(const ContinuousOrder &order) {
  for (size_t i = order.GetLowPrice(); i <= order.GetHighPrice(); ++i) {
    if (!book_[i].contains(order)) {
      break;
    }
    book_[i].erase(order);
  }
  if (order.GetSide() == OrderSide::Buy) {
    for (size_t i = 0; i < order.GetLowPrice(); ++i) {
      if (!book_[i].contains(order)) {
        break;
      }
      book_[i].erase(order);
    }
  } else {
    for (size_t i = order.GetHighPrice() + 1; i < book_.size(); ++i) {
      if (!book_[i].contains(order)) {
        break;
      }
      book_[i].erase(order);
    }
  }
}
