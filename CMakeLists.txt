cmake_minimum_required(VERSION 3.16)

project(go-gptj VERSION 0.1.0 LANGUAGES CXX)

set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTORCC ON)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

set(GGML_BUILD_EXAMPLES ON CACHE BOOL "ggml: build examples" FORCE)

add_subdirectory(gpt4all-chat/ggml EXCLUDE_FROM_ALL)

set(TARGET binding)
add_library(${TARGET}
    binding.h binding.cpp
    gpt4all-chat/gptj.h gpt4all-chat/gptj.cpp
    gpt4all-chat/llmodel.h
    )

target_link_libraries(${TARGET} PUBLIC
    ggml
    ggml_utils
    )

