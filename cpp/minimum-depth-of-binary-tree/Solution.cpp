#include <algorithm>
using namespace std;

class Solution {
public:
    int minDepth(TreeNode *root) {
        if (root == nullptr) {
            return 0;
        }
        if (root && !root->left) {
            return 1 + minDepth(root->right);
        }
        if (root && !root->right) {
            return 1 + minDepth(root->left);
        }
        int left = minDepth(root->left);
        int right = minDepth(root->right);
        return min(left, right) + 1;
    }
};
