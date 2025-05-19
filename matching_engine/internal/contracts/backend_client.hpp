#pragma once

#include "domain/limit_exchange_order.hpp"
#include "domain/fill_details.hpp"
#include "domain/continuous_order.hpp"

class IBackendClient {
  public:
  virtual void SendUpdate(const ContinuousOrder& order, const FillDetails& fill_details) = 0;

  virtual ~IBackendClient() = default;
};