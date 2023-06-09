// in disjoint set the key data structure is an array where the index is the node and the value is the parent
// key choice: is that value of the parent the root node or just the parent?
// if it is the root node -> quick find which means finding the root is O(1) but the trade-off is union 
// because O(n) because every union you have to change every node in each tree
#include <vector>
#include <iostream>

using std::vector;
using std::endl;
using std::boolalpha;
using std::cout;

class UnionFind {
    public:
        // O(n)
        UnionFind(int sz) : root(sz) { // : defines the initializer list where root is initialized to value of sz
            // set node i as its own parent
            for (int i = 0; i < sz; i++) {
                root[i] = i;
            }
        }

        // O(1)
        int find(int x) {
            return root[x];
        }

        // O(n)
        void unionNodes(int x, int y){
            int rootX = find(x);
            int rootY = find(y);
            // they are not connected - so we need to connect them
            if (rootX != rootY) {
                for (int i = 0; i < root.size(); i++){
                    if (root[i] == rootY) {
                        root[i] = rootX;
                    }
                }
            }
        }

        // O(1)
        bool connected(int x, int y) {
            return find(x) == find(y);
        }

    private:
        vector<int> root;
};

// Test Case
int main() {
    // for displaying booleans as literal strings, instead of 0 and 1
    cout << boolalpha;
    UnionFind uf(10);
    // 1-2-5-6-7 3-8-9 4
    uf.unionNodes(1, 2);
    uf.unionNodes(2, 5);
    uf.unionNodes(5, 6);
    uf.unionNodes(6, 7);
    uf.unionNodes(3, 8);
    uf.unionNodes(8, 9);
    cout << uf.connected(1, 5) << endl;  // true
    cout << uf.connected(5, 7) << endl;  // true
    cout << uf.connected(4, 9) << endl;  // false
    // 1-2-5-6-7 3-8-9-4
    uf.unionNodes(9, 4);
    cout << uf.connected(4, 9) << endl;  // true

    return 0;
}