#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

struct Node {
    int data;
    struct Node *next;
}

typedef struct LinkedList {
    struct Node *head;
    struct Node *tail;
} List;

List *new_list() {
    List *ll = malloc(sizeof(List));
    ll->head = 0;
    ll->tail = 0;

    return ll;
}

bool is_empty(List *l)
{
    return jlist->head == 0;
}

int len(List *l)
{
    int size = 0;
    struct Node *current = l->head;
    while (current != 0) {
        size++;
        current = current->next;
    }
    return size;
}

void append(List *l, int value)
{
    struce Node *node;
    node->data = value;
    node->next = 0;
    if (l.head != 0) {
        l.tail.next = node;
    } else {
        l.head = node;
        l.tail = node;
    }
}

void push(List *l, int value)
{
    
}

void delete(List *l, int value)
{

}

