#pragma once

class IGetClearingPrice {
 public:
  virtual size_t Execute() = 0;
  virtual ~IGetClearingPrice() = default;
};