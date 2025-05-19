#pragma once

#include <contracts/order_book_repo.hpp>
#include "domain/continuous_order.hpp"

class IOrderBook {
 public:
  virtual void AddOrder(const ContinuousOrder& order) = 0;

  virtual void DeleteOrder(const ContinuousOrder& order) = 0;

  virtual size_t GetQuantity(size_t ticket_number) = 0;

  virtual std::vector<CurveFracture> GetCurve(size_t left_bound_price, size_t right_bound_price) = 0;

  virtual ~IOrderBook() = default;
};