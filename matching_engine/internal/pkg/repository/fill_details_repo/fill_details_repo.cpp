#include "fill_details_repo.hpp"

bool FillDetailsRepo::Set(const ContinuousOrder& order, const FillDetails& fill_details) {
  fill_details_[order] = fill_details;
  return true;
}
std::unordered_map<ContinuousOrder, FillDetails> FillDetailsRepo::GetAll() {
  return fill_details_;
}