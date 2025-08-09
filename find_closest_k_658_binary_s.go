func findClosestElements(arr []int, k int, x int) []int {

    left := 0 
    right := len(arr)-k

    for left < right{

        mid := left+(right-left)/2
        if x - arr[mid] > arr[mid+k] - x{
            left = mid + 1
        }else{
            right = mid
        }

    }

    return arr[left:left+k]
    
}

// Special kind of binary search
// here we give left priority and have strict bounds for right according to k
// imagine array [1, 2, 3, 4, 5, 6, 7] and a x = 4, k = 5  [ 2, 3, 4, 5] , since the left or lower number is priority in case of the same difference we have 2 here.
// So it makes sense to have left bounds at 0 and right, as confining it within k limits. The other solution which involves sliding window + binary search is more intuitive though.
// But this is unique way to approach this problem. 
