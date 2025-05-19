#pragma once

#include "treap/treap.hpp"
#include "domain/continuous_order.hpp"
#include "contracts/user_order_book_repo.hpp"
#include <unordered_set>


class OrdersRepository: public IUserOrderBookRepo {
 public:
  OrdersRepository();

  void AddOrder(const ContinuousOrder& order) override;
  void RemoveOrder(const ContinuousOrder& order) override;

  std::unordered_set<ContinuousOrder> GetOrders(size_t price) override;

  ~OrdersRepository() override = default;

 private:
  Treap<ContinuousOrder> bid_orders_;
  Treap<ContinuousOrder> ask_orders_;
};