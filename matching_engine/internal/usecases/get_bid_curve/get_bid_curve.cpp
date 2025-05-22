#include "get_bid_curve.hpp"

std::vector<CurveFracture> GetBidCurve::Execute(size_t left_boundary_price, size_t right_boundary_price) {
  return bid_curve_->GetCurve(left_boundary_price, right_boundary_price);
}