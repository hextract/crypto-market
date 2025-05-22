#pragma once

#include "contracts/user_order_book_repo.hpp"
#include "contracts/user_order_book_usecase.hpp"
#include "contracts/backend_client.hpp"
#include "contracts/order_updates_sender.hpp"

#include "pkg/repository/fill_details_repo/fill_details_repo.hpp"
#include "usecases/user_account_manager/user_account_manager.hpp"
#include "domain/continuous_order.hpp"
#include "domain/matched_details.hpp"
#include "domain/config.hpp"

#include <memory>
#include <utility>
#include <mutex>

class UserOrderBook : public IUserOrderBook {
 public:
  UserOrderBook(const TradingPair& pair, const std::shared_ptr<IUserOrderBookRepo>& storage,
                const std::shared_ptr<IOrderUpdatesSender>& order_updates_sender)
      : pair_(pair),
        storage_(storage),
        order_updates_sender_(order_updates_sender) {}

  void AddOrder(const ContinuousOrder& order) override;

  void RemoveOrder(const ContinuousOrder& order) override;

  void CancellOrder(const ContinuousOrder& order) override;

  MatchedDetails MatchOrders(size_t clearing_price) override;

  ~UserOrderBook() override = default;

 private:
  size_t delta_ = Config::GetDeltaTime();
  TradingPair pair_;
  UserAccountManager account_manager_;

  std::shared_ptr<IUserOrderBookRepo> storage_{nullptr};
  std::shared_ptr<IOrderUpdatesSender> order_updates_sender_;
};