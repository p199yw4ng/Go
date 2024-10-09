// 1. 两数之和
func TwoSum(nums []int, target int) []int {
    diff:=make(map[int]int)
    for i,v:=range nums{
        if j,ok := diff[v];ok{
            
            return []int{j,i}
        }
        diff[target-v]=i
    }
    
    return []int{0,0}
}

//49. 字母异位词分组
func groupAnagrams(strs []string) [][]string {
    mapAn := make(map[string][]string)

    for _, str := range strs{
        s := []byte(str)
        sort.Slice(s, func(i,j int)bool{return s[i]<s[j]})
        sorted:=string(s)
        mapAn[sorted] = append(mapAn[sorted],str)
    }
    var ans [][]string
    fmt.Println(mapAn)
    for _,v := range mapAn{
        ans = append(ans,v)
    }
   
    return ans
}
// 2. 移动零
func moveZeroes(nums []int)  {
    i,j,n := 0,0,len(nums)
    for ;i<n;i++{
        if nums[i]!=0{
            nums[j]=nums[i]
            j++
        }
    }
    for ;j<n;j++{
        nums[j]=0
    }
}
// 3. 无重复字符的最长子串
func LengthOfLongestSubstring(s string) int {
    n := len(s)
    windowLen := n
    for windowLen>=0{
        for i:=0;i+windowLen<=n;i++{
            if noDuplicateString(s[i:i+windowLen]){
                return windowLen
            }
        }
        windowLen--
    }
    return 0
}

func noDuplicateString(s string) bool {
    mapS := make(map[rune]bool)
    for _, v := range s{
        if mapS[v]{
            return false
        }else{
            mapS[v] = true
        }
    }
    return true
}
// 5. 统计子串和为k的子数组的个数
func subarraySum(nums []int, k int) int {
    
    ans:=0
    // 1.暴力
    for i:=0;i<len(nums);i++{
        sum:=0
        for j:=i;j<len(nums);j++{
            sum+=nums[j]
            if sum==k{
                ans++
            }
        }
    }
	// 前缀和
	// count, pre := 0, 0
    // m := map[int]int{}
    // m[0] = 1
    // for i := 0; i < len(nums); i++ {
    //     pre += nums[i]
    //     if _, ok := m[pre - k]; ok {
    //         count += m[pre - k]
    //     }
    //     m[pre] += 1
    // }
    // return count
    return ans
}

// 6. 最大子数组的数
func maxSlidingWindow(nums []int, k int) []int {
    ans := make([]int,0)
    maxQueen := make([]int,0)
    for i:=0;i+k<=len(nums);i++{
        //初始化max
        if i==0{
            for j:=0;j<k;j++{
                maxQueen = mySortInsert(maxQueen,nums[j])
            }    
        }else{
            //移除
            maxQueen = mySortDelete(maxQueen,nums[i-1])
            //加入
            maxQueen = mySortInsert(maxQueen,nums[i+k-1])
        }
        //加入 最大数
        ans = append(ans,maxQueen[0])
    }
    return ans
}
func mySortInsert(nums []int, k int) []int{
    for i,v:=range nums{
        if v<=k{
            // 插入k
            return append(nums[:i],append([]int{k},nums[i:]...)...)
        }
    }
    return append(nums,k)
}
func mySortDelete(nums []int, k int) []int{
    for i,v:=range nums{
        if v==k{
            // 插入k
            nums = append(nums[:i],nums[i+1:]...)
            break
        }
    }
    return nums
}

// 最大子数组和
func maxSubArray(nums []int) int {
    max := nums[0]
    for i:=1;i<len(nums);i++{
        if nums[i]+nums[i-1]>nums[i]{
            nums[i]+=nums[i-1]
        }
        if nums[i]>max{
            max = nums[i]
        }
    }
    return max
}
// 合并区间
func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] <  intervals[j][0]

	})
    ans := [][]int{intervals[0]}
    for i:=1;i<len(intervals);i++{
        if intervals[i][0] > ans[len(ans)-1][1]{
            ans = append(ans,intervals[i])
        } else{
            ans[len(ans)-1][1] = max(intervals[i][1],ans[len(ans)-1][1])
        }
    }

    return ans
}
func max(a,b int) int{
    if a>b{
        return a
    }
    return b
}
// 
func setZeroes(matrix [][]int)  {

    //记录包含0的行列
    r,c := make([]int,0),make([]int,0)
    for i:=0;i<len(matrix);i++{
        for j:=0;j<len(matrix[0]);j++{
            if matrix[i][j]==0{
                r = append(r,i)
                c = append(c,j)
            }

        }
    }

    for _, i := range r{
        for j:=0;j<len(matrix[0]);j++{
            matrix[i][j]=0
        }
    }
    for _, j := range c{
        for i:=0;i<len(matrix);i++{
            matrix[i][j]=0
        }
    }

}

