#include "segment_treap.hpp"

// public

template <typename T, typename F>
void SegmentTreap<T, F>::Insert(const T& key) {
  std::lock_guard lock_guard(mutex_);

  Node* new_node = new Node(key, rnd_());
  new_node->accumulator = neutral_;
  new_node->value = neutral_;

  std::pair<Node*, Node*> result = Split(root_, key);
  root_ = Merge(Merge(result.first, new_node), result.second);
}

template <typename T, typename F>
void SegmentTreap<T, F>::Delete(const T& key) {
  std::lock_guard lock_guard(mutex_);

  if (root_ == nullptr) {
    return;
  }

  Node* parent = nullptr;
  Node* node = root_;
  while (node != nullptr && node->key != key) {
    parent = node;
    if (order_(key,node->key)) {
      node = node->left;
    } else {
      node = node->right;
    }
  }
  if (node == nullptr) {
    return;
  }

  Node* result = Merge(node->left, node->right);
  delete node;
  if (parent != nullptr) {
    if (order_(parent->key, key)) {
      parent->right = result;
    } else {
      parent->left = result;
    }
  } else {
    root_ = result;
  }
}

template <typename T, typename F>
std::optional<T> SegmentTreap<T, F>::GetUpperKey(const T& key) {
  std::lock_guard lock_guard(mutex_);

  std::pair<Node*, Node*> split = Split(root_, key);
  Node* result = GetMin(split.second);
  root_ = Merge(split.first, split.second);
  if (result == nullptr) {
    return std::nullopt;
  }
  return result->key;
}

template <typename T, typename F>
F SegmentTreap<T, F>::GetUpperValue(const T& key) {
  std::lock_guard lock_guard(mutex_);

  std::pair<Node*, Node*> split = Split(root_, key);
  Node* result = GetMin(split.second);
  root_ = Merge(split.first, split.second);
  if (result == nullptr) {
    return neutral_;
  }
  return GetValue(result);
}


template <typename T, typename F>
void SegmentTreap<T, F>::AddOnHalfInterval(const T& left_bound, const T& right_bound, const F& addition) {
  std::lock_guard lock_guard(mutex_);

  auto [left_beam, head] = Split(root_, right_bound);
  auto [tail, half_interval] = Split(left_beam, left_bound);

  if (half_interval != nullptr) {
    half_interval->accumulator = operation_(half_interval->accumulator, addition);
  }

  left_beam = Merge(tail, half_interval);
  root_ = Merge(left_beam, head);
}

template <typename T, typename F>
std::vector<T> SegmentTreap<T, F>::GetHalfIntervalKeys(const T& left_bound, const T& right_bound) {
  std::lock_guard lock_guard(mutex_);

  auto [left_beam, head] = Split(root_, right_bound);
  auto [tail, half_interval] = Split(left_beam, left_bound);

  auto keys = GetTreeKeys(half_interval);

  left_beam = Merge(tail, half_interval);
  root_ = Merge(left_beam, head);

  return keys;
}


template <typename T, typename F>
std::vector<F> SegmentTreap<T, F>::GetHalfIntervalValues(const T& left_bound, const T& right_bound) {
  std::lock_guard lock_guard(mutex_);

  auto [left_beam, head] = Split(root_, right_bound);
  auto [tail, half_interval] = Split(left_beam, left_bound);

  auto values = GetTreeValues(half_interval);

  left_beam = Merge(tail, half_interval);
  root_ = Merge(left_beam, head);

  return values;
}

// private

template <typename T, typename F>
std::vector<F> SegmentTreap<T, F>::GetTreeValues(Node* node) {
  std::vector<F> values;
  if (node != nullptr) {
    PushAccumulated(node);
    for (auto& value: GetTreeValues(node->left)) {
      values.emplace_back(value);
    }
    values.emplace_back(GetValue(node));
    for (auto& value: GetTreeValues(node->right)) {
      values.emplace_back(value);
    }
  }
  return values;
}

template <typename T, typename F>
std::vector<T> SegmentTreap<T, F>::GetTreeKeys(Node* node) {
  std::vector<T> keys;
  if (node != nullptr) {
    PushAccumulated(node);
    for (auto& key: GetTreeKeys(node->left)) {
      keys.emplace_back(key);
    }
    keys.emplace_back(node->key);
    for (auto& key: GetTreeKeys(node->right)) {
      keys.emplace_back(key);
    }
  }
  return keys;
}


template <typename T, typename F>
void SegmentTreap<T, F>::PushAccumulated(Node* node) {
  if (node == nullptr) {
    return;
  }
  node->value = operation_(node->value, node->accumulator);
  if (node->left != nullptr) {
    node->left->accumulator = operation_(node->left->accumulator, node->accumulator);
  }
  if (node->right != nullptr) {
    node->right->accumulator = operation_(node->right->accumulator, node->accumulator);
  }
  node->accumulator = neutral_;
}

template <typename T, typename F>
SegmentTreap<T, F>::Node* SegmentTreap<T, F>::GetMin(Node* node) {
  while (node != nullptr && node->left != nullptr) {
    PushAccumulated(node);
    node = node->left;
  }
  return node;
}

template <typename T, typename F>
SegmentTreap<T, F>::Node* SegmentTreap<T, F>::GetMax(Node* node) {
  while (node != nullptr && node->right != nullptr) {
    PushAccumulated(node);
    node = node->right;
  }
  return node;
}

template <typename T, typename F>
std::pair<typename SegmentTreap<T, F>::Node*, typename SegmentTreap<T, F>::Node*>
SegmentTreap<T, F>::Split(Node* node, const T& pivot) {
  if (node == nullptr) {
    return {nullptr, nullptr};
  }
  PushAccumulated(node);
  if (order_(node->key, pivot) || node->key == pivot) {
    std::pair<Node*, Node*> result = Split(node->right, pivot);
    node->right = result.first;
    Update(node);
    return {node, result.second};
  }
  std::pair<Node*, Node*> result = Split(node->left, pivot);
  node->left = result.second;
  Update(node);
  return {result.first, node};
}

template <typename T, typename F>
SegmentTreap<T, F>::Node* SegmentTreap<T, F>::Merge(Node* node1, Node* node2) {
  if (node1 == nullptr) {
    return node2;
  }
  if (node2 == nullptr) {
    return node1;
  }
  PushAccumulated(node1);
  PushAccumulated(node2);
  if (node1->priority > node2->priority) {
    node1->right = Merge(node1->right, node2);
    Update(node1);
    return node1;
  }
  node2->left = Merge(node1, node2->left);
  Update(node2);
  return node2;
}

template <typename T, typename F>
size_t SegmentTreap<T, F>::GetSize(Node* node) {
  if (node == nullptr) {
    return 0;
  }
  return node->size;
}

template <typename T, typename F>
F SegmentTreap<T, F>::GetValue(Node* node) {
  if (node == nullptr) {
    return neutral_;
  }
  return operation_(node->value, node->accumulator);
}

template <typename T, typename F>
void SegmentTreap<T, F>::Update(Node* node) {
  if (node != nullptr) {
    node->size = 1 + GetSize(node->left) + GetSize(node->right);
  }
}

template <typename T, typename F>
void SegmentTreap<T, F>::RecursiveDeleteNode(Node* node) {
  if (node != nullptr) {
    RecursiveDeleteNode(node->left);
    RecursiveDeleteNode(node->right);
    delete node;
  }
}