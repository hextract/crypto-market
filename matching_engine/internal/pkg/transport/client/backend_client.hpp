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

  void SendUpdate(const ContinuousOrder& order, const FillDetails& fill_details) override;

  ~BackendClient() override = default;

 private:
  struct AccumulatedOrderUpdate {
    double price = 0;
    double volume = 0;
    std::string status;
  };

  static const size_t kAccumulatedOrderUpdatesReleaseCount = 10;
  size_t accumulated_order_updates_count_ = 0;
  std::unordered_map<size_t, AccumulatedOrderUpdate> accumulated_order_updates_;
  std::shared_ptr<std::atomic<bool>> app_ready_;
  std::shared_ptr<IOrderBookConfiguration> cfg_{nullptr};
  std::shared_ptr<HttpClient> client_{nullptr};
};
