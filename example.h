#ifndef EXAMPLE_H__
#define EXAMPLE_H__

#ifdef WINDOWS
   #define API __declspec(dllexport)
#else
    #define API extern
#endif

#ifdef __cplusplus
extern "C" {
#endif // __cplusplus

   API int Sum(int a, int b);
   API int SumAll(const int* nums, int n);

   // Implemented in hello.c (C)
   void SayHello(const char* s);
   // Implemented in hello2.cpp (C++)
   void SayHello2(const char* s);
   // Implemented in hello.go (Go)
   void SayHelloGo(/*const*/ char* s);

#ifdef __cplusplus
}
#endif // __cplusplus
#endif // EXAMPLE_H__
