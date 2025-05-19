#pragma once

#include <queue>
#include <optional>
#include <mutex>

#include "domain/continuous_order.hpp"
#include "contracts/order_queue_usecase.hpp"

class OrderQueue : public IOrderQueue {
 public:
  OrderQueue() = default;

  [[nodiscard]] bool Empty() const override { return queue_.empty(); }

  void AddOrder(const ContinuousOrder &order) override;

  std::optional<ContinuousOrder> Pop() override;

  ~OrderQueue() override = default;

 private:
  std::mutex m_;
  std::queue<ContinuousOrder> queue_;
};
