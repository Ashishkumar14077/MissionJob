#ifndef STACK_HPP
#define STACK_HPP
#include <iostream>

//push(int)
//pop()
//peak()
//is_empty()
//is_full()

const int STACK_SIZE = 5;
class Stack {
    int arr[STACK_SIZE];
    int idx;
public:
    Stack();
    ~Stack();
    void push(int value);
    int pop();
    int peak();
    bool is_empty();
    bool is_full();
};
Stack::Stack(){
    std::cout<<"Constructor called...."<<"\n";
    idx = 0;
}

Stack::~Stack(){
    std::cout<<"Distructor called...\n";
}

void Stack::push(int val){
    if (is_full()){
        std::cout<<"Stack is full"<<std::endl;
        return;
    }
    arr[idx] = val;
    idx++;
}

int Stack::pop(){
    if(is_empty()){
        std::cout<<"stack is empty"<<std::endl;
        return -1;
    }
    return arr[idx--]; 
}

int Stack::peak(){
    if(is_empty()){
        std::cout<<"stack is empty"<<std::endl;
        return -1;
    }
    return arr[idx-1];
}

bool Stack::is_empty(){
    if (idx == 0){
        return true;
    } 
    return false;
}

bool Stack::is_full(){
    if(idx==STACK_SIZE){
       // std::cout<<"stack is full \n";
        return true;
    }
    return false;
}
#endif