SOURCE_FILES = ../main.cpp ../hello.c ../hello2.cpp
OBJ_FILES = main.o hello.o hello2.o

CXXFLAGS = -fPIC -std=c++11 -I../

all: shared static

clean:
	rm ../*.so ../*.a *.o ../liblinux-shared-example.so ../liblinux-shared-example.a

shared:
	g++ -shared ${CXXFLAGS} ${SOURCE_FILES} -o ../liblinux-shared-example.so

static:
	g++ ${CXXFLAGS} -c ${SOURCE_FILES}
	ar rcs ../liblinux-shared-example.a ${OBJ_FILES} 
