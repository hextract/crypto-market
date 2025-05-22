#pragma once

#include "domain/continuous_order.hpp"
#include "domain/fill_details.hpp"

class IOrderUpdatesSender {
 public:
  virtual bool SendFillDetails(const ContinuousOrder& order, const FillDetails& fill_details) = 0;
  virtual bool SendCancelled(const ContinuousOrder& order) = 0;
  virtual ~IOrderUpdatesSender() = default;
};