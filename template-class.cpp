#include<iostream>
using namespace std;

template<class T>
class Arithmetic{
    private:
        T a;
        T b;
    
    public:
        Arithmetic(T a,T b);
        T add();
        T sub();
};

template <class T>
Arithmetic<T>::Arithmetic(T a,T b){
    this->a=a;
    this->b=b;
}

template <class T>
T Arithmetic<T>::add(){
    return a+b;
}

template <class T>
T Arithmetic<T>::sub(){
    return a-b;
}


int main(){
    Arithmetic ar(10.9,5.8);
    cout<<"ADD "<<ar.add()<<endl;
    cout<<"SUB "<<ar.sub()<<endl;
    return 0;
}