#include <stdio.h>
#include <iostream>
using namespace std;

class Rectangle{
    public:
    int length;
    int breadth;


    void initialise(int l,int b){
        length=l;
        breadth=b;
    }

    int area(){
        return length*breadth;
    }

    int perimeter(){
        return 2*(length+breadth);
    }

};

int main(){
    Rectangle r = {0,0};

    int l,b;
    printf("Enter length and Breadth: ");
    cin>>l>>b;

    r.initialise(l,b);

    int ar = r.area();
    int peri = r.perimeter();

    cout<<"area"<<ar<<endl;
    cout<<"perimter"<<peri<<endl;

    return 0;
}