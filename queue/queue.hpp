#ifndef QUEUE_HPP
#define QUEUE_HPP
#include <iostream>

//enqueue(int)
//dequeue()
//front()
//is_empty()
//is_full()

const int QUEUE_SIZE = 5;
class Queue {
    int arr[QUEUE_SIZE];
    int front, rear;
public:
    Queue();
    ~Queue();
    void enqueue(int value);
    int dequeue();
    int get_front();
    bool is_empty();
    bool is_full();
};
Queue::Queue(){
    std::cout<<"Constructor called...."<<"\n";
    front = 0;
    rear = 0;
}

Queue::~Queue(){
    std::cout<<"Distructor called...\n";
}

void Queue::enqueue(int val){
    if(is_full()){
        return;
    }
    arr[rear] = val;
    rear++;
}

int Queue::dequeue(){
    if(is_empty()){
        return -1;
    }
    return arr[front++];
}

int Queue::get_front(){
    if(is_empty()){
        std::cout<<front<<"\n";
        return -1;
    }
    return arr[front];
}

bool Queue::is_empty(){
   if(front==rear){
    return true;
   }
   return false;
}

bool Queue::is_full(){
    if(rear==QUEUE_SIZE){
        return true;
    }
    return false;
}
#endif