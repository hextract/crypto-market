#include "ask_curve_repository.hpp"
#include <numeric>
#include "constants.hpp"

AskCurveRepository::AskCurveRepository() : constants_(0, std::plus(), std::greater()),
                                           koeffs_(0, std::plus(), std::greater()) {}

void AskCurveRepository::AddOrder(const ContinuousOrder& order) {
  double low_price = order.GetLowPrice();
  double high_price = order.GetHighPrice();
  double speed = order.GetSpeed();

  double koeff = speed / (high_price - low_price);
  double constant = order.GetSpeed();
  double line_shift = -koeff * low_price;

  // updating price fractures
  auto constants_first_fracture_upper = constants_.GetUpperKey(order.GetLowPrice() + 1);
  if (!constants_first_fracture_upper.has_value() ||
      constants_first_fracture_upper.value() != order.GetLowPrice()) {
    double upper_constant = constants_.GetUpperValue(order.GetLowPrice());
    constants_.Insert(order.GetLowPrice());
    constants_.AddOnHalfInterval(order.GetLowPrice() + 1,
                                 order.GetLowPrice(),
                                 upper_constant);
  }
  auto constants_second_fracture_upper = constants_.GetUpperKey(order.GetHighPrice() + 1);
  if (!constants_second_fracture_upper.has_value() ||
      constants_second_fracture_upper.value() != order.GetHighPrice()) {
    double upper_constant = constants_.GetUpperValue(order.GetHighPrice());
    constants_.Insert(order.GetHighPrice());
    constants_.AddOnHalfInterval(order.GetHighPrice() + 1,
                                 order.GetHighPrice(),
                                 upper_constant);
  }
  auto koeffs_first_fracture_upper = koeffs_.GetUpperKey(order.GetLowPrice() + 1);
  if (!koeffs_first_fracture_upper.has_value() ||
      koeffs_first_fracture_upper.value() != order.GetLowPrice()) {
    double upper_koeff = koeffs_.GetUpperValue(order.GetLowPrice());
    koeffs_.Insert(order.GetLowPrice());
    koeffs_.AddOnHalfInterval(order.GetLowPrice() + 1,
                              order.GetLowPrice(),
                              upper_koeff);
  }
  auto koeffs_second_fracture_upper = koeffs_.GetUpperKey(order.GetHighPrice() + 1);
  if (!koeffs_second_fracture_upper.has_value() ||
      koeffs_second_fracture_upper.value() != order.GetHighPrice()) {
    double upper_koeff = koeffs_.GetUpperValue(order.GetHighPrice());
    koeffs_.Insert(order.GetHighPrice());
    koeffs_.AddOnHalfInterval(order.GetHighPrice() + 1,
                              order.GetHighPrice(),
                              upper_koeff);
  }

  // adding order ask curve
  koeffs_.AddOnHalfInterval(order.GetHighPrice(),
                            order.GetLowPrice(), koeff);
  constants_.AddOnHalfInterval(order.GetHighPrice(),
                               order.GetLowPrice(), line_shift);
  constants_.AddOnHalfInterval(Constants::kInf,
                               order.GetHighPrice(),
                               constant);
}

void AskCurveRepository::DeleteOrder(const ContinuousOrder& order) {
  double low_price = order.GetLowPrice();
  double high_price = order.GetHighPrice();
  double speed = order.GetSpeed();

  double koeff = speed / (high_price - low_price);
  double constant = order.GetSpeed();
  double line_shift = -koeff * low_price;

  // removing order bid curve
  koeffs_.AddOnHalfInterval(order.GetHighPrice(),
                            order.GetLowPrice(), -koeff);
  constants_.AddOnHalfInterval(order.GetHighPrice(),
                               order.GetLowPrice(), -line_shift);
  constants_.AddOnHalfInterval(Constants::kInf,
                               order.GetHighPrice(),
                               -constant);

  // updating price fractures

  double constants_first_fracture_upper =
      constants_.GetUpperValue(order.GetLowPrice() + 1);
  double constants_second_fracture_upper =
      constants_.GetUpperValue(order.GetHighPrice() + 1);
  double koeffs_first_fracture_upper = koeffs_.GetUpperValue(order.GetLowPrice() + 1);
  double koeffs_second_fracture_upper = koeffs_.GetUpperValue(order.GetHighPrice() + 1);

  double constants_first_fracture_upper_next =
      constants_.GetUpperValue(order.GetLowPrice());
  double constants_second_fracture_upper_next =
      constants_.GetUpperValue(order.GetHighPrice());
  double koeffs_first_fracture_upper_next = koeffs_.GetUpperValue(order.GetLowPrice());
  double koeffs_second_fracture_upper_next = koeffs_.GetUpperValue(order.GetHighPrice());

  if (abs(constants_first_fracture_upper - constants_first_fracture_upper_next) < Constants::kConstantEps
      && abs(koeffs_first_fracture_upper - koeffs_first_fracture_upper_next) < Constants::kKoeffEps) {
    constants_.Delete(order.GetLowPrice());
    koeffs_.Delete(order.GetLowPrice());
  }
  if (abs(constants_second_fracture_upper - constants_second_fracture_upper_next) < Constants::kConstantEps
      && abs(koeffs_second_fracture_upper - koeffs_second_fracture_upper_next) < Constants::kKoeffEps) {
    constants_.Delete(order.GetHighPrice());
    koeffs_.Delete(order.GetHighPrice());
  }
}

size_t AskCurveRepository::GetQuantity(size_t price) {
  double koeff = koeffs_.GetUpperValue(price);
  double constant = constants_.GetUpperValue(price);
  return koeff * static_cast<double>(price) + constant;
}

std::vector<CurveFracture> AskCurveRepository::GetCurve(size_t left_bound_price, size_t right_bound_price) {
  auto prices = koeffs_.GetHalfIntervalKeys(right_bound_price, left_bound_price);
  prices.insert(prices.begin(), right_bound_price);
  if (prices.back() > left_bound_price) {
    prices.insert(prices.end(), left_bound_price);
  }

  std::vector<CurveFracture> curve_fractures;
  for (size_t price : prices) {
    curve_fractures.emplace_back(price, GetQuantity(price));
  }
  std::reverse(curve_fractures.begin(), curve_fractures.end());
  return curve_fractures;
}