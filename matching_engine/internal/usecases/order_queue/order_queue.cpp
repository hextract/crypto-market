#include "order_queue.hpp"
#include <mutex>
#include <iostream>

void OrderQueue::AddOrder(const ContinuousOrder &order) {
  std::lock_guard l(m_);
  queue_.push(order);
}

std::optional<ContinuousOrder> OrderQueue::Pop() {
  std::lock_guard l(m_);
  if (queue_.empty()) {
    return std::nullopt;
  }
  ContinuousOrder top = std::move(queue_.front());
  queue_.pop();
  return std::make_optional(top);
}
