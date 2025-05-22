#pragma once

#include "contracts/order_updates_sender.hpp"
#include "contracts/fill_details_repo.hpp"
#include "contracts/backend_client.hpp"
#include "pkg/repository/fill_details_repo/fill_details_repo.hpp"
#include <memory>

class OrderUpdatesSender : public IOrderUpdatesSender {
 public:
  OrderUpdatesSender(std::shared_ptr<IBackendClient> backend_client)
      : fill_details_repo_(std::make_shared<FillDetailsRepo>()),
        backend_client_(std::move(backend_client)) {}

  bool SendFillDetails(const ContinuousOrder& order, const FillDetails& fill_details) override;
  bool SendCancelled(const ContinuousOrder& order) override;

  ~OrderUpdatesSender() override = default;
 private:
  void FlushAccumulatedIfNeeded();

  static const size_t kAccumulatedOrderUpdatesCount = 10;
  size_t order_updates_count_ = 0;

  std::mutex fill_details_repo_mutex_;
  std::shared_ptr<IFillDetailsRepo> fill_details_repo_;

  std::mutex cancelled_orders_mutex_;
  std::vector<ContinuousOrder> cancelled_orders_;

  std::shared_ptr<IBackendClient> backend_client_;
};