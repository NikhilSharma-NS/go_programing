## Part:1

1)Diffrence between Linear and hierarchical data structure
2)What is a tree
3)Basic terminologgies used in tree (root,parent,ancestor,descendant,sibling,depth,height,skew,tree)

1.1)Linear data Structure 

Arrays
Linked List
Stack
Queue

1.2)Hierarchical data Structure
Tree

2) What is tree

2.1) Similar to Linked list but instead of each node pointing to a simply to next node in a linear fashion; each node points to a number of nodes
2.2) Non linear data strucuture 
2.3) Used to represent hierarchial data

3) Terminologies

3.1) Root : It is the node that has no parent
3.2) Edge : It is a link from parent to its child
3.3) Leaf nodes : Nodes with no children
3.4) Sibling : Childern of same parent
3.5) Ancestors  : node P is an ancestor Q ,if there is a path from root to Q such that P appears in the path
```
       root
    P
Q      Q  
     Q    Q
```
3.6) Descendants : node P is an ancestor of Q
Q is a descendant of P
3.6) Depth : Lenth of path from root to node
```
           25
    20            30
18       22     28     70
     21     24
```
depth of node 22 is 2: (25->20->22)     
3.7) Height :  Lenth of path from that node to deepest node
height of node 22 is 2:(22->21 and 24)
3.8) Skew trees
a tree that has only left or only right nodes are called left skew tree and right skew tree



## Part:2

1) What is binary tree
2) Types of Binary tree
3) Properties of Binary tree
4) Structure of Binary tree

1)  What is binary tree

Binary trees are the one in which a parent node have at most 2 children . i.e, either 0,1,2 children only

```
             1
      2              3
                  4     5
               6     7

```

2) Types of binary tree

2.1) Strict 
2.2) Full/perfect
2.3) Complete

2.1) Strict Binay tree
Each node has exactly 2 child nodes or no nodes

```
   1
2     3
    4   5
  6   7
```

2.2) Full/perfect Binay tree
Each node has exactly 2 children and the leaf nodes are at same level

```
         1
    2        3
 4     5   6    7
```

2.3) Complete Binay tree
Every level except possibility last are completely filled and all the node are as far left as possible

```
           1
      2        3
   4     5   6    7
8    9

```
maximum number of nodes at a level are 2nd level
level 0 : 2^0 = 1 
level 1 : 2^1 = 2
level 2 : 2^2 = 4

Example:
```
           1
      2        3
   4     5   6    7

```

Maximum number of nodes in a tree: 2^(h+1)-1 -> 2^3-1=7

Maximum height of tree given n node : log2(n+1)-1 -> log2(7+1)-1=2

heights
```
null -> height = -1
```
```
1 -> height = 0
```
```
 1    -> height =1
2 3
``` 

Balanced Binary tree

difference between height of left and right subtree is not more than k ( mostly 1)

leftHeight - RightHeight = 1

```
             1
         2      3
      4      6     7
   8     9
```

3) Properties of Binary tree

3.1) Number of nodes in a full/perfect binary tree : 2^(n+1)-1
3.2) Number of nodes in a complete binary tree are between 2^n and 2^(n+1)-1
3.3) Minimum height of a binary tree is log2(n+1)-1 Or Floor (log2(n+1))

4) Structure of Binary tree

```                     (root)
                        data
       left                                        right
 (pointer to left node)                   (pointer to right node) 

```
In Go 
```
type TreeNode struct{
left   *TreeNode
right  *TreeNode
data int64
}
```
In Java
```
class TreeNode{
    int data;
    TreeNode left;
    TreeNode right;
}
```

## Part3


1) Binary Search Tree
2) Properties of Binary Search Tree


1) Binary Search Introduction

Binary search tree is the type of binary tree in which the value of left child is always less than the right child

It is mainly used in Binary search technique

2) Properties of Binary Search Tree

2.1) The left subtree of a node only contains key less than the nodes key
2.2) The right subtree of a node only contains key greater than the nodes key
2.3) Both right and left subtree must also be binary search tree

```
                    10
              5          14
          4     8    12      17       
```

## Part 4

Tree Traversal

1) Inorder
2) Preorder
3) Postorder

What is tree traversal

the process of visiting all the nodes of a tree

How can we visit nodes

starting the root node; we can perform 3 operation
1) Process current node ( we will call this opeartion as D)
2) Traverse to the left child node (Lets call it;L)
3) Traverse to the right child node (Lets call it:R)

How can we visit nodes?

1)LDR : ProcessLeftTree,ProcessCurrent,ProcessRightTree
2)LRD : ProcessLeftTree,ProcessRightTree,ProcessCurrent
3)DLR : ProcessCurrent,ProcessLeftTree,ProcessRightTree
4)DRL : ProcessCurrent,ProcessRightTree,ProcessLeftTree
5)RDL : ProcessRightTree,ProcessCurrent,ProcessLeftTree
6)RLD : ProcessRightTree,ProcessLeftTree,ProcessCurrent

Classification of Traversal

1) Preorder (DLR) Traversal
2) Inorder  (LDR) Traversal
3) Postorder (LRD) Traversal
4) Level order Traversal (inspired by BFS of Graph traversal)

Example : 
``` 
                               1
                    2                    3
            4               5    6                7

```

Preorder Traversal

DLR
```
Step1) Process root 
o/p : 1
Step2)Process Left Tree
2.1) 2,left,right
o/p : 1,2
2.2) process left(4)
o/p : 1,2,4
2.3) Process right (5)
o/p : 1,2,4,5
Step3) Process Right
3.1) 3,left,right
o/p : 1,2,4,5,3
3.2) process left(6)
o/p : 1,2,4,5,3,6
3.3) Process right(7) 
o/p : 1,2,4,5,3,6,7

```

Inorder Traversal
LDR
```
Step1) Process Left
1.1) left,root,right -> left,2,right
1.2) 4,2,right
o/p : 4
1.3) 2,right
o/p : 4,2
1.4) right(5)
o/p : 4,2,5
Step2) Process root(1)
o/p: 4,2,5,1
Step3) Process right
3.1) right
3.2)left(6),root,right
o/p: 4,2,5,1,6
3.3)root(3),right
o/p: 4,2,5,1,6,3
3.4)right(7)
o/p: 4,2,5,1,6,3,7

```

Post Traversal
LRD
```
Step1) Process Left
1.1)left(4),right,root
o/p: 4
1.2)right,node
o/p: 4,5
1.3) node
o/p: 4,5,2
Step2) Process Right
2.1)left(6),right,root
o/p:4,5,2,6
2.2)right,root
o/p:4,5,2,6,7
2.3)root(3)
o/p:4,5,2,6,7,3
Step3) Process root(1)
o/p:4,5,2,6,7,3,1

```


