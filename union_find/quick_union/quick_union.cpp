// in disjoint set the key data structure is an array where the index is the node and the value is the parent
// key choice: is that value of the parent the root node or just the parent?
// if it is the parent then -> quick union
// we can optimise slightly where given two trees we always have the root of one tree connect to the root of another
// this means if it is a single node being connected they are all connected to the root - but it does mean
// if there is a tree being connected you have to search through

#include <vector>
#include <iostream>

using std::vector;
using std::endl;
using std::boolalpha;
using std::cout;

class UnionFind {
    public:
        // O(n)
        UnionFind(int sz) : root(sz) {
            for (int i = 0; i < sz; ++i) {
                root[i] = i;
            }
        }

        // O(n)
        int find(int x) {
            while (x != root[x]) {
                x = root[x];
            }
            return x;
        }

        // O(n)
        int find(int x) {
            while (x != root[x]) {
                x = root[x];
            }
            return x;
        }

        // O(n) - O(1 itself but O(n) because of find
        void unionSet(int x, int y){
            int rootX = find(x);
            int rootY = find(y);
            if (rootX != rootY) {
                root[rootX] = rootY;
            }
        }

        // O(n) - O(1 itself but O(n) because of find
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
    uf.unionSet(1, 2);
    uf.unionSet(2, 5);
    uf.unionSet(5, 6);
    uf.unionSet(6, 7);
    uf.unionSet(3, 8);
    uf.unionSet(8, 9);
    cout << uf.connected(1, 5) << endl;  // true
    cout << uf.connected(5, 7) << endl;  // true
    cout << uf.connected(4, 9) << endl;  // false
    // 1-2-5-6-7 3-8-9-4
    uf.unionSet(9, 4);
    cout << uf.connected(4, 9) << endl;  // true

    return 0;
}