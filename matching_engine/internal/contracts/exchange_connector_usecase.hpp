#pragma once

#include "domain/limit_exchange_order.hpp"

class IExchangeConnector {
 public:
  virtual void PlaceOrder(const LimitExchangeOrder& order) = 0;

  virtual ~IExchangeConnector() = default;
};
