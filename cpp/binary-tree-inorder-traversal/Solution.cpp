#include <vector>
using namespace std;

class Solution {
public:
    void inorder(TreeNode *root, vector<int> &v) {
        if (root == nullptr) { return; }
        inorder(root->left, v);
        v.push_back(root->val);
        inorder(root->right, v);
    }

    vector<int> inorderTraversal(TreeNode *root) {
        vector<int> inorderValues;
        inorder(root, inorderValues);
        return inorderValues;
    }
};
