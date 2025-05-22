#pragma once

#include "domain/trading_pair.hpp"

class IOrderBookConfiguration {
 public:
  [[nodiscard]] virtual size_t GetTickSize(const TradingPair& pair) const = 0;

  [[nodiscard]] virtual size_t GetMinPartCnt(const TradingPair& pair) const = 0;

  virtual void ChangeTickSize(const TradingPair& pair, size_t tick_size) = 0;

  virtual void ChangeMinPartCnt(const TradingPair& pair, size_t tick_size) = 0;

  virtual ~IOrderBookConfiguration() = default;
};