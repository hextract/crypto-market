#pragma once

#include "crow.h"
#include <string>
#include <memory>
#include <cmath>

#include "domain/continuous_order.hpp"
#include "domain/curve_fracture.hpp"
#include "domain/trading_pair.hpp"
#include "domain/asset.hpp"
#include "contracts/order_book_configuration_usecase.hpp"

class ServerMapper {
 public:
  ServerMapper(const std::shared_ptr<IOrderBookConfiguration>& order_book_config)
      : order_book_config_(order_book_config) {}

  void SetPair(const TradingPair& pair) {
    pair_ = pair;
  }

  ContinuousOrder ToDomainOrder(const crow::json::rvalue& body) {
    auto& pair_json = body["pair"];
    if (pair_json.t() != crow::json::type::List || pair_json.size() != 2) {
      throw std::runtime_error("Field 'pair' must be an array of two strings");
    }
    size_t order_id = body["order_id"].i();
    TradingPair pair(Asset(pair_json[0].s()), Asset(pair_json[1].s()));
    OrderSide side = body["buy_sell_indicator"].b() ? OrderSide::Buy : OrderSide::Sell;
    size_t amount = std::floor(body["amount"].d() * static_cast<double>(order_book_config_->GetMinPartCnt(pair)));
    size_t price_low = std::floor(body["price_low"].d() * static_cast<double>(order_book_config_->GetTickSize(pair)));
    size_t price_high = std::floor(body["price_high"].d() * static_cast<double>(order_book_config_->GetTickSize(pair)));
    size_t
        speed = std::floor(body["speed"].d() * static_cast<double>(order_book_config_->GetMinPartCnt(pair)) / 3600 / 1000);

    if (price_high <= price_low) {
      throw std::runtime_error("price_high must be GREATER than price_low");
    }

    return ContinuousOrder{
        order_id, pair, side, amount, price_low, price_high, speed
    };
  }

  double ToDtoPrice(size_t price) {
    return static_cast<double>(price) / static_cast<double>(order_book_config_->GetTickSize(pair_));
  }

  crow::json::wvalue ToJsonCurveFracture(const CurveFracture& curve_fracture) {
    crow::json::wvalue json;
    json["price"] =
        static_cast<double>(curve_fracture.price) / static_cast<double>(order_book_config_->GetTickSize(pair_));
    json["volume"] =
        static_cast<double>(curve_fracture.volume) / static_cast<double>( order_book_config_->GetMinPartCnt(pair_));
    return json;
  }

  static std::pair<double, double> ToDtoCurvePriceBoundaries(const crow::json::rvalue& body) {
    std::pair<double, double> boundaries;
    boundaries.first = body["left_boundary_price"].d();
    boundaries.second = body["right_boundary_price"].d();
    return boundaries;
  }

  std::pair<size_t, size_t> ToDomainCurvePriceBoundaries(std::pair<double, double> dto_boundaries) {
    std::pair<size_t, size_t> domain_boundaries;
    domain_boundaries.first =
        llround(dto_boundaries.first * static_cast<double>(order_book_config_->GetTickSize(pair_)));
    domain_boundaries.second =
        llround(dto_boundaries.second * static_cast<double>(order_book_config_->GetTickSize(pair_)));
    return domain_boundaries;
  }

  static CancelOrder ToDomainCancelOrder(const crow::json::rvalue& body) {
    size_t order_id = body["order_id"].i();
    return CancelOrder{
        order_id
    };
  }

 private:
  TradingPair pair_ = TradingPair(Asset(""), Asset(""));
  std::shared_ptr<IOrderBookConfiguration> order_book_config_{nullptr};
};