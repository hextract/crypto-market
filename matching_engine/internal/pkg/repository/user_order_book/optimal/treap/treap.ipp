#include "treap.hpp"

// public

template <typename T>
void Treap<T>::Insert(const T& key) {
  std::lock_guard lock_guard(mutex_);

  if (Find(root_, key) == nullptr) {
    Node* new_node = new Node(key, rnd_());
    std::pair<Node*, Node*> result = Split(root_, key);
    root_ = Merge(Merge(result.first, new_node), result.second);
  }
}

template <typename T>
void Treap<T>::Delete(const T& key) {
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


template <typename T>
std::optional<T> Treap<T>::GetUpper(const T& key) {
  std::lock_guard lock_guard(mutex_);

  std::pair<Node*, Node*> split = Split(root_, key);
  Node* result = GetMin(split.second);
  root_ = Merge(split.first, split.second);
  if (result == nullptr) {
    return std::nullopt;
  }
  return result->key;
}

template <typename T>
std::vector<T> Treap<T>::GetTail(const T& pivot) {
  std::lock_guard lock_guard(mutex_);

  auto [tail, head] = Split(root_, pivot);
  auto tail_elems = GetTreeElems(tail);
  root_ = Merge(tail, head);
  return tail_elems;
}

// private

template <typename T>
Treap<T>::Node* Treap<T>::GetMin(Node* node) {
  while (node != nullptr && node->left != nullptr) {
    node = node->left;
  }
  return node;
}

template <typename T>
Treap<T>::Node* Treap<T>::GetMax(Node* node) {
  while (node != nullptr && node->right != nullptr) {
    node = node->right;
  }
  return node;
}

template <typename T>
std::vector<T> Treap<T>::GetTreeElems(Node* node) {
  std::vector<T> elems;
  if (node != nullptr) {
    for (auto& elem: GetTreeElems(node->left)) {
      elems.emplace_back(std::move(elem));
    }
    elems.emplace_back(node->key);
    for (auto& elem: GetTreeElems(node->right)) {
      elems.emplace_back(std::move(elem));
    }
  }
  return elems;
}

template <typename T>
std::pair<typename Treap<T>::Node*, typename Treap<T>::Node*> Treap<T>::Split(Node* node, const T& pivot) {
  if (node == nullptr) {
    return {nullptr, nullptr};
  }
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

template <typename T>
Treap<T>::Node* Treap<T>::Merge(Node* node1, Node* node2) {
  if (node1 == nullptr) {
    return node2;
  }
  if (node2 == nullptr) {
    return node1;
  }
  if (node1->priority > node2->priority) {
    node1->right = Merge(node1->right, node2);
    Update(node1);
    return node1;
  }
  node2->left = Merge(node1, node2->left);
  Update(node2);
  return node2;
}

template <typename T>
int Treap<T>::GetSize(Node* node) {
  if (node == nullptr) {
    return 0;
  }
  return node->size;
}

template <typename T>
void Treap<T>::Update(Node* node) {
  if (node != nullptr) {
    node->size = 1 + GetSize(node->left) + GetSize(node->right);
  }
}

template <typename T>
Treap<T>::Node* Treap<T>::Find(Node* curr_node, const T& key) {
  while (true) {
    if (curr_node == nullptr || curr_node->key == key) {
      return curr_node;
    }
    if (order_(curr_node->key, key)) {
      curr_node = curr_node->left;
    } else {
      curr_node = curr_node->right;
    }
  }
}

template <typename T>
void Treap<T>::RecursiveDeleteNode(Node* node) {
  if (node != nullptr) {
    RecursiveDeleteNode(node->left);
    RecursiveDeleteNode(node->right);
    delete node;
  }
}
