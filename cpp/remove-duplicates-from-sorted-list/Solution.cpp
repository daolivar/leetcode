class Solution {
public:
    ListNode *deleteDuplicates(ListNode *head) {
        if (head == nullptr || head->next == nullptr) { return head; }
        ListNode *cur = head;
        while (cur->next != nullptr) {
            if (cur->val == cur->next->val) {
                ListNode *temp = cur->next;
                cur->next = cur->next->next;
                delete temp;
            } else {
                cur = cur->next;
            }
        }
        return head;
    }
};
