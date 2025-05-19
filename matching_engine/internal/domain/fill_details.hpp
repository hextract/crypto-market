#pragma once


class FillDetails {
 public:
  FillDetails() = default;

  FillDetails(size_t filled_base_size, size_t filled_quote_size, size_t average_price)
      : filled_base_size_(filled_base_size), filled_quote_size_(filled_quote_size),
        average_price_(average_price) {}

  [[nodiscard]] size_t GetFilledBaseSize() const { return filled_base_size_; }

  [[nodiscard]] size_t GetFilledQuoteSize() const { return filled_quote_size_; }

  [[nodiscard]] size_t GetAveragePrice() const { return average_price_; }

  void AddBaseCoinFill(size_t quantity) {
    filled_base_size_ += quantity;
  }

  void AddQuoteCoinFill(size_t quantity) {
    filled_quote_size_ += quantity;
  }

  void SetAveragePrice(size_t price) {
    average_price_ = price;
  }

 private:
  size_t filled_base_size_ = 0;
  size_t filled_quote_size_ = 0;
  size_t average_price_ = 0;
};