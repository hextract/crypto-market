#pragma once

#include "contracts/order_book_configuration_usecase.hpp"
#include <unordered_map>

class OrderBookConfiguration : public IOrderBookConfiguration {
 public:
  OrderBookConfiguration();

  [[nodiscard]] size_t GetTickSize(const TradingPair &pair) const override;

  [[nodiscard]] size_t GetMinPartCnt(const TradingPair &pair) const override;

  void ChangeTickSize(const TradingPair &pair, size_t tick_size) override;

  void ChangeMinPartCnt(const TradingPair &pair, size_t tick_size) override;

  ~OrderBookConfiguration() override = default;

 private:
  std::unordered_map<TradingPair, size_t> tick_sizes_;
  std::unordered_map<TradingPair, size_t> min_part_cnts_;
};