// 螺旋矩阵
func spiralOrder(matrix [][]int) []int {

    ans := make([]int,0)
    if len(matrix)==0{
        return ans
    }
    // 上下左右的边界
    u,r,d,l := 0,len(matrix[0])-1,len(matrix)-1,0
    // 遍历
    for{
        for i:=l;i<=r;i++{
            ans = append(ans,matrix[u][i])
        }
        u++
        if u>d{
            break
        }
        for i:=u;i<=d;i++{
            ans = append(ans,matrix[i][r])
        }
        r--
        if r<l{
            break
        }
        for i:=r;i>=l;i--{
            ans = append(ans,matrix[d][i])
        }
        d--
        if d<u{
            break
        }
        for i:=d;i>=u;i--{
            ans = append(ans,matrix[i][l])
        }
        l++
        if l>r{
            break
        }
    }

    return ans
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // 相交链表
 func getIntersectionNode(headA, headB *ListNode) *ListNode {
    p,q := headA,headB
    // 双指针
    // 拼接 两个链表  两个指针 同时走  相遇时 即为相交点    
    // 
    for {
        if p==q{
            return p
        }
        p=p.Next
        q=q.Next
        if p==nil && q==nil{
            break
        }
        if p==nil{
            p = headB
        }
        if q==nil{
            q=headA
        }
  

    }
    return nil
}
// 反转链表
 func reverseList(head *ListNode) *ListNode {
    if head == nil{
        return nil
    }
    // 双指针
    var pre *ListNode
    cur:= head
    // 双指针 反转

    for cur != nil{
        next := cur.Next
        cur.Next = pre
        pre = cur
        cur = next

    }
    return pre

}
// 回文链表
func isPalindrome(head *ListNode) bool {
    //快慢指针找到中间节点
    slow, fast := head,head.Next
    for fast!=nil && fast.Next!=nil{
        slow=slow.Next
        fast=fast.Next.Next
    }

    // 反转后半部分
    tail := reverseList(slow.Next)
    // 比较
    for tail != nil{
        if head.Val != tail.Val{
            return false
        }
        head=head.Next
        tail=tail.Next
    }
    return true
    

}

func hasCycle(head *ListNode) bool {
    if head ==nil{
        return false
    }
    slow, fast := head,head.Next
    for fast!=nil && fast.Next!=nil{
        if slow==fast{
            return true
        }
        slow=slow.Next
        fast=fast.Next.Next
    }
    return false

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func inorderTraversal(root *TreeNode) []int {

    var ans []int
    var inorder func(root *TreeNode)
    inorder = func(root *TreeNode){
        if root == nil{
            return
        }
        inorder(root.Left)
        ans = append(ans,root.Val)
        inorder(root.Right)
        // 中序遍历 左 中 右
        // 前序遍历 中 左 右
        // 后序遍历 左 右 中
    }

    inorder(root)

    return ans

}
// 二叉树的最大深度
func maxDepth(root *TreeNode) int {
    if root == nil{
        return 0
    }
    return max(maxDepth(root.Left),maxDepth(root.Right))+1
}

// 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
    if root==nil{
        return nil
    }
    Right := invertTree(root.Left)
    Left := invertTree(root.Right)
    root.Right = Right
    root.Left = Left
    return root
}

func isSymmetric(root *TreeNode) bool {
    var check func(p,q *TreeNode) bool

    check = func(p,q *TreeNode) bool{
        if p==nil && q==nil{
            return true
        }
        if p==nil || q == nil {
            return false
        }

        return p.Val==q.Val && check(p.Left,q.Right) && check(p.Right,q.Left)
    }
    return check(root,root)
}

// 二叉树的直径

func diameterOfBinaryTree(root *TreeNode) int {
    ans := 0
    var maxDepth func(root *TreeNode) int
    maxDepth = func maxDepth(root *TreeNode) int {
        if root == nil{
            return 0
        }
        L := maxDepth(root.Left)
        R := maxDepth(root.Right)
        ans = max(ans,L+R+1)
        return max(L,R)+1
    }
    maxDepth(root)
    return ans-1
}


// 有序数组构建二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {

    var build func(left,right int) *TreeNode
    build = func(left,right int) *TreeNode{
        if left>right{
            return nil
        }
        mid := (left+right)/2
        root := &TreeNode{Val:nums[mid]}
        root.Left = build(left,mid-1)
        root.Right = build(mid+1,right)
        return root
    }

    return build(0,len(nums)-1)

    
}

// 全排列
func permute(nums []int) [][]int {
    l:=len(nums)
    res:=[][]int{}
    var backtrack func(int)
    backtrack=func(first int){
        if first==l{
            res=append(res,append([]int(nil),nums...))
            return
        }
        for i:=first;i<l;i++{
            nums[first],nums[i]=nums[i],nums[first]
            backtrack(first+1)
            nums[first],nums[i]=nums[i],nums[first]
        }
    }
    backtrack(0)
    return res
}

