#include<iostream>

using namespace std;

class MyLinkedList {
public:
    struct LinkedNode {
        int val;
        LinkedNode* next;
        LinkedNode(int x): val(x), next(nullptr){};
    };

    MyLinkedList() {
        size = 0;
        dummyHead = new LinkedNode(0);
    }
    
    int get(int index) {
        if(index < 0 || index > size-1) {
            return -1;
        }

        LinkedNode* cur = dummyHead->next;
        while(index--) {
            cur = cur->next;
        }
        return cur->val;
    }
    
    void addAtHead(int val) {
        LinkedNode* newNode = new LinkedNode(val);
        newNode->next = dummyHead->next;
        dummyHead->next = newNode;
        size++;
    }
    
    void addAtTail(int val) {
        LinkedNode* cur = dummyHead;
        while(cur->next != nullptr) {
            cur = cur->next;
        }
        cur->next = new LinkedNode(val);
        size++;
    }
    
    void addAtIndex(int index, int val) {
        if(index > size) return;
        if(index < 0) index = 0;
        LinkedNode* cur = dummyHead;
        while(index--) {
            cur = cur->next;
        }
        LinkedNode* newNode = new LinkedNode(val);
        newNode->next = cur->next;
        cur->next = newNode;
        size++;
    }
    
    void deleteAtIndex(int index) {
        if(index < 0 || index > size-1) return;
        LinkedNode* cur = dummyHead;
        while(index--) {
            cur = cur->next;
        }
        cur->next = cur->next->next;
        size--;
    }

private:
    int size;
    LinkedNode* dummyHead;
};

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * MyLinkedList* obj = new MyLinkedList();
 * int param_1 = obj->get(index);
 * obj->addAtHead(val);
 * obj->addAtTail(val);
 * obj->addAtIndex(index,val);
 * obj->deleteAtIndex(index);
 */

class MyLinkedList2 {
public:
    struct LinkedNode {
        int val;
        LinkedNode* next;
        LinkedNode(int x): val(x), next(nullptr) {};
    };

    MyLinkedList2() {
        size = 0;
        head = nullptr;
    }
    
    int get(int index) {
        if(index < 0 || index > size-1) {
            return -1;
        }
        LinkedNode* cur = head;
        while(index--) {
            cur = cur->next;
        }
        return cur->val;
    }
    
    void addAtHead(int val) {
        LinkedNode* newNode = new LinkedNode(val);
        newNode->next = head;
        head = newNode;
        size++;
    }
    
    void addAtTail(int val) {
        LinkedNode* newNode = new LinkedNode(val);
        if(size == 0) {
            head = newNode;
        } else {
            LinkedNode* cur = head;
            while(cur->next != nullptr) {
                cur = cur->next;
            }
            cur->next = newNode;
        }
        size++;
    }
    
    void addAtIndex(int index, int val) {
        if(index <= 0) {
            addAtHead(val);
            return;
        }

        if(index == size) {
            addAtTail(val);
            return;
        }

        if(index > size) return;
        
        LinkedNode* newNode = new LinkedNode(val);
        LinkedNode* cur = head;
        while(--index) {
            cur = cur->next;
        } 
        newNode->next = cur->next;
        cur->next = newNode;
        size++;
    }
    
    void deleteAtIndex(int index) {
        if(index < 0 || index > size-1) {
            return;
        }

        if(index == 0) {
            LinkedNode* tmp = head;
            head = head->next;
            delete tmp;
            size--;
            return;
        }

        LinkedNode* cur = head;
        LinkedNode* tmp;
        while(--index) {
            cur = cur->next;
        }
        tmp = cur->next;
        if(tmp->next == nullptr) {
            cur->next = nullptr;
        } else{
            cur->next = cur->next->next;
        }
        size--;
        delete tmp;
    }

private:
    int size;
    LinkedNode* head;
};

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */

struct ListNode {
    int val;
    ListNode *next;
    ListNode(): val(0), next(nullptr) {}
    ListNode(int x): val(x), next(nullptr) {}
    ListNode(int x, ListNode *next): val(x), next(next) {}
};

class Solution {
public:
    /* 双指针法
    ListNode* reverseList(ListNode* head) {
        ListNode* cur = head;
        ListNode* pre = nullptr;
        ListNode* tmp;
        while(cur != nullptr) {
            tmp = cur->next;
            cur->next = pre;
            pre = cur;
            cur = tmp;
        }
        return pre;
    }
    */

    /* 递归法1 
    ListNode* reverse(ListNode* pre, ListNode* cur) {
        if(cur == nullptr) {
            return pre;
        }
        ListNode* tmp = cur->next;
        cur->next = pre;
        return reverse(cur, tmp);
    }    

    ListNode* reverseList(ListNode* head) {
        return reverse(nullptr, head);
    }
    */

    /* 递归法2 */
    ListNode* reverseList(ListNode* head) {
        if(head == nullptr || head->next == nullptr) {
            return head;
        }
        ListNode* last = reverseList(head->next);
        head->next->next = head; // 不能是last->next = head, 因为last是反转后的头结点
        head->next = nullptr;
        return last;
    }
};

int main() {
}