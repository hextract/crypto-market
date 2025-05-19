#pragma once

#include "domain/curve_fracture.hpp"
#include <vector>

class IGetBidCurve {
 public:
  virtual std::vector<CurveFracture> Execute(size_t left_boundary_price, size_t right_boundary_price) = 0;
  virtual ~IGetBidCurve() = default;
};