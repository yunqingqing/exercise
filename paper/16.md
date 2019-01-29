# 删除链表的节点

> 在O(1)时间内删除链表节点
> 给定单向链表的头指针和一个节点指针，定义一个函数在O(1)时间内删除该节点。链表节点和函数定义如下:

```c++
struct ListNode
{
    int m_nValue;
    ListNode* m_pNext;
}
void DeleteNode(ListNode**pListHead, ListNode*pToBeDeleted)
```

> 删除链表中的重复节点
> 在一个排序的链表中，1->2->3->3->4->4->5，删除重复节点后1->2->5