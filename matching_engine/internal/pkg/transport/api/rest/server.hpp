#pragma once

#include "crow.h"
#include <unordered_map>
#include <functional>
#include <string>
#include <algorithm>

#include "contracts/order_book_configuration_usecase.hpp"
#include "contracts/order_book_usecase.hpp"
#include "domain/config.hpp"
#include "domain/curve_fracture.hpp"

#include "mappers.hpp"

class HTTPServer {
 public:
  using GetCurveHandlerType = std::function<std::vector<CurveFracture>(size_t left_boundary_price,
                                                                       size_t right_boundary_price)>;
  explicit HTTPServer(const std::shared_ptr<IOrderBookConfiguration>& order_book_config)
      : order_book_config_(order_book_config), mapper_(order_book_config) {
    SetupHandlers();
  }

  void SetCreateOrderCallback(const std::function<void(const ContinuousOrder&)>& create_order_cb) {
    create_order_cb_ = create_order_cb;
  }

  void SetCancelOrderCallback(const std::function<void(const CancelOrder&)>& cancel_order_cb) {
    cancel_order_cb_ = cancel_order_cb;
  }

  void SetGetBidCurveHandler(const GetCurveHandlerType& get_bid_curve_handler) {
    get_bid_curve_handler_ = get_bid_curve_handler;
  }

  void SetGetAskCurveHandler(const GetCurveHandlerType& get_ask_curve_handler) {
    get_ask_curve_handler_ = get_ask_curve_handler;
  }

  void SetGetClearingPriceHandler(const std::function<std::optional<size_t>()>& get_clearing_price_handler) {
    get_clearing_price_handler_ = get_clearing_price_handler;
  }

  void SetPair(const TradingPair& pair) {
    mapper_.SetPair(pair);
  }

  void Run() {
    app_.port(Config::GetServerPort()).run();
  }

 private:
  void SetupHandlers() {
    CreateOrderHandler();
    CancelOrderHandler();
    GetAskCurveHandler();
    GetBidCurveHandler();
    GetClearingPriceHandler();
  }

  static crow::response GetCorsResponse() {
    crow::response response;
    response.add_header("Access-Control-Allow-Origin", "*");
    response.add_header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS");
    response.add_header("Access-Control-Allow-Headers", "*");
    return response;
  }

  void CreateOrderHandler() {
    CROW_ROUTE(app_, "/create-order").methods(crow::HTTPMethod::Post, crow::HTTPMethod::Options)(
        [this](const crow::request& req) {
          crow::response response = GetCorsResponse();
          if (req.method == crow::HTTPMethod::Options) {
            return response;
          }
          try {
            auto body_obj = crow::json::load(req.body);
            if (!body_obj) {
              return crow::response(400, "Invalid JSON!");
            }
            ContinuousOrder order = mapper_.ToDomainOrder(body_obj);

            create_order_cb_(order);

            crow::json::wvalue result;
            result["status"] = "ok";

            response.body = result.dump();
            return response;
          } catch (const std::exception& ex) {
            response.code = 400;
            response.body = std::string("Error parsing fields: ") + ex.what();
            return response;
          }
        }
    );
  }

  void GetClearingPriceHandler() {
    CROW_ROUTE(app_, "/clearing-price").methods(crow::HTTPMethod::GET, crow::HTTPMethod::Options)(
        [this](const crow::request& req) {
          crow::response response = GetCorsResponse();
          if (req.method == crow::HTTPMethod::Options) {
            return response;
          }
          try {
            std::optional<size_t> domain_price = get_clearing_price_handler_();
            if (!domain_price.has_value()) {
              throw std::runtime_error("There are no ask or bids, so the clearing price is undefined");
            }

            double price = mapper_.ToDtoPrice(domain_price.value());

            crow::json::wvalue result;
            result["price"] = price;

            response.body = result.dump();
            return response;
          } catch (const std::exception& ex) {
            response.code = 400;
            response.body = std::string("Error during getting clearing price: ") + ex.what();
            return response;
          }
        }
    );
  }

