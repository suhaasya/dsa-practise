#include <stdio.h>
#include <iostream>
using namespace std;

int add(int a, int b){
    return a + b;
}

int main(){
    int num1 = 10,num2=5,sum;
    
    sum = add(num1,num2);
    cout<<sum<<endl;
    return 0;
}