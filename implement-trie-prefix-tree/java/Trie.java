/**
 * Trie data structure implementation.
 */
public class Trie {
    /** The root node of the trie. */
    public TrieNode root;

    /**
     * Initializes an empty trie with the root node.
     */
    public Trie() {
        this.root = new TrieNode();
    }

    /**
     * Inserts a word into the trie.
     * 
     * Time Complexity: O(n), where 'n' is the length of the word.
     * Space Complexity: O(n), considering the worst-case scenario of creating nodes
     * for each character.
     * 
     * @param word The word to insert.
     */
    public void insert(String word) {
        TrieNode current = root;
        for (char c : word.toCharArray()) {
            current.children.putIfAbsent(c, new TrieNode());
            current = current.children.get(c);
        }
        current.isEndOfWord = true;
    }

    /**
     * Searches for a word in the trie.
     * 
     * Time Complexity: O(n), where 'n' is the length of the word.
     * Space Complexity: O(1).
     * 
     * @param word The word to search for.
     * @return true if the word is in the trie, false otherwise.
     */
    public boolean search(String word) {
        TrieNode current = root;
        for (char c : word.toCharArray()) {
            if (!current.children.containsKey(c)) {
                return false;
            }
            current = current.children.get(c);
        }
        return current.isEndOfWord;
    }

    /**
     * Checks if there is any word in the trie that starts with the given prefix.
     * 
     * Time Complexity: O(n), where 'n' is the length of the prefix.
     * Space Complexity: O(1).
     * 
     * @param prefix The prefix to check.
     * @return true if there is any word with the given prefix, false otherwise.
     */
    public boolean startsWith(String prefix) {
        TrieNode current = root;
        for (char c : prefix.toCharArray()) {
            if (!current.children.containsKey(c)) {
                return false;
            }
            current = current.children.get(c);
        }
        return true;
    }
}

/**
 * TrieNode class representing a node in the trie.
 */
class TrieNode {
    /** Indicates whether the node represents the end of a word. */
    public boolean isEndOfWord;

    /** Mapping of characters to their corresponding child nodes. */
    public Map<Character, TrieNode> children;

    /**
     * Initializes a trie node.
     */
    public TrieNode() {
        this.isEndOfWord = false;
        this.children = new HashMap<>();
    }
}
