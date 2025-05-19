#include "get_clearing_price.hpp"

size_t GetClearingPrice::Execute() {
  return price_calculator_->CalculatePrice();
}