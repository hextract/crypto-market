#pragma once

#include "contracts/order_book_configuration_usecase.hpp"
#include "contracts/backend_client.hpp"
#include "domain/config.hpp"
#include "domain/fill_details.hpp"
#include "domain/continuous_order.hpp"

#include <drogon/HttpClient.h>
#include <drogon/HttpAppFramework.h>
#include <json/json.h>
#include <atomic>

using namespace drogon;

class BackendClient : public IBackendClient {
 public:
  explicit BackendClient(const std::shared_ptr<IOrderBookConfiguration>& cfg)
      : cfg_(cfg), app_ready_(std::make_shared<std::atomic<bool>>(false)) {
    drogon::app().registerBeginningAdvice([this]() {
      client_ = HttpClient::newHttpClient(Config::GetBackendUrl(),
                                          drogon::app().getIOLoop(0),
                                          true,
                                          false);
      client_->setPipeliningDepth(10);
      app_ready_->store(true);
      app_ready_->notify_all();
    });
  }

  bool SendFillDetails(const std::unordered_map<ContinuousOrder,FillDetails>& orders_fill_details) override;
  bool SendCancelled(const std::vector<ContinuousOrder>& cancelled_orders) override;

  ~BackendClient() override = default;

 private:
  std::shared_ptr<std::atomic<bool>> app_ready_;
  std::shared_ptr<IOrderBookConfiguration> cfg_{nullptr};
  std::shared_ptr<HttpClient> client_{nullptr};
};
