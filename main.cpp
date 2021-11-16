// The functions contained in this file are pretty dummy
// and are included only as a placeholder. Nevertheless,
// they *will* get included in the shared library if you
// don't remove them :)
//
// Obviously, you 'll have to write yourself the super-duper
// functions to include in the resulting library...
// Also, it's not necessary to write every function in this file.
// Feel free to add more files in this project. They will be
// included in the resulting library.

#include "example.h"

API int Sum(int a, int b){
   return a + b;
}

API int SumAll(const int* nums, int n){
   int r = 0;
   for (int i = 0; i < n; i++){
      r += nums[i];
   }
   return r;
}
