#include "backend_client.hpp"

#include <thread>

void BackendClient::SendUpdate(const ContinuousOrder& order, const FillDetails& fill_details) {
  while (!app_ready_->load()) {
    app_ready_->wait(true);
  }
  client_->getLoop()->queueInLoop([this, order, fill_details]() {
    double tickSize = static_cast<double>(cfg_->GetTickSize(order.GetPair()));
    double partCount = static_cast<double>(cfg_->GetMinPartCnt(order.GetPair()));

    double boughtAmount = static_cast<double>(fill_details.GetFilledBaseSize()) / partCount;
    double payedAmount = static_cast<double>(fill_details.GetFilledQuoteSize()) / tickSize / partCount;

    std::string status = (fill_details.GetFilledBaseSize() >= order.GetVolume())
                         ? "finished"
                         : "partial";
    if (fill_details.GetAveragePrice() > 0) {
      accumulated_order_updates_[order.GetOrderId()].volume = boughtAmount;
      accumulated_order_updates_[order.GetOrderId()].price = payedAmount;
      accumulated_order_updates_[order.GetOrderId()].status = status;
      ++accumulated_order_updates_count_;
    }

    if (accumulated_order_updates_count_ >= kAccumulatedOrderUpdatesReleaseCount) {
      Json::Value body;
      for (auto [order_id, accumulated_order_update] : accumulated_order_updates_) {
        Json::Value order_update;
        order_update["order_id"] = static_cast<int64_t>(order_id);
        order_update["payed_amount"] = accumulated_order_update.price;
        order_update["bought_amount"] = accumulated_order_update.volume;
        order_update["status"] = accumulated_order_update.status;
        body.append(order_update);
      }

      auto req = HttpRequest::newHttpJsonRequest(body);
      accumulated_order_updates_.clear();
      accumulated_order_updates_count_ = 0;

      req->setMethod(HttpMethod::Patch);
      req->setPath("/market-maker/statuses");
      client_->sendRequest(
          req,
          [](ReqResult result,
                                         const HttpResponsePtr& resp) {
            std::cout << "First team backend notified about last " << kAccumulatedOrderUpdatesReleaseCount << " updates" << std::endl;
            if (result != ReqResult::Ok || !resp || resp->getStatusCode() != k200OK) {
              std::cout << "Failed to update status for orders"
                        << "; result=" << to_string(result)
                        << ", status="
                        << (resp ? resp->getStatusCode() : 0);
            }
          });
    }
  });
}
