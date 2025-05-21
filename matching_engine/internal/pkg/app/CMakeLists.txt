add_library(app
        matching_engine.cpp
        ../../usecases/price_calculator/segment_treap/segment_treap_price_calculator.cpp
        ../../usecases/user_order_book/user_order_book.cpp
        ../../usecases/order_book/order_book.cpp
        ../../usecases/engine_manager/engine_manager.cpp
)

target_link_libraries(app
        PUBLIC
        crow_support
)
target_include_directories(app PUBLIC
        ${PROJECT_SOURCE_DIR}/internal/
)
