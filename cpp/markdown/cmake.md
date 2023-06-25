# cmake 

https://cmake.org/cmake/help/latest/guide/tutorial/A%20Basic%20Starting%20Point.html#exercise-1-building-a-basic-project

 CMakeLists.txt 를 만들어준다. 

- cmake_minimum_required(VERSION 3.10)
- project(Hello)
- set(CMAKE_CXX_STANDARD 11)
- set(CMAKE_CXX_STANDARD_REQUIRED True)
- add_executable(Hello hello.cpp)


build 폴더에서

cmake .. 
cmake --build .
