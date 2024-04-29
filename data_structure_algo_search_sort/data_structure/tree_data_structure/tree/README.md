## Tree

### Introduction to Binary Tree


Full Binary Tree - Either has 0 or 2 Childern 

Compelte Binary Tree 
- All levels are completly Filled except the last level 
- The last level has all nodes as left as possible 

Perfect Binary Tree
- All leaf nodes are at same level 

Balanced Binary Tree
- Height of tree max Log(N) (n=number of nodes)

Degenerate Tree
- every nodes has sigle childern 

### Traversal of Binary Tree(BFS/DFS)

#### Inorder Traversal
- Left Root Right

#### Preorder Traversal
- Root Left Right
  
#### Postorder Traversal
- Left Right Root    

```
          1
    2          3
4      5    6     7
```

- Inorder : 4 2 5 1 6 3 7
- Preorder : 1 2 4 5 3 6 7
- PostOder : 4 5 2 6 7 3 1 
- Breadth First Search or BFS - 1 2 3 4 5 6 7
- Depth First Search or DFS - 4 5 6 7 2 3 1







