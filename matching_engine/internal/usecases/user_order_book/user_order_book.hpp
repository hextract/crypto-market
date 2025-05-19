#pragma once

#include "contracts/user_order_book_repo.hpp"
#include "contracts/user_order_book_usecase.hpp"
#include "contracts/backend_client.hpp"
#include <memory>
#include <utility>
#include "usecases/user_account_manager/user_account_manager.hpp"
#include "domain/continuous_order.hpp"
#include "domain/matched_details.hpp"
#include "domain/config.hpp"


class UserOrderBook : public IUserOrderBook {
 public:
  UserOrderBook(const TradingPair& pair, const std::shared_ptr<IUserOrderBookRepo>& storage)
      : pair_(pair), storage_(storage) {}

  void SetBackendClient(const std::shared_ptr<IBackendClient>& client) override;

  void AddOrder(const ContinuousOrder& order) override;

  void RemoveOrder(const ContinuousOrder& order) override;

  MatchedDetails MatchOrders(size_t clearing_price) override;

  ~UserOrderBook() override = default;

 private:
  size_t delta_ = Config::GetDeltaTime();
  TradingPair pair_;
  std::shared_ptr<IUserOrderBookRepo> storage_{nullptr};
  UserAccountManager account_manager_;
  std::shared_ptr<IBackendClient> backend_client_{nullptr};
};