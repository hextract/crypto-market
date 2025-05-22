#pragma once

#include "domain/limit_exchange_order.hpp"

class IExchangeConnectorClient {
 public:
  virtual void PlaceOrder(const LimitExchangeOrder& order, const std::function<void(const Success&)>& on_success,
       const std::function<void(const Error&)>& on_error) = 0;

  virtual ~IExchangeConnectorClient() = default;
};