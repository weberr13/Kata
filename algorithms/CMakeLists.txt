cmake_minimum_required (VERSION 3.0)
ENABLE_LANGUAGE(CXX)

project (Kata_Algorithms)
set(PROJECT_SRC ${Kata_Algorithms_SOURCE_DIR}/src)
set(TEST_SRC ${Kata_Algorithms_SOURCE_DIR}/test)
set(THIRDPARTY_SRC ${Kata_Algorithms_SOURCE_DIR}/thirdParty)
SET(PLATFORM_LINK_LIBRARIES rt)
SET(CMAKE_CXX_FLAGS "-Wall -rdynamic -Wunused -std=c++14 -pthread -D_GLIBCXX_USE_NANOSLEEP")

file(GLOB SRC_FILES ${Kata_Algorithms_SOURCE_DIR}/src/*.h ${Kata_Algorithms_SOURCE_DIR}/src/*.hpp ${Kata_Algorithms_SOURCE_DIR}/src/*.cpp)
SET(THIS_LIB Algorithmss)
SET(MAJOR_VERSION 0)
SET(MINOR_VERSION 1)
SET(BUILD_NUMBER 0)
SET(VERSION ${MAJOR_VERSION}.${MINOR_VERSION}.${BUILD_NUMBER})

include_directories(${PROJECT_SRC})
add_library(${THIS_LIB} SHARED ${SRC_FILES})
SET(VERSION_STRING ${VERSION})
SET_TARGET_PROPERTIES(${THIS_LIB} PROPERTIES LINKER_LANGUAGE CXX SOVERSION ${VERSION} )

set(GTEST_DIR ${THIRDPARTY_SRC}/googletest-release-1.8.0/googletest)
set(GTEST_INCLUDE_DIRECTORIES ${Kata_Algorithms_SOURCE_DIR}/thirdParty/googletest-release-1.8.0/googletest/include/ ${Kata_Algorithms_SOURCE_DIR}/thirdParty/googletest-release-1.8.0/googletest ${GTEST_DIR}/src)
include_directories(${GTEST_INCLUDE_DIRECTORIES})
add_library(gtest_180_lib ${Kata_Algorithms_SOURCE_DIR}/thirdParty/googletest-release-1.8.0/googletest/src/gtest-all.cc)
set_target_properties(gtest_180_lib PROPERTIES COMPILE_DEFINITIONS "GTEST_HAS_RTTI=0")
enable_testing(true)
include_directories(test)
include_directories(${PROJECT_SRC})
file(GLOB TEST_SRC_FILES "test/*.cpp")
SET(CMAKE_CXX_FLAGS "-Wall -Wno-sign-compare -Wunused -std=c++14 -pthread -D_GLIBCXX_USE_NANOSLEEP")

add_executable (UnitTests thirdParty/test_main.cpp ${TEST_SRC_FILES})
target_link_libraries (UnitTests ${THIS_LIB} gtest_180_lib)

set_target_properties(${test} PROPERTIES COMPILE_FLAGS "-isystem -pthread ")

