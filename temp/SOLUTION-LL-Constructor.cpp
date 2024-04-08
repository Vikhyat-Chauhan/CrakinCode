#include <iostream>

using namespace std;

class Node
{
public:
    int value;
    Node *next;

    Node(int value)
    {
        this->value = value;
        next = nullptr;
    }
};

class LinkedList
{
private:
    Node *head;
    Node *tail;
    int length = 0;

public:
    LinkedList(int value)
    {
        Node *newNode = new Node(value);
        head = newNode;
        tail = newNode;
        length++;
    }

    ~LinkedList()
    {
        Node *temp = head;
        while (head)
        {
            head = head->next;
            delete temp;
            temp = head;
        }
    }

    void printList()
    {
        Node *temp = head;
        while (temp != nullptr)
        {
            cout << temp->value << endl;
            temp = temp->next;
        }
    }

    void getHead()
    {
        if (head == nullptr)
        {
            cout << "Head: nullptr" << endl;
        }
        else
        {
            cout << "Head: " << head->value << endl;
        }
    }

    void getTail()
    {
        if (tail == nullptr)
        {
            cout << "Tail: nullptr" << endl;
        }
        else
        {
            cout << "Tail: " << tail->value << endl;
        }
    }

    void getLength()
    {
        cout << "Length: " << length << endl;
    }

    void append(int value)
    {
        cout << "Appending to the list " << value << std::endl;
        Node *appendNode = new Node(value);
        tail->next = appendNode;
        tail = appendNode;
        length++;
    }

    void deleteLast()
    {
        std::cout << "deleteLast()" << std::endl;
        Node *previous = head;
        Node *current = head;
        if (length == 0)
        {
            cout << "List Empty " << std::endl;
            return;
        }
        else if (length == 1)
        {
            cout << "Deleted Last Node with Value " << head->value << std::endl;
            head = nullptr;
            tail = nullptr;
        }
        else
        {

            while (current->next)
            {
                previous = current;
                current = current->next;
            }
            previous->next = nullptr;
            tail = previous;
            cout << "Deleted Last Node with Value " << current->value << std::endl;
        }
        delete current;
        length--;
    }

    void prePend(int value)
    {
        Node *currentNode = new Node(value);
        if (length == 0)
        {
            tail = currentNode;
            currentNode->next = nullptr;
        }
        else
        {
            currentNode->next = head;
        }
        head = currentNode;
        length++;
        cout << "PrePended node with Value " << value << std::endl;
    }

    void deleteFirst()
    {
        if (head->next != nullptr)
        {
            Node *currentNode = head;
            delete currentNode;
            head = head->next;
        }
    }

    void insert()
    {
       
    }
};

int main()
{

    LinkedList *myLinkedList = new LinkedList(2);

    myLinkedList->append(3); // myLinkedList->append(3);myLinkedList->append(4);myLinkedList->append(5);

    cout << "\nLinked List:\n";
    myLinkedList->printList();
    myLinkedList->prePend(1);
    myLinkedList->printList();
}
