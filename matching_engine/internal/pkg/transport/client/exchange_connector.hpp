#pragma once

#include "contracts/exchange_connector_client.hpp"
#include "domain/config.hpp"
#include "domain/limit_exchange_order.hpp"

#include <drogon/HttpClient.h>
#include <drogon/HttpAppFramework.h>

using namespace drogon;

class ExchangeConnectorClient : public IExchangeConnectorClient {
  public:
    ExchangeConnectorClient()
      : client_(HttpClient::newHttpClient(Config::GetExchangeConnectorUrl())) {
    }

    void PlaceOrder(const LimitExchangeOrder &order,
                   const std::function<void(const Success &)> &on_success,
                   const std::function<void(const Error &)> &on_error) override {
      Json::Value body;
      body["order_id"] = static_cast<Json::Int64>(order.GetOrderId());
      body["type"] = order.GetOrderType();
      body["indicator"] = order.GetIndicator();
      body["price"] = order.GetPrice();
      body["volume"] = order.GetVolume();
      body["ticker"] = order.GetTicker();

      auto req = HttpRequest::newHttpJsonRequest(body);
      req->setMethod(HttpMethod::Post);
      req->setPath("/place-order");

      client_->sendRequest(req,
                           [on_success, on_error](ReqResult result, const HttpResponsePtr &resp) {
                             if (result == ReqResult::Ok && resp) {
                               auto json = resp->getJsonObject();
                               if (resp->getStatusCode() == k200OK && json && json->isMember("message")) {
                                 Success s;
                                 s.message = (*json)["message"].asString();
                                 on_success(s);
                               } else {
                                 Error e;
                                 e.status_code = resp->getStatusCode();
                                 if (json && json->isMember("error_message"))
                                   e.message = (*json)["error_message"].asString();
                                 else
                                   e.message = resp->getBody();
                                 on_error(e);
                               }
                             } else {
                               Error e;
                               e.status_code = resp ? resp->getStatusCode() : 0;
                               e.message = result == ReqResult::Ok
                                             ? std::string("Empty response")
                                             : std::string("Request failed: ") + to_string(result);
                               on_error(e);
                             }
                           }
      );
    }

    ~ExchangeConnectorClient() override = default;

  private:
    std::shared_ptr<HttpClient> client_;
};
