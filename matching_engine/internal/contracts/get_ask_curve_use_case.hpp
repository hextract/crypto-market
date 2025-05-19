#pragma once

#include "domain/curve_fracture.hpp"
#include <vector>

class IGetAskCurve {
 public:
  virtual std::vector<CurveFracture> Execute(size_t left_boundary_price, size_t right_boundary_price) = 0;
  virtual ~IGetAskCurve() = default;
};