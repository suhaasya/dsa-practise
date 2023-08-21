#include <stdio.h>
#include <iostream>
using namespace std;

struct Rectangle{
    int length;
    int breadth;
};

void initialise(struct Rectangle *r,int l,int b){
    r->length=l;
    r->breadth=b;
}

int area(struct Rectangle r){
    return r.length*r.breadth;
}

int perimeter(struct Rectangle r){
    return 2*(r.length+r.breadth);
}

int main(){
    Rectangle r = {0,0};

    int l,b;
    printf("Enter length and Breadth: ");
    cin>>l>>b;

    initialise(&r,l,b);

    int ar = area(r);
    int peri = perimeter(r);

    cout<<"area"<<ar<<endl;
    cout<<"perimter"<<peri<<endl;

    return 0;
}