#pragma once

#include <vector>
#include "domain/continuous_order.hpp"

class MatchedDetails {
 public:
  MatchedDetails() = default;

  MatchedDetails(int64_t imbalance, const std::vector<ContinuousOrder>& buy_filled,
      const std::vector<ContinuousOrder>& sell_filled)
    : imbalance_(imbalance), buy_filled_(buy_filled), sell_filled_(sell_filled) {}

  [[nodiscard]] int64_t GetImbalance() const {
    return imbalance_;
  }

  [[nodiscard]] std::vector<ContinuousOrder> GetBuyFilled() const {
    return buy_filled_;
  }

  [[nodiscard]] std::vector<ContinuousOrder> GetSellFilled() const {
    return sell_filled_;
  }

  void AddImbalance(int64_t add) {
    imbalance_ += add;
  }

  void AddBuyFilled(const ContinuousOrder& order) {
    buy_filled_.push_back(order);
  }

  void AddSellFilled(const ContinuousOrder& order) {
    sell_filled_.push_back(order);
  }

 private:
  int64_t imbalance_ = 0;
  std::vector<ContinuousOrder> buy_filled_;
  std::vector<ContinuousOrder> sell_filled_;
  std::vector<ContinuousOrder> buy_finished_;
  std::vector<ContinuousOrder> sell_finished_;
};
