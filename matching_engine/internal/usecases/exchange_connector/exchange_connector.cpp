#include "exchange_connector.hpp"

void ExchangeConnector::PlaceOrder(const LimitExchangeOrder &order) {
  client_->PlaceOrder(order, GetOnSuccessCb(), GetOnErrorCb());
}

std::function<void(const Success &)> ExchangeConnector::GetOnSuccessCb() {
  return [](const Success& ) {};
}

std::function<void(const Error &)> ExchangeConnector::GetOnErrorCb() {
  return [](const Error& ) {};
}
