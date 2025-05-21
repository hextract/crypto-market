#pragma once

#include <iostream>
#include <cmath>
#include "trading_pair.hpp"

enum class OrderSide {
 Buy,
 Sell,
};

class ContinuousOrder {
 public:
  ContinuousOrder(size_t order_id, const TradingPair& pair, const OrderSide& side, size_t volume,
                 size_t low_price, size_t high_price, size_t speed)
    : order_id_(order_id), pair_(pair), side_(side), volume_(volume), low_price_(low_price),
      high_price_(high_price), speed_(speed) {}

  [[nodiscard]] size_t GetOrderId() const { return order_id_; }

  [[nodiscard]] TradingPair GetPair() const { return pair_; }

  [[nodiscard]] OrderSide GetSide() const { return side_; }

  [[nodiscard]] size_t GetVolume() const { return volume_; }

  [[nodiscard]] size_t GetLowPrice() const { return low_price_; }

  [[nodiscard]] size_t GetHighPrice() const { return high_price_; }

  [[nodiscard]] size_t GetSpeed() const { return speed_; }

  [[nodiscard]] size_t GetRoundedSpeed(size_t price) const {
    if (side_ == OrderSide::Buy) {
      if (price <= low_price_) {
        return speed_;
      }
      if (price >= high_price_) {
        return 0;
      }
      double d_price = double(high_price_ - price) / double(high_price_ - low_price_) * double(speed_);
      return std::floor(d_price);
    }
    if (price >= high_price_) {
      return speed_;
    }
    if (price <= low_price_) {
      return 0;
    }
    double d_price = double(price - low_price_) / double(high_price_ - low_price_) * double(speed_);
    return std::floor(d_price);
  }

  ContinuousOrder& operator=(const ContinuousOrder& other_order) = default;

  ContinuousOrder(const ContinuousOrder& other_order) = default;

 private:
  size_t order_id_;
  TradingPair pair_;
  OrderSide side_;
  size_t volume_;
  size_t low_price_;
  size_t high_price_;
  size_t speed_;
};

inline bool operator==(const ContinuousOrder& lhs, const ContinuousOrder& rhs) {
 return lhs.GetOrderId() == rhs.GetOrderId();
}

template<>
struct std::hash<ContinuousOrder> {
 size_t operator()(const ContinuousOrder& order) const {
  return std::hash<std::size_t>{}(order.GetOrderId());
 }
};

class CancelOrder {
 public:
  explicit CancelOrder(size_t order_id) : order_id_(order_id) {}

  [[nodiscard]] size_t GetOrderId() const { return order_id_; }
 private:
  size_t order_id_{0};
};
