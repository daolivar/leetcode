class Solution
{
public:

    bool isSymmetric(TreeNode *root) { 
        return mirrorNodes(root->left, root->right);
    }

    bool mirrorNodes(TreeNode *left, TreeNode *right) {
        if (left == nullptr || right == nullptr) {
            return left == right;
        }
        if (left->val != right->val) {
            return false;
        }
        return mirrorNodes(left->left, right->right) && mirrorNodes(left->right, right->left);
    }
};
