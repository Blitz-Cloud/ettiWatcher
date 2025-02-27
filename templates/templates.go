package templates

var CppTemplate = `#include <iostream>

using namespace std;

int main()
{
  cout << "Hello World" << endl;
  return 0;
}`

var CTemplate = `#include <stdio.h>

int main(){
  printf("Hello world");
  return 0;
}
`
var CMakeForC = `cmake_minimum_required(VERSION 3.29)
project(%s C)

set(CMAKE_C_STANDARD 11)

add_executable(%s main.c)

`
var CMakeForCpp = `cmake_minimum_required(VERSION 3.29)
project(%s)

set(CMAKE_CXX_STANDARD 20)

add_executable(%s main.cpp)`
