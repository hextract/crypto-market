#pragma once

#include <iostream>
#include <string>


class Asset {
 public:
  explicit Asset(std::string name) : name_(std::move(name)) {}

  [[nodiscard]] std::string GetName() const {return name_; }

 private:
  std::string name_;
};

inline bool operator==(const Asset& lhs, const Asset& rhs) {
  return lhs.GetName() == rhs.GetName();
}

template<>
struct std::hash<Asset> {
  size_t operator()(const Asset& asset) const noexcept {
    return std::hash<std::string>{}(asset.GetName());
  }
};
