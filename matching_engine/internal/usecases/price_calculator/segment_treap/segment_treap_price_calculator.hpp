#pragma once

#include "contracts/price_calc.hpp"
#include "contracts/order_book_repo.hpp"
#include <memory>
#include <cstdint>

class SegmentTreapPriceCalculator : public IPriceCalculator {
 public:
  SegmentTreapPriceCalculator(std::shared_ptr<IOrderBookRepo> bid_curve, std::shared_ptr<IOrderBookRepo> ask_curve)
      : bid_curve_(std::move(bid_curve)),
        ask_curve_(std::move(ask_curve)) {}

  std::optional<size_t> CalculatePrice() override;

  void ChangeImbalance(int64_t delta) override;

  ~SegmentTreapPriceCalculator() override = default;

 private:
  std::shared_ptr<IOrderBookRepo> bid_curve_;
  std::shared_ptr<IOrderBookRepo> ask_curve_;
  int64_t imbalance_ = 0;
};
