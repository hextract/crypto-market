add_library(usecases
        order_book/order_book.cpp
        user_order_book/user_order_book.cpp
        user_account_manager/user_account_manager.cpp
        order_queue/order_queue.cpp
        engine_manager/engine_manager.cpp
        order_translator/order_translator.cpp
        exchange_connector/exchange_connector.cpp
        price_calculator/linear/linear_price_calculator.cpp
        price_calculator/segment_treap/segment_treap_price_calculator.cpp
        order_book_configuration/order_book_configuration.cpp
        cancel_queue/cancel_queue.cpp
)

target_include_directories(usecases PUBLIC
        ${PROJECT_SOURCE_DIR}/internal/usecases/order_book
        ${PROJECT_SOURCE_DIR}/internal/usecases/user_order_book
        ${PROJECT_SOURCE_DIR}/internal/usecases/price_calculator
        ${PROJECT_SOURCE_DIR}/internal/usecases/user_account_manager
        ${PROJECT_SOURCE_DIR}/internal/usecases/order_queue
        ${PROJECT_SOURCE_DIR}/internal/usecases/engine_manager
        ${PROJECT_SOURCE_DIR}/internal/usecases/order_translator
        ${PROJECT_SOURCE_DIR}/internal/usecases/exchange_connector
        ${PROJECT_SOURCE_DIR}/internal/usecases/order_book_configuration
        ${PROJECT_SOURCE_DIR}/internal/usecases/cancel_queue
        ${PROJECT_SOURCE_DIR}/internal/
)