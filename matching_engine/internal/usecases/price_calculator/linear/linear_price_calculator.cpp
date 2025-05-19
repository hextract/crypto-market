#include "linear_price_calculator.hpp"

void LinearPriceCalculator::ChangeImbalance(int64_t delta) {
  imbalance_ += delta;
}

std::optional<size_t> LinearPriceCalculator::CalculatePrice() {
  for (size_t i = 0; i < ticket_amount_; ++i) {
    if (buying_curve_->GetQuantity(i) == 0) {
      return std::nullopt; // no buyers left;
    }

    if (buying_curve_->GetQuantity(i) == selling_curve_->GetQuantity(i)) {
      return std::make_optional(i); // don't need to change imbalance
    }
    if (buying_curve_->GetQuantity(i) < selling_curve_->GetQuantity(i)) {
      if (i == 0) {
        // must not happend in real life
        return std::make_optional(0);
      }
      int64_t left_point_delta = buying_curve_->GetQuantity(i - 1) - selling_curve_->GetQuantity(i - 1);
      int64_t right_point_delta = buying_curve_->GetQuantity(i) - selling_curve_->GetQuantity(i);
      if (abs(imbalance_ + left_point_delta) <= abs(imbalance_ + right_point_delta)) {
        imbalance_ += left_point_delta;
        return std::make_optional(i - 1);
      }
      imbalance_ += right_point_delta;
      return std::make_optional(i);
    }
  }
  return std::nullopt;
}
