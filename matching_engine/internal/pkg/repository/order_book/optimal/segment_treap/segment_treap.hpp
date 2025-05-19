#include <random>
#include <utility>
#include <functional>
#include <optional>
#include <mutex>
#pragma once

template <typename T, typename F>
class SegmentTreap {
 public:
  // > comparator for direct ordering
  using Order = std::function<bool(const T&, const T&)>;
  using Operation = std::function<F(const F&, const F&)>;

  SegmentTreap(const F& neutral, Operation operation, Order order) : rnd_(time(nullptr)), neutral_(neutral), operation_(operation), order_(order) {};

  void Insert(const T& key);
  void Delete(const T& key);

  // add on (left, right]
  void AddOnHalfInterval(const T& left_bound, const T& right_bound, const F& addition);

  std::vector<F> GetHalfIntervalValues(const T& left_bound, const T& right_bound);
  std::vector<T> GetHalfIntervalKeys(const T& left_bound, const T& right_bound);

  F GetUpperValue(const T& key);
  std::optional<T> GetUpperKey(const T& key);

  ~SegmentTreap() { RecursiveDeleteNode(root_); }

 private:
  struct Node {
    Node(const T& key, unsigned int priority) : key(key), priority(priority) {}

    T key;
    F accumulator;
    F value;

    size_t size = 1;
    unsigned int priority = 0;
    Node* left = nullptr;
    Node* right = nullptr;
  };

  std::vector<F> GetTreeValues(Node* node);
  std::vector<T> GetTreeKeys(Node* node);

  Node* GetMin(Node* node);
  Node* GetMax(Node* node);

  void PushAccumulated(Node* node);

  std::pair<Node*, Node*> Split(Node* node, const T& pivot);

  Node* Merge(Node* node1, Node* node2);

  static size_t GetSize(Node* node);
  F GetValue(Node* node);

  static void Update(Node* node);

  void RecursiveDeleteNode(Node* node);

  std::mutex mutex_;
  std::mt19937 rnd_;
  Node* root_ = nullptr;

  F neutral_;
  Operation operation_;
  Order order_;
};

#include "segment_treap.ipp"