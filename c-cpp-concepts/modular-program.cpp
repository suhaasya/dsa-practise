#include <stdio.h>
#include <iostream>
using namespace std;


int area(int l,int b){
    return l*b;
}

int perimeter(int l,int b){
    return 2*(l+b);
}

int main(){
    int l=0,b=0;
    printf("Enter length and Breadth: ");
    cin>>l>>b;

    int ar = area(l,b);
    int peri = perimeter(l,b);

    cout<<"area"<<ar<<endl;
    cout<<"perimter"<<peri<<endl;

    return 0;
}