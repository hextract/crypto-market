#pragma once

#include <string>

#include "asset.hpp"

// TODO: where to store tick size for every pair.
class TradingPair {
 public:
  TradingPair(const Asset& base_asset, const Asset& quote_asset)
     : base_asset_(base_asset), quote_asset_(quote_asset) {}

  [[nodiscard]] Asset GetBase() const { return base_asset_; }

  [[nodiscard]] Asset GetQuote() const { return quote_asset_; }

  [[nodiscard]] std::string GetTicker() const { return base_asset_.GetName() + quote_asset_.GetName(); }

 private:
  Asset base_asset_;
  Asset quote_asset_;
};

inline bool operator==(const TradingPair& lhs, const TradingPair& rhs) {
  return lhs.GetBase() == rhs.GetBase() && lhs.GetQuote() == rhs.GetQuote();
}

template<>
struct std::hash<TradingPair> {
  size_t operator()(const TradingPair& pair) const noexcept {
    return std::hash<std::string>{}(pair.GetBase().GetName() + pair.GetQuote().GetName());
  }
};