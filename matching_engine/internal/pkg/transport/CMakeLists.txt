add_library(transport
        notification/kafka/kafka.cpp
        client/backend_client.cpp
        order_updates_sender/order_updates_sender.cpp
)

target_link_libraries(transport
        PUBLIC
        crow_support
        Drogon::Drogon
)

target_include_directories(transport PUBLIC
        ${PROJECT_SOURCE_DIR}/internal
        ${PROJECT_SOURCE_DIR}/internal/pkg/transport/client
)