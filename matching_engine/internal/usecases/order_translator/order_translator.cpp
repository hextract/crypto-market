#include "order_translator.hpp"

#include <domain/config.hpp>
#include <cmath>

OrderTranslator::OrderTranslator() {
}

ContinuousOrder OrderTranslator::ToContinuous(
  const std::vector<LimitExchangeOrder> &limit_orders) {
  if (limit_orders.empty()) {
    return ContinuousOrder{0, TradingPair(Asset(""), Asset("")), OrderSide::Buy, 0, 0, 0, 0};
  }

  const auto &first = limit_orders.front();
  size_t id = first.GetOrderId();

  std::string ticker = first.GetTicker();
  Asset base(ticker.substr(0, ticker.size() / 2));
  Asset quote(ticker.substr(ticker.size() / 2));
  TradingPair pair{base, quote};

  OrderSide side = (first.GetIndicator() == "BUY" ? OrderSide::Buy : OrderSide::Sell);

  size_t tick_size = cfg_->GetTickSize(pair);
  size_t min_parts = cfg_->GetMinPartCnt(pair);
  size_t dt = Config::GetDeltaTime();

  size_t min_tick = std::numeric_limits<size_t>::max();
  size_t max_tick = 0;
  size_t sum_parts = 0;

  for (const auto& order: limit_orders) {
    size_t p_ticks = static_cast<size_t>(std::floor(order.GetPrice() / tick_size));
    min_tick = std::min(min_tick, p_ticks);
    max_tick = std::max(max_tick, p_ticks);

    size_t v_parts = static_cast<size_t>(std::floor(order.GetVolume() * min_parts));
    sum_parts += v_parts;
  }

  size_t speed = (dt > 0 ? sum_parts / dt : sum_parts);
  return ContinuousOrder{id, pair, side, sum_parts, min_tick, max_tick, speed};
}

std::vector<LimitExchangeOrder> OrderTranslator::ToLimit(
  const ContinuousOrder &o) {
  std::vector<LimitExchangeOrder> res;

  size_t dt = Config::GetDeltaTime();
  size_t tickSize = cfg_->GetTickSize(o.GetPair());
  size_t minParts = cfg_->GetMinPartCnt(o.GetPair());

  double partsExecuted = static_cast<double>(o.GetSpeed()) * static_cast<double>(dt);
  double vol = partsExecuted / static_cast<double>(minParts);

  double price = static_cast<double>(o.GetLowPrice()) * static_cast<double>(tickSize);

  std::string type = "limit";
  std::string indicator = (o.GetSide() == OrderSide::Buy ? "Buy" : "Sell");
  std::string ticker = o.GetPair().GetTicker();

  res.emplace_back(o.GetOrderId(), type, indicator, price, vol, ticker);
  return res;
}
