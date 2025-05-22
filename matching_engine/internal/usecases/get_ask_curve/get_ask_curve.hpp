#pragma once

#include "contracts/get_ask_curve_use_case.hpp"
#include "contracts/order_book_repo.hpp"
#include <memory>

class GetAskCurve : IGetAskCurve {
 public:
  GetAskCurve(std::shared_ptr<IOrderBookRepo> ask_curve) : ask_curve_(ask_curve) {}
  std::vector<CurveFracture> Execute(size_t left_boundary_price, size_t right_boundary_price) override;
  ~GetAskCurve() override = default;
 private:
  std::shared_ptr<IOrderBookRepo> ask_curve_;
};