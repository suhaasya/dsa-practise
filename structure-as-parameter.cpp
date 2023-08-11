#include <stdio.h>
#include <iostream>
using namespace std;

struct Rectangle{
    int lenght;
    int breadth;
};
// call by value
void fun(struct Rectangle r){
    r.lenght=20;
    cout<<"lenght "<<r.lenght<<endl; 
    cout<<"breadth "<<r.breadth<<endl; 
}

// call by address
void fun2(struct Rectangle *p){
    p->lenght=20;
    cout<<"lenght "<<p->lenght<<endl; 
    cout<<"breadth "<<p->breadth<<endl; 
}

struct Rectangle *fun3(){
    struct Rectangle *p;
    p= new Rectangle;
    p->lenght = 15;
    p->breadth = 7;
    return p;
}

int main(){
    
    struct Rectangle r = {10,5};

    fun(r);
    fun2(&r);

    cout<<"lenght "<<r.lenght<<endl; 
    cout<<"breadth "<<r.breadth<<endl; 

    struct Rectangle *ptr = fun3();

    cout<<"lenght "<<ptr->lenght<<endl; 
    cout<<"breadth "<<ptr->breadth<<endl; 


    return 0;
}