#pragma once

#include "contracts/order_book_repo.hpp"
#include "contracts/order_book_usecase.hpp"
#include "domain/curve_fracture.hpp"
#include <memory>

class OrderBook : public IOrderBook {
public:
  explicit OrderBook(std::shared_ptr<IOrderBookRepo> storage)
    : storage_(std::move(storage)) {}

  void AddOrder(const ContinuousOrder& order) override;

  void DeleteOrder(const ContinuousOrder& order) override;

  std::vector<CurveFracture> GetCurve(size_t left_bound_price, size_t right_bound_price) override;

  size_t GetQuantity(size_t ticket_number) override;

  ~OrderBook() override = default;

private:
  std::shared_ptr<IOrderBookRepo> storage_;
};