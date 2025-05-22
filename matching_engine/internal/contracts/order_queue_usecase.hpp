#pragma once

#include <optional>

#include "domain/continuous_order.hpp"

class IOrderQueue {
 public:
  [[nodiscard]] virtual bool Empty() const = 0;

  virtual void AddOrder(const ContinuousOrder &order) = 0;

  virtual std::optional<ContinuousOrder> Pop() = 0;

  virtual ~IOrderQueue() = default;
};