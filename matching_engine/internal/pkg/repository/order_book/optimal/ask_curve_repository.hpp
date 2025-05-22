#pragma once

#include "pkg/repository/order_book/optimal/segment_treap/segment_treap.hpp"
#include "domain/continuous_order.hpp"
#include "domain/curve_fracture.hpp"
#include "contracts/order_book_repo.hpp"
#include <functional>
#include <mutex>

class AskCurveRepository : public IOrderBookRepo {
 public:
  AskCurveRepository();

  void AddOrder(const ContinuousOrder& order) override;
  void DeleteOrder(const ContinuousOrder& order) override;

  size_t GetQuantity(size_t price) override;
  std::vector<CurveFracture> GetCurve(size_t left_bound_price, size_t right_bound_price) override;

  ~AskCurveRepository() override = default;

 private:
  SegmentTreap<size_t, double> constants_;
  SegmentTreap<size_t, double> koeffs_;
};