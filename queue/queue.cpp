#include "queue.hpp"
#include <iostream>

int main(){
    Queue q;
    q.enqueue(2);
    q.enqueue(3);
    q.enqueue(4);
    q.enqueue(6);

    std::cout<<q.get_front()<<std::endl;

    q.dequeue();
    q.dequeue();
    q.dequeue();

    std::cout<<q.get_front()<<std::endl;
}