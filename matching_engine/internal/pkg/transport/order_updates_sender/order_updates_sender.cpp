#include "order_updates_sender.hpp"
#include <thread>

// public

bool OrderUpdatesSender::SendFillDetails(const ContinuousOrder& order, const FillDetails& fill_details) {
  std::lock_guard lock_guard(fill_details_repo_mutex_);
  fill_details_repo_->Set(order, fill_details);
  ++order_updates_count_;

  FlushAccumulatedIfNeeded();
  return true;
}

bool OrderUpdatesSender::SendCancelled(const ContinuousOrder& order) {
  std::lock_guard lock_guard(cancelled_orders_mutex_);
  cancelled_orders_.push_back(order);
  ++order_updates_count_;
  return true;
}

// private

void OrderUpdatesSender::FlushAccumulatedIfNeeded() {
  // maybe it is bad to copy all updates below
  if (order_updates_count_ >= kAccumulatedOrderUpdatesCount) {
    // sending fill details
    std::thread fill_details_worker([this, fill_details_repo=fill_details_repo_](){
      std::lock_guard lock_guard(fill_details_repo_mutex_);
      backend_client_->SendFillDetails(fill_details_repo->GetAll());
    });
    fill_details_worker.detach();

    // sending cancelled orders
    std::thread cancelled_worker([this, cancelled_orders=cancelled_orders_](){
      std::lock_guard lock_guard(cancelled_orders_mutex_);
      backend_client_->SendCancelled(cancelled_orders);
    });
    cancelled_worker.detach();

    fill_details_repo_ = std::make_shared<FillDetailsRepo>();
    cancelled_orders_.clear();
    order_updates_count_ = 0;
  }
}
