import sys

class Node:
    def __init__(self,data):
        self.right=self.left=None
        self.data = data
class Solution:
    def insert(self,root,data):
        if root==None:
            return Node(data)
        else:
            if data<=root.data:
                cur=self.insert(root.left,data)
                root.left=cur
            else:
                cur=self.insert(root.right,data)
                root.right=cur
        return root

    def levelOrder(self,root):
        #Write your code here
        queue = []
        output = []
        if isinstance(root,Node):
            queue.append(root)
            
            while queue:
                current = queue.pop(0)
                output.append(current.data)
                if isinstance(current.left,Node):
                    queue.append(current.left)
                if isinstance(current.right,Node):
                    queue.append(current.right)
        print(*output)      
                
                

T=int(input())
myTree=Solution()
root=None
for i in range(T):
    data=int(input())
    root=myTree.insert(root,data)
myTree.levelOrder(root)
