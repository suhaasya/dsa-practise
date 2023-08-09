#include<iostream>
#include<stdio.h>

using namespace std;

int main(){
    // int a = 10;
    // int *p;
    // p=&a;

    // cout<<a<<endl;
    // cout<<p<<endl;
    // cout<<*p<<endl;
    // cout<<&a<<endl;
    // int A[5] = {2,4,5,6,8};
    int *p;
    p = new int[5];
    p[0]=1;
    p[1]=3;
    p[2]=5;
    p[3]=8;
    p[4]=4;
    // p = &A; don't do this when assigning pointer with array
    // p = A;
    for (int i = 0; i < 5; i++)
    {
        cout<<p[i]<<endl;
        cout<< sizeof(p[i])<<endl;
    }
    
    delete [] p;
    return 0;
}