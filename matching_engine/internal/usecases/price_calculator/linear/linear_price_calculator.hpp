#pragma once

#include "contracts/price_calc.hpp"
#include "contracts/order_book_repo.hpp"
#include <memory>

class LinearPriceCalculator : public IPriceCalculator {
public:
  LinearPriceCalculator(std::shared_ptr<IOrderBookRepo> buying_curve, std::shared_ptr<IOrderBookRepo> selling_curve,
                        size_t ticket_amount)
    : buying_curve_(std::move(buying_curve)),
      selling_curve_(std::move(selling_curve)),
      ticket_amount_(ticket_amount) {}

  std::optional<size_t> CalculatePrice() override;

  void ChangeImbalance(int64_t delta) override;

  ~LinearPriceCalculator() override = default;

private:
  std::shared_ptr<IOrderBookRepo> buying_curve_;
  std::shared_ptr<IOrderBookRepo> selling_curve_;
  int64_t imbalance_ = 0;
  size_t ticket_amount_;
};
