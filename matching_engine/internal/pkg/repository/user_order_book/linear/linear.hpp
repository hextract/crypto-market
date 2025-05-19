#pragma once

#include <map>
#include <unordered_set>
#include <vector>

#include "contracts/user_order_book_repo.hpp"
#include "domain/trading_pair.hpp"
#include "domain/fill_details.hpp"
#include "domain/continuous_order.hpp"

class LinearUserOrderBookRepo : public IUserOrderBookRepo {
 public:
  explicit LinearUserOrderBookRepo(const TradingPair& pair);

  void AddOrder(const ContinuousOrder& order) override;

  std::unordered_set<ContinuousOrder> GetOrders(size_t price) override;

  void RemoveOrder(const ContinuousOrder &order) override;

  ~LinearUserOrderBookRepo() override = default;

 private:
  TradingPair pair_;
  std::vector<std::unordered_set<ContinuousOrder>> book_;
};