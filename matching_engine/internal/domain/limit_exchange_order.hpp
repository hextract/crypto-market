#pragma once

class LimitExchangeOrder {
 public:
  LimitExchangeOrder(
    size_t order_id, std::string order_type, std::string indicator,
    double price, double volume, std::string ticker)
      : order_id_(order_id), order_type_(std::move(order_type)),
        indicator_(std::move(indicator)), price_(price), volume_(volume),
        ticker_(std::move(ticker)) {}

  [[nodiscard]] size_t GetOrderId() const { return order_id_; }

  [[nodiscard]] std::string GetOrderType() const { return order_type_; }

  [[nodiscard]] std::string GetIndicator() const { return indicator_; }

  [[nodiscard]] double GetPrice() const { return price_; }

  [[nodiscard]] double GetVolume() const { return volume_; }

  [[nodiscard]] std::string GetTicker() const { return ticker_; }

 private:
  size_t order_id_{0};
  std::string order_type_{};
  std::string indicator_{};
  double price_{0};
  double volume_{0};
  std::string ticker_{};
};

struct Success {
  std::string message;
};

struct Error {
  int         status_code;
  std::string message;
};