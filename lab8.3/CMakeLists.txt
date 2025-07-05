cmake_minimum_required(VERSION 3.10)
project(lab8.3)

set(CMAKE_CXX_STANDARD 17)

find_package(OpenCV REQUIRED)
find_package(OpenSSL REQUIRED)

include_directories(include)

add_executable(lab8.3 
    main.cpp 
    image_hasher.cpp)

target_link_libraries(lab8.3 
    ${OpenCV_LIBS} 
    OpenSSL::SSL 
    OpenSSL::Crypto)