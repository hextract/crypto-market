add_library(repository
        user_order_book/linear/linear.cpp
        user_order_book/segment_tree/segment_tree.cpp
        user_order_book/optimal/orders_repository.cpp
        user_order_book/optimal/treap/treap.ipp
        order_book/linear/linear.cpp
        order_book/optimal/ask_curve_repository.cpp
        order_book/optimal/bid_curve_repository.cpp
        order_book/optimal/segment_treap/segment_treap.ipp
        fill_details_repo/fill_details_repo.cpp
)

target_include_directories(repository PUBLIC
        ${PROJECT_SOURCE_DIR}/internal/
        ${PROJECT_SOURCE_DIR}/internal/domain
        ${PROJECT_SOURCE_DIR}/internal/contracts
        ${PROJECT_SOURCE_DIR}/internal/pkg/repository/order_book/optimal
        ${PROJECT_SOURCE_DIR}/internal/pkg/repository/order_book/optimal/segment_treap
        ${PROJECT_SOURCE_DIR}/internal/pkg/repository/user_order_book/optimal
        ${PROJECT_SOURCE_DIR}/internal/pkg/repository/user_order_book/optimal/treap
)