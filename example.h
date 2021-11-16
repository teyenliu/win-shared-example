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

#ifdef __cplusplus
}
#endif // __cplusplus
#endif // EXAMPLE_H__
