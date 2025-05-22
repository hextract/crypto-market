#pragma once

#include <unordered_set>

#include "domain/continuous_order.hpp"

class IUserOrderBookRepo {
 public:
  virtual void AddOrder(const ContinuousOrder& order) = 0;

  virtual std::unordered_set<ContinuousOrder> GetOrders(size_t price) = 0;

  // TODO: think more about price parameter, may be changed
  virtual void RemoveOrder(const ContinuousOrder& order) = 0;

  virtual ~IUserOrderBookRepo() = default;
};