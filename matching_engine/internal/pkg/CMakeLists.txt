add_subdirectory(app)
add_subdirectory(repository)
add_subdirectory(transport)

add_library(pkg INTERFACE)

target_link_libraries(pkg INTERFACE
        app
        repository
        transport
)
