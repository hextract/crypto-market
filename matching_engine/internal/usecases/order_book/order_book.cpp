#include "order_book.hpp"

void OrderBook::AddOrder(const ContinuousOrder &order) {
  storage_->AddOrder(order);
}

void OrderBook::DeleteOrder(const ContinuousOrder &order) {
  storage_->DeleteOrder(order);
}

size_t OrderBook::GetQuantity(size_t ticket_number) {
  return storage_->GetQuantity(ticket_number);
}

std::vector<CurveFracture> OrderBook::GetCurve(size_t left_bound_price, size_t right_bound_price) {
  return storage_->GetCurve(left_bound_price, right_bound_price);
}
