#include "stack.hpp"

int main(){
    Stack s;
    s.push(4);
    s.push(2);
    s.push(3);
    std::cout<<"peak"<<s.peak()<<"\n";
    s.pop();
    std::cout<<"peak"<<s.peak()<<"\n";
    s.pop();
    s.pop();
    s.pop();
}
