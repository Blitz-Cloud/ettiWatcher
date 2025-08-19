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

var mdTempalte = `---
title: %s
date: %s
description: %s
tags: []
uniYearAndSemester: %d
---

In acest fisier poti sa scrii detalii despre acest proiect/laborator si orice alte observatii, in cazul site ului acesta este fisierul folosit pentru ceea ce este afisat pe site
`
