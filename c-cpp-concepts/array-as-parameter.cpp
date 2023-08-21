#include <stdio.h>
#include <iostream>
using namespace std;


void fun(int *A,int n){
    for(int i=0;i<n;i++){
        cout<<A[i]<<" ";
    }
    A[0] = 15;
}

int * fun2(int size){
    int *p;
    p = new int[size];
    for(int i=0;i<size;i++){
        p[i]=i+1;
    }
    return p;
}


int main(){
   int a[]={2,4,5,6,7};
   int n = 5;
   
   fun(a,n);

   for(int x:a){
    cout<<x<<endl;
   }
   int *ptr = fun2(n);

    for(int i=0;i<n;i++){
        cout<<ptr[i]<<" ";
    }

    return 0;
}