// 搜索插入位置
func searchInsert(nums []int, target int) int {
    l,r:=0,len(nums)-1
    mid := (l+r)/2
    for l<r{
        if nums[mid]==target{
            return mid
        }else if nums[mid]<target{
            l=mid+1
        }else {
            r=mid-1
        }
        mid = (l+r)/2
    }
    if nums[mid]<target{
        mid++
    }
    return mid
}

// 74. 搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
    r,c:=0,0
    for r<len(matrix){
        if target<matrix[r][0]{
            break
        }
        r++
    }
    r--
    if r>=0{
        for ;c<len(matrix[0]);c++{
            if target==matrix[r][c]{
                return true
            }
        }
    }

    return false
}

// 34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
    p,q := 0,len(nums)-1
    pflag,qflag:= false,false
    ans := []int{-1,-1}
    for p<=q{
        if nums[p]==target{
            if !pflag{
                ans[0]=p
                pflag=true
            }
        }
        if nums[q]==target{
            if !qflag{
                ans[1]=q
                qflag=true
            }
        }
        if !pflag{
            p++
        }
        if !qflag{
            q--
        }
        if pflag&&qflag{
            break
        }
    }
    return ans
}   

// 20. 有效的括号
func isValid(s string) bool {
    stack := []rune{}
    kuohao   := map[rune]rune{'{':'}','(':')','[':']'}


    for _,v:=range s{
        //入栈
        if _, ok:=kuohao[v];ok{
            stack = append(stack, v)
        }
        //出栈
        if v=='}' || v==']' || v==')'{
            // 栈不为空
            if len(stack)>0{
                left := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                if kuohao[left] != v{
                    return false
                }
            }else{
                return false
            }

        }
            
    }
    if len(stack)>0{
        return false
    }
    return true
}
//
// 347. 前 K 个高频元素
func topKFrequent(nums []int, k int) []int {
    maptop := make(map[int]int)

    for _, v := range nums{
        maptop[v]++
    }
    var listtop [][]int
    for k,v := range maptop{
        listtop = append(listtop,[]int{k,v})
    }

    sort.Slice(listtop, func(i,j int) bool {
        return listtop[i][1] > listtop[j][1]
    })
    var ans []int
    for i:=0;i<k;i++{
        ans = append(ans, listtop[i][0])
    }
    return ans
}
// 121. 买卖股票的最佳时机,d 动归
func maxProfit(prices []int) int {
    min := math.MaxInt
    maxprofit := 0
    for _,v:=range prices{
        if v < min{
            min = v
        }
        if v - min > maxprofit{
            maxprofit = v-min
        }

    }
    return maxprofit

}

// 55. 跳跃游戏
func canJump(nums []int) bool {
    maxPos := 0
    target := len(nums)
    for i,v := range nums{
        //若可达
        if maxPos >= i && i+v > maxPos{
            maxPos = i+v
        }
    }
    fmt.Println(maxPos)
    if maxPos>=target{
        return true

    }

    return false
}
// 45. 跳跃游戏 II
func jump(nums []int) int {
    //记录所有位置的可达
    maxpos := make([]int,len(nums))
    maxjump:=0
    for i,v:=range nums {
        if i+v > maxjump{
            maxjump = i+v
        }
        maxpos[i]=maxjump
    }

    target := len(nums)-1
    ans:=0
    fmt.Println(maxpos)
    //目前的位置 < target则
    for i:=0;i<target;{
        ans++
        i=maxpos[i]
    }
    return ans
}   


// 70. 爬楼梯
func climbStairs(n int) int {
    dp := make([]int,n+1)

    dp[0]=1
    dp[1]=1
    for i:=2;i<=n;i++{
        dp[i] = dp[i-1]+dp[i-2]
    }
    return dp[n]
}

// 118. 杨辉三角
func generate(numRows int) [][]int {
    
    dp:=make([][]int,numRows+1)
    var ans [][]int
    
    for i:=1;i<=numRows;i++{
        dp[i]=make([]int,i)
        dp[i][0]=1
        dp[i][i-1]=1
    }

    for i:=1;i<=numRows;i++{
        for j:=1;j<i-1;j++{
            dp[i][j] = dp[i-1][j]+dp[i-1][j-1]
        }
        ans = append(ans,dp[i])
    }
    return ans
}

// 198. 打家劫舍
func rob(nums []int) int {
    n:=len(nums)
    dp:=make([][]int,n+1)
    for i:=0;i<n+1;i++{
        //[no][yes]
        dp[i]=make([]int,2)
    }
    for i,v := range nums{
        //不偷就是上一轮取大值
        dp[i+1][0]+=max(dp[i][0],dp[i][1])
        //偷了上一轮就不能投
        dp[i+1][1]+=dp[i][0] + v
    }

    return max(dp[n][0],dp[n][1])
}
func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}