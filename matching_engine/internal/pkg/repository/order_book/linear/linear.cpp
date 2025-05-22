#include "linear.hpp"
#include <stdexcept>

size_t LinearOrderBookRepo::GetQuantity(size_t ticket_number) {
  if (ticket_number >= curve_.size()) {
    throw std::runtime_error("Ticket number is more than exposed max ticket");
  }
  return curve_[ticket_number];
}

void LinearOrderBookRepo::AddOrder(const ContinuousOrder &order) {
  for (size_t price = 0; price < curve_.size(); ++price) {
    curve_[price] += order.GetRoundedSpeed(price);
  }
}

void LinearOrderBookRepo::DeleteOrder(const ContinuousOrder &order) {
  for (size_t i = order.GetLowPrice(); i <= order.GetHighPrice(); ++i) {
    curve_[i] -= order.GetSpeed();
  }
}

// TODO: write work variant
std::vector<CurveFracture> LinearOrderBookRepo::GetCurve(size_t left_bound_price, size_t right_bound_price) {
  return {};
}
