#pragma once
#include <random>
#include <utility>
#include <mutex>
#include <functional>
#include <optional>

template <typename T>
class Treap {
 public:
  // > comparator for direct ordering
  using Order = std::function<bool(const T&, const T&)>;

  Treap(Order order) : rnd_(time(nullptr)), order_(order) {};

  void Insert(const T& key);

  void Delete(const T& key);

  std::optional<T> GetUpper(const T& key);

  std::vector<T> GetTail(const T& pivot);

  ~Treap() { RecursiveDeleteNode(root_); }

 private:
  struct Node {
    Node(const T& key, unsigned int priority) : key(key), priority(priority) {}

    T key;
    int size = 1;
    unsigned int priority = 0;
    Node* left = nullptr;
    Node* right = nullptr;
  };

  Node* GetMin(Node* node);
  Node* GetMax(Node* node);

  std::vector<T> GetTreeElems(Node* node);

  std::pair<Node*, Node*> Split(Node* node, const T& pivot);
  Node* Merge(Node* node1, Node* node2);

  static int GetSize(Node* node);

  static void Update(Node* node);

  Node* Find(Node* curr_node, const T& key);

  void RecursiveDeleteNode(Node* node);

  std::mutex mutex_;
  std::mt19937 rnd_;
  Node* root_ = nullptr;
  Order order_;
};

#include "treap.ipp"

