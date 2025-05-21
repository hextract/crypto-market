#pragma once

#include "domain/limit_exchange_order.hpp"
#include "domain/fill_details.hpp"
#include "domain/continuous_order.hpp"
#include <unordered_map>
#include <vector>

class IBackendClient {
  public:
  virtual bool SendFillDetails(const std::unordered_map<ContinuousOrder,FillDetails>& orders_fill_details) = 0;
  virtual bool SendCancelled(const std::vector<ContinuousOrder>& cancelled_orders) = 0;

  virtual ~IBackendClient() = default;
};