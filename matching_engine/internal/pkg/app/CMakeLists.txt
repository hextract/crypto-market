add_library(app
        matching_engine.cpp
        ../../usecases/price_calculator/segment_treap/segment_treap_price_calculator.cpp
)

target_link_libraries(app
        PUBLIC
        crow_support
)
target_include_directories(app PUBLIC
        ${PROJECT_SOURCE_DIR}/internal/
)
