add_library(transport
        notification/kafka/kafka.cpp
)

target_link_libraries(transport
        PUBLIC
        crow_support
)
