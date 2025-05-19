#include "segment_treap_price_calculator.hpp"
#include <limits>
#include <functional>
#include <algorithm>

namespace Constants {
const size_t kInfPrice = std::numeric_limits<size_t>::max();
} // namespace Constants

// finds the first value that satisfies predicate
template <typename Type>
Type BinarySearch(Type left, Type right, std::function<bool(const Type&)> predicate) {
  while (left < right) {
    Type mid = (left + right) / 2;
    if (predicate(mid)) {
      right = mid;
    } else {
      left = mid + 1;
    }
  }
  return left;
}

std::optional<size_t> SegmentTreapPriceCalculator::CalculatePrice() {
  auto ask_greater_bid_predicate = [this](size_t price){
    return ask_curve_->GetQuantity(price) > bid_curve_->GetQuantity(price);
  };
  auto ask_less_bid_predicate = [this](size_t price){
    return ask_curve_->GetQuantity(price) < bid_curve_->GetQuantity(price);
  };

  if (!ask_less_bid_predicate(0) || !ask_greater_bid_predicate(Constants::kInfPrice)) {
    return std::nullopt;
  }

  std::pair<size_t, size_t> intersection_interval;
  intersection_interval.second = BinarySearch<size_t>(0, Constants::kInfPrice, ask_greater_bid_predicate);
  intersection_interval.first = BinarySearch<size_t>(0, Constants::kInfPrice, std::not_fn(ask_less_bid_predicate)) - 1;

  if (intersection_interval.second - intersection_interval.first % 2 == 1) {
    if (imbalance_ > 0) {
      --imbalance_;
    } else {
      ++imbalance_;
      return (intersection_interval.second + intersection_interval.first + 1) / 2;
    }
  }
  return (intersection_interval.second + intersection_interval.first) / 2;
}

void SegmentTreapPriceCalculator::ChangeImbalance(int64_t delta) {
  imbalance_ += delta;
}