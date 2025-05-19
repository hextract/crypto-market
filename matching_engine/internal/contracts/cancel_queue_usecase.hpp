#pragma once

#include <optional>

#include "domain/continuous_order.hpp"

class ICancelQueue {
  public:
  [[nodiscard]] virtual bool Empty() const = 0;

  virtual void AddOrder(const CancelOrder &order) = 0;

  virtual std::optional<CancelOrder> Pop() = 0;

  virtual ~ICancelQueue() = default;
};