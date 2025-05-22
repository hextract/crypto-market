#include "get_ask_curve.hpp"

std::vector<CurveFracture> GetAskCurve::Execute(size_t left_boundary_price, size_t right_boundary_price) {
  return ask_curve_->GetCurve(left_boundary_price, right_boundary_price);
}