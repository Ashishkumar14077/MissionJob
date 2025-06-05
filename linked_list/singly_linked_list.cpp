#include <iostream>
using namespace std;

//create structure of Node
struct Node{
    int data;
    Node *next;
};

int main(){
    //create instanse of Node 
    Node n;
    //insert value in first node
    n.data = 10;
    //cout<<n.next<<"\n"; // 0x1048c312c
   // cout<<n.data<<"\n"; // 10

    //create head to point self
    Node *head = new Node;
    //head of next point to the n 
    head = &n;
    //cout<<head->next<<"\n"; // 0x1048c312c
    //cout<<head->data<<"\n"; // 10

    //create tail 
    Node *tail = new Node;
    //tail point to head in initial
    tail = head;

    //create new node
    Node n1;
    n1.data = 20;
    //cout<<n1.data<<"\n";
    n1.next = nullptr;
    //head->next = &n1; //corrrect
    tail->next = &n1;
    tail = &n1;

    // create another node
    Node n2;
    n2.data = 30;
    //cout<<n2.data<<"\n";
    tail->next = &n2;
    tail = &n2;
    n2.next = nullptr;

    //create another node
    Node n3;
    n3.data = 40;
    tail->next = &n3;
    tail = &n3;
    n3.next = nullptr;

    //create new node
    Node n4;
    n4.data = 50;
    tail->next = &n4;
    tail = &n4;
    n4.next = nullptr;

    Node *x = new Node;
    x = head;

    while(x!=nullptr){
        if(x->data == 40){
            Node n5;
            n5.data = 45;
            Node *y = new Node;
            y = x->next;
            x->next = &n5;
            n5.next = y;
            break;
        }
        x = x->next;
    }

    cout<<"Iterate"<<"\n";
    x = head;
    while(x!=nullptr){
        cout<<x->data<<"\n";
        x = x->next;
    }
}