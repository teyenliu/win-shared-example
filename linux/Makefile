SOURCE_FILES = ../main.cpp ../hello.c ../hello2.cpp

CXXFLAGS = -fPIC -std=c++11 -shared -I../

all: shared

clean:
        rm ../*.so

shared:
        g++ -shared ${CXXFLAGS} ${SOURCE_FILES} -o linux-shared-example.so
