#pragma once

#include "contracts/get_clearing_price_use_case.hpp"
#include "contracts/price_calc.hpp"
#include <memory>

class GetClearingPrice : IGetClearingPrice {
 public:
  GetClearingPrice(std::shared_ptr<IPriceCalculator> price_calculator) : price_calculator_(price_calculator) {}
  size_t Execute() override;
  ~IGetClearingPrice() override = default;
 private:
  std::shared_ptr<IPriceCalculator> price_calculator_;
};