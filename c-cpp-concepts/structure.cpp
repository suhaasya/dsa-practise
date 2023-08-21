#include <stdio.h>
#include <iostream>
using namespace std;

struct Rectangle{
    int lenght;
    int breadth;
};

int main(){
    struct Rectangle r;
    r.lenght = 5;
    r.breadth = 10;
    // printf("%d", r.lenght);
    cout<<r.lenght<<endl;
    return 0;
}