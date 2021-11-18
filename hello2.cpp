// hello2.cpp

#include <iostream>

extern "C" {
    #include "example.h"
}

void SayHello2(const char* s) {
     std::cout << s;
}
