#pragma once

#include <queue>
#include <optional>
#include <mutex>

#include "domain/continuous_order.hpp"
#include "contracts/cancel_queue_usecase.hpp"

class CancelQueue : public ICancelQueue {
  public:
  CancelQueue() = default;

  [[nodiscard]] bool Empty() const override { return queue_.empty(); }

  void AddOrder(const CancelOrder &order) override;

  std::optional<CancelOrder> Pop() override;

  ~CancelQueue() override = default;

  private:
  std::mutex m_;
  std::queue<CancelOrder> queue_;
};
