#pragma once

#include <memory>
#include <functional>

#include <domain/limit_exchange_order.hpp>
#include "contracts/exchange_connector_usecase.hpp"
#include "contracts/exchange_connector_client.hpp"

class ExchangeConnector : public IExchangeConnector {
 public:
  explicit ExchangeConnector(const std::shared_ptr<IExchangeConnectorClient>& client)
     : client_(client) {}

  void PlaceOrder(const LimitExchangeOrder &order) override;

 private:
  std::function<void(const Success&)> GetOnSuccessCb();

  std::function<void(const Error&)> GetOnErrorCb();

  std::shared_ptr<IExchangeConnectorClient> client_{nullptr};
};