  void GetBidCurveHandler() {
    CROW_ROUTE(app_, "/bid-curve").methods(crow::HTTPMethod::Get, crow::HTTPMethod::Options)(
        [this](const crow::request& req) {
          crow::response response = GetCorsResponse();
          if (req.method == crow::HTTPMethod::Options) {
            return response;
          }
          try {
            auto body_obj = crow::json::load(req.body);
            if (!body_obj) {
              return crow::response(400, "Invalid JSON!");
            }

            std::pair<double, double> dto_boundaries = mapper_.ToDtoCurvePriceBoundaries(body_obj);
            if (dto_boundaries.first < 0 || dto_boundaries.second < 0) {
              throw std::runtime_error("non-positive price boundaries");
            }
            std::pair<double, double> boundaries = mapper_.ToDomainCurvePriceBoundaries(dto_boundaries);

            std::vector<CurveFracture> fractures = get_bid_curve_handler_(boundaries.first, boundaries.second);
            std::vector<crow::json::wvalue> json_fractures(fractures.size());
            std::transform(fractures.begin(),
                           fractures.end(),
                           json_fractures.begin(),
                           [this](const CurveFracture& fracture) {
                             return mapper_.ToJsonCurveFracture(fracture);
                           });
            std::transform(fractures.begin(),
                           fractures.end(),
                           json_fractures.begin(),
                           [this](const CurveFracture& fracture) {
                             return mapper_.ToJsonCurveFracture(fracture);
                           });

            crow::json::wvalue result = std::move(json_fractures);

            response.body = result.dump();
            return response;
          } catch (const std::exception& ex) {
            response.code = 400;
            response.body = std::string("Error parsing fields: ") + ex.what();
            return response;
          }
        }
    );
  }

  void GetAskCurveHandler() {
    CROW_ROUTE(app_, "/ask-curve").methods(crow::HTTPMethod::Get, crow::HTTPMethod::Options)(
        [this](const crow::request& req) {
          crow::response response = GetCorsResponse();
          if (req.method == crow::HTTPMethod::Options) {
            return response;
          }
          try {
            auto body_obj = crow::json::load(req.body);
            if (!body_obj) {
              return crow::response(400, "Invalid JSON!");
            }

            std::pair<double, double> dto_boundaries = mapper_.ToDtoCurvePriceBoundaries(body_obj);
            if (dto_boundaries.first < 0 || dto_boundaries.second < 0) {
              throw std::runtime_error("non-positive price boundaries");
            }
            std::pair<double, double> boundaries = mapper_.ToDomainCurvePriceBoundaries(dto_boundaries);

            std::vector<CurveFracture> fractures = get_ask_curve_handler_(boundaries.first, boundaries.second);
            crow::json::wvalue::list json_fractures(fractures.size());

            std::transform(fractures.begin(),
                           fractures.end(),
                           json_fractures.begin(),
                           [this](const CurveFracture& fracture) {
                             return mapper_.ToJsonCurveFracture(fracture);
                           });

            crow::json::wvalue result = std::move(json_fractures);

            response.body = result.dump();
            return response;
          } catch (const std::exception& ex) {
            response.code = 400;
            response.body = std::string("Error parsing fields: ") + ex.what();
            return response;
          }
        }
    );
  }

  void CancelOrderHandler() {
    CROW_ROUTE(app_, "/cancel-order").methods(crow::HTTPMethod::Post, crow::HTTPMethod::Options)(
        [this](const crow::request& req) {
          crow::response response = GetCorsResponse();
          if (req.method == crow::HTTPMethod::Options) {
            return response;
          }
          try {
            auto body_obj = crow::json::load(req.body);
            if (!body_obj) {
              return crow::response(400, "Invalid JSON!");
            }
            CancelOrder cancel_order = mapper_.ToDomainCancelOrder(body_obj);

            cancel_order_cb_(cancel_order);

            crow::json::wvalue result;
            result["status"] = "ok";

            response.body = result.dump();
            return response;
          } catch (const std::exception& ex) {
            response.code = 400;
            response.body = std::string("Error parsing fields: ") + ex.what();
            return response;
          }
        }
    );
  }

  std::shared_ptr<IOrderBookConfiguration> order_book_config_{nullptr};
  ServerMapper mapper_;
  crow::SimpleApp app_{};

  std::function<void(const ContinuousOrder&)> create_order_cb_;
  std::function<void(const CancelOrder&)> cancel_order_cb_;
  std::function<std::optional<size_t>()> get_clearing_price_handler_;

  GetCurveHandlerType get_ask_curve_handler_;
  GetCurveHandlerType get_bid_curve_handler_;
};
