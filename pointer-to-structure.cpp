#include <stdio.h>
#include <stdlib.h>
#include <iostream>
using namespace std;

struct Rectangle{
    int lenght;
    int breadth;
};

int main(){
    // Rectangle r = {10,5};
    // cout<<r.lenght<<endl;
    // cout<<r.breadth<<endl;

    // Rectangle *p=&r;
    // cout<<p->lenght<<endl;
    // cout<<p->breadth<<endl;

    Rectangle *p;
    p = new Rectangle;
    p->lenght = 8;
    p->breadth = 6;
    
    cout<<p->lenght<<endl;
    cout<<p->breadth<<endl;

    return 0;
}