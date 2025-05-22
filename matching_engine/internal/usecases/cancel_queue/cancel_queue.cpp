#include "cancel_queue.hpp"
#include "domain/continuous_order.hpp"
#include <mutex>


void CancelQueue::AddOrder(const CancelOrder &order) {
  std::lock_guard l(m_);
  queue_.push(order);
}

std::optional<CancelOrder> CancelQueue::Pop() {
  std::lock_guard l(m_);
  if (queue_.empty()) {
    return std::nullopt;
  }
  CancelOrder top = std::move(queue_.front());
  queue_.pop();
  return std::make_optional(top);
}
