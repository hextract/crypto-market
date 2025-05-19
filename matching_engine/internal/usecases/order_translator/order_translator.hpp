#pragma once

#include "domain/continuous_order.hpp"
#include "domain/limit_exchange_order.hpp"
#include "contracts/order_translator_usecase.hpp"
#include "contracts/order_book_configuration_usecase.hpp"
#include <vector>
#include <memory>

class OrderTranslator : public IOrderTranslator {
 public:
  OrderTranslator();

  ContinuousOrder ToContinuous(const std::vector<LimitExchangeOrder>& limit_orders);

  std::vector<LimitExchangeOrder> ToLimit(const ContinuousOrder& order);
 private:
  std::shared_ptr<IOrderBookConfiguration> cfg_{nullptr};
};