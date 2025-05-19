#pragma once

#include "domain/continuous_order.hpp"
#include "domain/curve_fracture.hpp"
#include <vector>

class IOrderBookRepo {
public:
  virtual void AddOrder(const ContinuousOrder& order) = 0;

  virtual void DeleteOrder(const ContinuousOrder& order) = 0;

  virtual size_t GetQuantity(size_t ticket_number) = 0;

  virtual std::vector<CurveFracture> GetCurve(size_t left_bound_price, size_t right_bound_price) = 0;

  virtual ~IOrderBookRepo() = default;
};