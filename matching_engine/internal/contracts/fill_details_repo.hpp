#pragma once

#include "domain/continuous_order.hpp"
#include "domain/fill_details.hpp"

#include <unordered_map>

class IFillDetailsRepo {
 public:
  virtual bool Set(const ContinuousOrder& order, const FillDetails& fill_details) = 0;
  virtual std::unordered_map<ContinuousOrder, FillDetails> GetAll() = 0;
  virtual ~IFillDetailsRepo() = default;
};