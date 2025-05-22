cmake_minimum_required(VERSION 3.22.1)
project(matching_engine LANGUAGES CXX)
set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++20")

find_package(Boost 1.66 REQUIRED COMPONENTS system thread)
find_package(Drogon REQUIRED)

add_library(crow_support INTERFACE)

target_include_directories(crow_support INTERFACE
        ${PROJECT_SOURCE_DIR}/include
)
target_compile_definitions(crow_support INTERFACE
        CROW_ENABLE_BOOST_ASIO
)
target_link_libraries(crow_support INTERFACE
        Boost::system
        Boost::thread
)

add_subdirectory(internal/usecases)
add_subdirectory(internal/pkg)
add_subdirectory(tests)

add_executable(matching_engine main.cpp)
target_link_libraries(matching_engine
        PRIVATE pkg usecases crow_support Drogon::Drogon
)
enable_testing()
