#include <iostream>

using namespace std;


class Node { 
    public:
        int value;
        Node* next;

        Node(int value) {
            this->value = value;
            next = nullptr;
        }
}; 


class LinkedList {
    private:
        Node* head;
        Node* tail;
        int length;

    public:
        LinkedList(int value) {
            Node* newNode = new Node(value);
            head = newNode;
            tail = newNode;
            length = 1;
        }

        ~LinkedList() {
            Node* temp = head;
            while (head) {
                head = head->next;
                delete temp;
                temp = head;
            }
        }

        void printList() {
            Node* temp = head;
            while (temp != nullptr) {
                cout << temp->value << " -> ";
                temp = temp->next;
            }
            cout << "nullptr \n";
        }
        
        void printListAdv() {
            Node* temp = head;
            while (temp != nullptr) {
                cout << temp->value << " -> ";
                temp = temp->next;
            }
            cout << "nullptr | H ";
            if(head != nullptr)
            cout<<head->value;
            cout<<" | T ";
            if(tail != nullptr)
            cout<<tail->value;
            cout<<" | L "<<length <<"\n";
        }

        void getHead() {
            if (head == nullptr) {
                cout << "Head: nullptr" << endl;
            } else {
                cout << "Head: " << head->value << endl;
            }
        }

        void getTail() {
            if (tail == nullptr) {
                cout << "Tail: nullptr" << endl;
            } else { 
                cout << "Tail: " << tail->value << endl;
            }  
        }

        void getLength() {
            cout << "Length: " << length << endl;
        }

        void appendList(int value) {
            Node* newNode = new Node(value);
            newNode->next = nullptr;
            tail->next = newNode;
            tail = newNode;
            length++;
        }
        
        void deleteLast() {
            /*Edge case the list is already empty*/
            if(length == 0){
                cout << "List Empty" << endl;
                return;
            }
            else if(length == 1){
                cout << "Emptying List" << endl;
                delete tail;
                head = nullptr;
                tail = nullptr;
            }
            else{
            Node* current = head;
            /* Length Dependent method
            for(int i=0;i<length-2;i++){
                current = current->next;
            }*/
            /*Lenght Independent method*/
            while((current->next)->next != nullptr){
                current = current->next;
            }
            delete tail;
            current->next = nullptr;
            tail = current;
            }
            length--;
        }
        
        void prependList(int value){
            Node* newnode = new Node(value);
            /*Edge Cases*/
            if(length == 0){
                newnode->next = nullptr;
                tail = newnode;
            }else{
                newnode->next = head;
            }
            head = newnode;
            length++;
        }
        
        void deleteFirst(){
            /*Edge Cases*/
            if(length == 0){
                cout<<"List Empty"<<endl;
                return;
            }
            Node* temp = head;
            if(length == 1){
                cout<<"Emptying List"<<endl;
                head = nullptr;
                tail = nullptr;
            }else{
                head = head->next;
            }
            delete temp;
            length--;
        }
        
        Node* getItem(int index){
            /*Edge Cases*/
            if(length == 0){
                cout<<"List is Empty"<<endl;
                return nullptr;
            }
            if((index>length) || (index<0)){
                cout<<"Index out of Bound"<<endl;
                return nullptr;
            }else{
                Node* returnNode = head;
                for(int i=0; i<index;i++){
                    returnNode = returnNode->next;
                }
                return returnNode;
            }
        }
        
        bool setItem(int index, int value){
            Node* item = getItem(index);
            /*Edge Case*/
            if(item == nullptr)
            return false;
            item->value = value;
            return true;
        }
        
        bool insertItem(int index, int value){
            /*Edge Case*/
            if((index>=length) || (index<0)){
                cout<<"Index out of Bound"<<endl;
                return false;
            }
            if(index == 0){
                prependList(value);
                return true;
            }
            if(index == length-1){
                appendList(value);
                return true;
            }
            Node* insertNode = new Node(value);
            Node* itemBefore = getItem(index-1);
            insertNode->next = itemBefore->next;
            itemBefore->next = insertNode;
            length++;
            return true;
        }
        
        bool deleteNode(int index){
            /*Edge Case*/
            if((index>=length) || (index<0)){
                cout<<"Index out of Bound"<<endl;
                return false;
            }
            if(index == 0){
                deleteFirst();
                return true;
            }
            if(index == length-1){
                deleteLast();
                return true;
            }
            Node* itemBefore = getItem(index-1);
            Node* item = itemBefore->next;
            itemBefore->next = item->next;
            delete item;
            return true;
        }
        
        void reverseList(){
            Node* before = nullptr;
            Node* current = head;
            //Reversing Head and Tail after they are used for intilization
            Node* temp = tail;
            tail = head;
            head = temp;
            
            for(int i=0; i<length;i++){
                Node* after = current->next;
                current->next = before;
                //All this is for next Execution
                before = current;
                current = after;
            }
        }
};

int main() {
        
    LinkedList* myLinkedList = new LinkedList(1);

    myLinkedList->appendList(2);
    myLinkedList->appendList(3);
    myLinkedList->appendList(4);
    
    cout << "\nLinked List:\n";
    myLinkedList->printListAdv();
    myLinkedList->deleteNode(4);
    myLinkedList->printListAdv();
}