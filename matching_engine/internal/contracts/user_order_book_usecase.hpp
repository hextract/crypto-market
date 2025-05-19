#pragma once

#include "domain/continuous_order.hpp"
#include "domain/matched_details.hpp"
#include "contracts/backend_client.hpp"
#include <memory>

class IUserOrderBook {
 public:
  virtual void AddOrder(const ContinuousOrder& order) = 0;

  virtual void SetBackendClient(const std::shared_ptr<IBackendClient>& client) = 0;

  virtual void RemoveOrder(const ContinuousOrder& order) = 0;

  virtual MatchedDetails MatchOrders(size_t clearing_price) = 0;

  virtual ~IUserOrderBook() = default;
};