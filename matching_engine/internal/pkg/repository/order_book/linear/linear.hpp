#pragma once

#include <vector>

#include "contracts/order_book_repo.hpp"

class LinearOrderBookRepo : public IOrderBookRepo {
public:
  explicit LinearOrderBookRepo(size_t tickets_amount) : curve_(tickets_amount, 0) {}

  void AddOrder(const ContinuousOrder& order) override;

  void DeleteOrder(const ContinuousOrder& order) override;

  size_t GetQuantity(size_t ticket_number) override;

  std::vector<CurveFracture> GetCurve(size_t left_bound_price, size_t right_bound_price) override;

  ~LinearOrderBookRepo() override = default;

private:
  std::vector<size_t> curve_;
};