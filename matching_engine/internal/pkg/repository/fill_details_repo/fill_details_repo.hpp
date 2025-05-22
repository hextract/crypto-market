#pragma once

#include "contracts/fill_details_repo.hpp"

class FillDetailsRepo : public IFillDetailsRepo {
 public:
  bool Set(const ContinuousOrder& order, const FillDetails& fill_details) override;
  std::unordered_map<ContinuousOrder, FillDetails> GetAll() override;
  ~FillDetailsRepo() override = default;
 private:
  std::unordered_map<ContinuousOrder, FillDetails> fill_details_;
};