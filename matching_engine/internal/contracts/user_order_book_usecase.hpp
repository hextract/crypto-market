#pragma once

#include "domain/continuous_order.hpp"
#include "domain/matched_details.hpp"

class IUserOrderBook {
 public:
  virtual void AddOrder(const ContinuousOrder& order) = 0;

  virtual void RemoveOrder(const ContinuousOrder& order) = 0;

  virtual MatchedDetails MatchOrders(size_t clearing_price) = 0;

  virtual ~IUserOrderBook() = default;
};