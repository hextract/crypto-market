#include "bid_curve_repository.hpp"
#include "constants.hpp"
#include <functional>

BidCurveRepository::BidCurveRepository() : constants_(0, std::plus(), std::less()),
                                           koeffs_(0, std::plus(), std::less()) {}

void BidCurveRepository::AddOrder(const ContinuousOrder& order) {
  double koeff = static_cast<double>(order.GetSpeed())
      / (static_cast<double>(order.GetLowPrice()) - static_cast<double>(order.GetHighPrice()));
  double constant = order.GetSpeed();
  double line_shift = -koeff * order.GetLowPrice();

  // updating price fractures
  if (order.GetLowPrice() > 0) {
    auto constants_first_fracture_upper = constants_.GetUpperKey(order.GetLowPrice() - 1);
    if (!constants_first_fracture_upper.has_value() ||
        constants_first_fracture_upper.value() != order.GetLowPrice()) {
      double upper_constant = constants_.GetUpperValue(order.GetLowPrice());
      constants_.Insert(order.GetLowPrice());
      constants_.AddOnHalfInterval(order.GetLowPrice() - 1,
                                   order.GetLowPrice(),
                                   upper_constant);
    }
  }
  auto constants_second_fracture_upper = constants_.GetUpperKey(order.GetHighPrice() - 1);
  if (!constants_second_fracture_upper.has_value() ||
      constants_second_fracture_upper.value() != order.GetHighPrice()) {
    double upper_constant = constants_.GetUpperValue(order.GetHighPrice());
    constants_.Insert(order.GetHighPrice());
    constants_.AddOnHalfInterval(order.GetHighPrice() - 1,
                                 order.GetHighPrice(),
                                 upper_constant);
  }
  if (order.GetLowPrice() > 0) {
    auto koeffs_first_fracture_upper = koeffs_.GetUpperKey(order.GetLowPrice() - 1);
    if (!koeffs_first_fracture_upper.has_value() ||
        koeffs_first_fracture_upper.value() != order.GetLowPrice()) {
      double upper_koeff = koeffs_.GetUpperValue(order.GetLowPrice());
      koeffs_.Insert(order.GetLowPrice());
      koeffs_.AddOnHalfInterval(order.GetLowPrice() - 1,
                                order.GetLowPrice(),
                                upper_koeff);
    }
  }
  auto koeffs_second_fracture_upper = koeffs_.GetUpperKey(order.GetHighPrice() - 1);
  if (!koeffs_second_fracture_upper.has_value() ||
      koeffs_second_fracture_upper.value() != order.GetHighPrice()) {
    double upper_koeff = koeffs_.GetUpperValue(order.GetHighPrice());
    koeffs_.Insert(order.GetHighPrice());
    koeffs_.AddOnHalfInterval(order.GetHighPrice() - 1,
                              order.GetHighPrice(),
                              upper_koeff);
  }

  // adding order bid curve
  koeffs_.AddOnHalfInterval(order.GetLowPrice(),
                            order.GetHighPrice(), koeff);
  constants_.AddOnHalfInterval(order.GetLowPrice(),
                               order.GetHighPrice(), line_shift);
  constants_.AddOnHalfInterval(0, order.GetHighPrice(), constant);
}

void BidCurveRepository::DeleteOrder(const ContinuousOrder& order) {
  double koeff = static_cast<double>(order.GetSpeed())
      / (static_cast<double>(order.GetLowPrice()) - static_cast<double>(order.GetHighPrice()));
  double constant = order.GetSpeed();
  double line_shift = -koeff * order.GetLowPrice();

  // removing order bid curve
  koeffs_.AddOnHalfInterval(order.GetLowPrice(),
                            order.GetHighPrice(), -koeff);
  constants_.AddOnHalfInterval(order.GetLowPrice(),
                               order.GetHighPrice(), -line_shift);
  constants_.AddOnHalfInterval(0,
                               order.GetHighPrice(), -constant);

  // updating price fractures

  if (order.GetLowPrice() > 0) {
    double constants_first_fracture_upper =
        constants_.GetUpperValue(order.GetLowPrice() - 1);
    double koeffs_first_fracture_upper = koeffs_.GetUpperValue(order.GetLowPrice() - 1);
    double constants_first_fracture_upper_next =
        constants_.GetUpperValue(order.GetLowPrice());
    double koeffs_first_fracture_upper_next = koeffs_.GetUpperValue(order.GetLowPrice());

    if (abs(constants_first_fracture_upper - constants_first_fracture_upper_next) < Constants::kConstantEps &&
        abs(koeffs_first_fracture_upper - koeffs_first_fracture_upper_next) < Constants::kKoeffEps) {
      koeffs_.Delete(order.GetLowPrice());
      constants_.Delete(order.GetLowPrice());
    }
  }

  double constants_second_fracture_upper =
      constants_.GetUpperValue(order.GetHighPrice() - 1);
  double koeffs_second_fracture_upper = koeffs_.GetUpperValue(order.GetHighPrice() - 1);
  double constants_second_fracture_upper_next =
      constants_.GetUpperValue(order.GetHighPrice());
  double koeffs_second_fracture_upper_next = koeffs_.GetUpperValue(order.GetHighPrice());

  if (abs(constants_second_fracture_upper - constants_second_fracture_upper_next) < Constants::kConstantEps
      && abs(koeffs_second_fracture_upper - koeffs_second_fracture_upper_next) < Constants::kKoeffEps) {
    constants_.Delete(order.GetHighPrice());
    koeffs_.Delete(order.GetHighPrice());
  }
}

size_t BidCurveRepository::GetQuantity(size_t price) {
  double koeff = koeffs_.GetUpperValue(price);
  double constant = constants_.GetUpperValue(price);
  return koeff * price + constant;
}

std::vector<CurveFracture> BidCurveRepository::GetCurve(size_t left_bound_price, size_t right_bound_price) {
  auto prices = koeffs_.GetHalfIntervalKeys(left_bound_price,
                                            right_bound_price);
  prices.insert(prices.begin(), left_bound_price);
  if (prices.back() < right_bound_price) {
    prices.insert(prices.end(), right_bound_price);
  }

  std::vector<CurveFracture> curve_fractures;
  for (size_t price : prices) {
    curve_fractures.emplace_back(price, GetQuantity(price));
  }
  return curve_fractures;
}