/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//var max int;
//var n   int;
func maxDepth(root *TreeNode) int {
    
    if(root == nil){
        return 0;
    }
    l:=0;
    r:=0;
    if(root.Left!=nil){
        l=maxDepth(root.Left);
    }
    if(root.Right != nil){
        r=maxDepth(root.Right);
    }
    
    if(l>r){
        return l+1;
    }else{
        return r+1;
    }
    
}
/* 
    if(root==nil){
        return 0;
    }
    max =0;
    n=1;
    visit(root, n);
    return max;
}
func visit(head * TreeNode, n int) {
    if(head != nil){
        if(n>max){
            max = n;
        }
       fmt.Println(n);
       visit(head.Left, n+1);

       visit(head.Right, n+1);
       
    }else{
        
    }
}*/
