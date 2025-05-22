#pragma once

#include <unordered_map>
#include "domain/continuous_order.hpp"
#include "domain/fill_details.hpp"

class UserAccountManager {
 public:
  UserAccountManager() = default;

  [[nodiscard]] FillDetails GetFillDetails(const ContinuousOrder& order) const {
    if (!filled_details_.contains(order)) {
      return {};
    }
    return filled_details_.at(order);
  }

  [[nodiscard]] size_t GetBaseCoinFilled(const ContinuousOrder& order) const {
    if (!filled_details_.contains(order)) {
      return 0;
    }
    return filled_details_.at(order).GetFilledBaseSize();
  }

  [[nodiscard]] size_t GetQuoteCoinFilled(const ContinuousOrder& order) const {
    if (!filled_details_.contains(order)) {
      return 0;
    }
    return filled_details_.at(order).GetFilledQuoteSize();
  }

  [[nodiscard]] size_t GetAveragePrice(const ContinuousOrder& order) const {
    if (!filled_details_.contains(order)) {
      return 0;
    }
    return filled_details_.at(order).GetAveragePrice();
  }

  void AddFilledQuantity(const ContinuousOrder& order, size_t price, size_t quantity) {
    if (!filled_details_.contains(order)) {
      filled_details_[order] = FillDetails();
    }
    AddBaseCoinFill(order, quantity);
    AddQuoteCoinFill(order, price * quantity);
    UpdateAverage(order);
  }

  void AddBaseCoinFill(const ContinuousOrder& order, size_t quantity) {
    if (!filled_details_.contains(order)) {
      filled_details_[order] = FillDetails();
    }
    filled_details_[order].AddBaseCoinFill(quantity);
  }

  void AddQuoteCoinFill(const ContinuousOrder& order, size_t quantity) {
    if (!filled_details_.contains(order)) {
      filled_details_[order] = FillDetails();
    }
    filled_details_[order].AddQuoteCoinFill(quantity);
  }

  void UpdateAverage(const ContinuousOrder& order) {
    size_t base_size = filled_details_[order].GetFilledBaseSize();
    size_t quote_size = filled_details_[order].GetFilledQuoteSize();
    filled_details_[order].SetAveragePrice(quote_size / base_size);
  }

  void RemoveOrder(const ContinuousOrder& order) {
    if (!filled_details_.contains(order)) {
      return;
    }
    filled_details_.erase(order);
  }


 private:
  std::unordered_map<ContinuousOrder, FillDetails> filled_details_;
};
