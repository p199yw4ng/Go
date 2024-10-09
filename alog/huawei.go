// 21. 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    head,cur := &ListNode{}, &ListNode{}
    head = cur
    for list1 != nil && list2 !=nil{
        if list1.Val < list2.Val{
            cur.Next = list1
            list1 = list1.Next
        }else{
            cur.Next = list2
            list2 = list2.Next
        }
        cur = cur.Next
    }

    if list1==nil{
        cur.Next = list2
    }
    if list2==nil{
        cur.Next = list1
    }
    return head.Next
}

//27. 移除元素
func removeElement(nums []int, val int) int {
    
    j:=0
    for i,v :=range nums{
        if v!=val{
            nums[j]=nums[i]
            j++
        }
    }
    ans := j
    for j<len(nums){
        nums[j]=val
        j++
    }
   

    return ans
}


//14. 最长公共前缀
func longestCommonPrefix(strs []string) string {

    i:=0
    for {
        if Prefix(strs,i){
            i++
        }else{
            break
        }
    }
    return strs[0][:i]
}

func Prefix(strs []string, i int) bool {
    if i>=len(strs[0]){
        return false
    }
    tamp:=strs[0][i]
    for _, str:=range strs{
        
        if i < len(str) {
            if tamp != str[i]{
                return false
            }
            
        }else{
            return false
        }
    }
    return true
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

//	83. 删除排序链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
    mapNode := make(map[int]bool)

    if head==nil{
        return nil
    }

    pre := head
    mapNode[head.Val]=true
    cur := pre.Next

    for cur!=nil{
        if mapNode[cur.Val]{
            pre.Next = pre.Next.Next
        }else{
            mapNode[cur.Val]=true
            pre = pre.Next
        }
        cur = cur.Next
    
    }
    return head
}

//217. 存在重复元素
func containsDuplicate(nums []int) bool {
    mapNums := make(map[int]bool)
    for _, v:=range nums{
        if mapNums[v]{
            return true
        }else{
            mapNums[v] = true
        }
    }
    return false
}   
//268. 丢失的数字
func missingNumber(nums []int) int {
    sort.Ints(nums)
    for i,v:=range nums{
        if i!=v{
            return i
        }
    }

    return len(nums)
}


//415. 字符串相加

func addStrings(num1 string, num2 string) string {
    l1, l2 := len(num1), len(num2)
    var flag int
    var ans []byte
    for i,j := l1-1,l2-1;i>=0||j>=0;i,j=i-1,j-1{
        var x,y int
        if i>=0{
            x = int(num1[i]-'0')
        }
        if j>=0{
            y = int(num2[j]-'0')
        }
        tamp := x+y
        if flag == 1{
            tamp++
        }
        flag = tamp / 10
        tamp %= 10
        ans = append(ans,byte(tamp+'0'))
    }
    if flag==1{
        ans = append(ans, byte(1+'0'))
    }
    fmt.Println(string(ans))
    var strAns []byte
    for i:=len(ans)-1;i>=0;i--{
        strAns = append(strAns, ans[i])
    }

    return string(strAns)
}

//605. 种花问题
func canPlaceFlowers(flowerbed []int, n int) bool {

    flowerbed = append([]int{0},flowerbed...)
    flowerbed = append(flowerbed,0)
    pre, cur, next := 0,0,0

    for i:=1;i<len(flowerbed)-1;i++{
        
        pre = flowerbed[i-1]
        cur = flowerbed[i]
        next = flowerbed[i+1]
        
        if pre==0 && cur==0 && next==0{
            n--
            flowerbed[i] = 1
        }
        
    }
    fmt.Println(n,flowerbed)
    return n<=0
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 //637. 二叉树的层平均值
 func averageOfLevels(root *TreeNode) []float64 {
    var ans []float64
    stack :=[]*TreeNode{root}
    for len(stack)>0 {
        sum :=0.0
        var tamp []*TreeNode
        for _, v:=range stack{
            sum+=float64(v.Val)
            if v.Left != nil{
                tamp = append(tamp,v.Left)
            }
            if v.Right != nil{
                tamp = append(tamp,v.Right)
            }
        }
        ans = append(ans,sum/float64(len(stack)))
        stack = tamp
    }
    return ans
}

// 746. 使用最小花费爬楼梯
func minCostClimbingStairs(cost []int) int {
    n:=len(cost)
    dp := make([]int,n+1)
    cost = append(cost,0)
    dp[0] = cost[0]
    dp[1] = cost[1]
    for i:=2;i<=n;i++{
        dp[i] = min(dp[i-1],dp[i-2])+cost[i]
    }
    return dp[n]
}


//1047. 删除字符串中的所有相邻重复项
func removeDuplicates(s string) string {
    flag := true
    sByte := []byte(s)
    for flag{
        flag = false
        for i:=0;i<len(sByte)-1;i++{
            if sByte[i]==sByte[i+1]{
                sByte = append(sByte[:i],sByte[i+2:]...)
                flag = true
                break
            }
        }
    }
    return string(sByte)
}


// 1614. 括号的最大嵌套深度

func maxDepth(s string) int {
    var max,ans int
    
    for _, v:= range s{
        if v=='('{
            max++
            if max>ans{
                ans = max
            }
        }
        if v==')'{
            max--
        }
    }
    return ans
}



// 面试题 04.02. 最小高度树:给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。



func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums)==0{
        return nil
    }
    mid := len(nums)/2
    root := &TreeNode{Val:nums[mid]}
    root.Left=sortedArrayToBST(nums[:mid])
    root.Right=sortedArrayToBST(nums[mid+1:])
    return root
}