#include "domain/trading_pair.hpp"
#include "domain/asset.hpp"
#include "order_book_configuration.hpp"

#include <vector>

OrderBookConfiguration::OrderBookConfiguration() {
  std::vector<std::pair<TradingPair, size_t>> trading_pairs = {
    {TradingPair(Asset("ETH"), Asset("USDT")), 2},
  };
  for (const auto& [pair, tick_size]: trading_pairs) {
    tick_sizes_[pair] = tick_size;
  }
}

size_t OrderBookConfiguration::GetTickSize(const TradingPair &pair) const {
  if (!tick_sizes_.contains(pair)) {
    throw std::runtime_error(pair.GetTicker() + " tick size not found!");
  }
  return tick_sizes_.at(pair);
}

size_t OrderBookConfiguration::GetMinPartCnt(const TradingPair &pair) const {
  if (!min_part_cnts_.contains(pair)) {
    throw std::runtime_error(pair.GetTicker() + " min part count not found!");
  }
  return min_part_cnts_.at(pair);
}

void OrderBookConfiguration::ChangeTickSize(const TradingPair &pair, size_t tick_size) {
  tick_sizes_[pair] = tick_size;
}

void OrderBookConfiguration::ChangeMinPartCnt(const TradingPair &pair, size_t min_part_cnt) {
  min_part_cnts_[pair] = min_part_cnt;
}
