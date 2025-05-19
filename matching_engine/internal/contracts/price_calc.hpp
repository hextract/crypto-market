#pragma once

#include <optional>
#include <cstdint>

class IPriceCalculator {
public:
  virtual std::optional<size_t> CalculatePrice() = 0;

  virtual void ChangeImbalance(int64_t delta) = 0;

  virtual ~IPriceCalculator() = default;
};