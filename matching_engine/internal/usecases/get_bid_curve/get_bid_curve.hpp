#pragma once

#include "contracts/get_bid_curve_use_case.hpp"
#include "contracts/order_book_repo.hpp"
#include <memory>

class GetBidCurve : IGetBidCurve {
 public:
  GetBidCurve(std::shared_ptr<IOrderBookRepo> bid_curve) : bid_curve_(bid_curve) {}
  std::vector<CurveFracture> Execute(size_t left_boundary_price, size_t right_boundary_price) override;
  ~GetBidCurve() override = default;
 private:
  std::shared_ptr<IOrderBookRepo> bid_curve_;
};