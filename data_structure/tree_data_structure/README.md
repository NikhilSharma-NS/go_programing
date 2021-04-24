##Part:1

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
       root
    P
Q      Q  
     Q    Q
3.6) Descendants : node P is an ancestor of Q
Q is a descendant of P
3.6) Depth : Lenth of path from root to node
           25
    20            30
18       22     28     70
     21     24

depth of node 22 is 2: (25->20->22)     
3.7) Height :  Lenth of path from that node to deepest node
height of node 22 is 2:(22->21 and 24)
3.8) Skew trees
a tree that has only left or only right nodes are called left skew tree and right skew tree



##Part:2

1) What is binary tree
2) Types of Binary tree
3) Properties of Binary tree
4) Structure of Binary tree

1)  What is binary tree

Binary trees are the one in which a parent node have at most 2 children . i.e, either 0,1,2 children only


             1
      2              3
                  4     5
               6     7