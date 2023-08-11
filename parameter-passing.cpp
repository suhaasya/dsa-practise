#include <stdio.h>
#include <iostream>
using namespace std;


int add(int a, int b){
    int c;
    c = a+b;
    return c;
}

int swap(int *x, int *y){
    int temp;
    temp = *x;
    *x=*y;
    *y=temp;
}

int swap2(int &x, int &y){
    int temp;
    temp = x;
    x=y;
    y=temp;
}

int main(){
    
    int num1=10,num2=15,sum;

    sum = add(num1,num2);
    // swap(&num1,&num2);
    swap2(num1,num2);
    cout<<"Sum is "<<sum<<endl;
    cout<<"num1 "<<num1<<endl;
    cout<<"num2 "<<num2<<endl;
    return 0;
}