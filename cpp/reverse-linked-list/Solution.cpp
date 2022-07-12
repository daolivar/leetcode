class Solution {
public:
    ListNode *reverseList(ListNode *head) {
        if (head == nullptr) {
            return head;
        } else {
            ListNode *prev = nullptr;
            ListNode *next = nullptr;
            ListNode *curr = head;
            while (curr) {
                next = curr->next;
                curr->next = prev;
                prev = curr;
                curr = next;
            }
            head = prev;
        }
        return head;
    }
};
