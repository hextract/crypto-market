#include "backend_client.hpp"

#include <thread>

bool BackendClient::SendFillDetails(const std::unordered_map<ContinuousOrder, FillDetails>& orders_fill_details) {
  Json::Value body;
  for (const auto& [order, fill_details] : orders_fill_details) {
    Json::Value order_update;
    double tickSize = static_cast<double>(cfg_->GetTickSize(order.GetPair()));
    double partCount = static_cast<double>(cfg_->GetMinPartCnt(order.GetPair()));

    double boughtAmount = static_cast<double>(fill_details.GetFilledBaseSize()) / partCount;
    double payedAmount = static_cast<double>(fill_details.GetFilledQuoteSize()) / tickSize / partCount;

    order_update["order_id"] = static_cast<int64_t>(order.GetOrderId());
    order_update["paid_price"] = payedAmount;
    order_update["bought_amount"] = boughtAmount;

    std::string status = (fill_details.GetFilledBaseSize() >= order.GetVolume())
                         ? "finished"
                         : "partial";
    order_update["status"] = status;
    body.append(order_update);
  }

  auto req = std::make_shared<HttpRequestPtr>(HttpRequest::newHttpJsonRequest(body));
  (*req)->setMethod(HttpMethod::Post);
  (*req)->setPath("/market-maker/statuses");

  while (!app_ready_->load()) {
    app_ready_->wait(true);
  }
  client_->getLoop()->queueInLoop([this, req, updated_orders_count = orders_fill_details.size()]() {
    client_->sendRequest(
        *req,
        [updated_orders_count](ReqResult result,
                               const HttpResponsePtr& resp) {
          if (result != ReqResult::Ok || !resp || resp->getStatusCode() != k200OK) {
            std::cout << "Failed to update status for orders"
                      << "; result=" << to_string(result)
                      << ", status="
                      << (resp ? resp->getStatusCode() : 0);
          } else {
//            std::cout << "First team backend notified about " << updated_orders_count << " orders"
//                      << std::endl;
          }
        });
  });
  return true;
}

bool BackendClient::SendCancelled(const std::vector<ContinuousOrder>& cancelled_orders) {
  if (cancelled_orders.size() == 0) {
    return false;
  }
  Json::Value body;
  for (const auto& order : cancelled_orders) {
    Json::Value order_update;
    order_update["order_id"] = static_cast<int64_t>(order.GetOrderId());
    order_update["status"] = "cancelled";
    body.append(order_update);
  }

  auto req = std::make_shared<HttpRequestPtr>(HttpRequest::newHttpJsonRequest(body));
  (*req)->setMethod(HttpMethod::Patch);
  (*req)->setPath("/market-maker/statuses");

  while (!app_ready_->load()) {
    app_ready_->wait(true);
  }
  client_->getLoop()->queueInLoop([this, req, cancelled_count = cancelled_orders.size()]() {
    client_->sendRequest(
        *req,
        [cancelled_count](ReqResult result,
                               const HttpResponsePtr& resp) {
          if (result != ReqResult::Ok || !resp || resp->getStatusCode() != k200OK) {
            std::cout << "Failed to update status for orders"
                      << "; result=" << to_string(result)
                      << ", status="
                      << (resp ? resp->getStatusCode() : 0);
          } else {
            std::cout << "First team backend notified about cancelling " << cancelled_count << " orders"
                      << std::endl;
          }
        });
  });
  return true;
}