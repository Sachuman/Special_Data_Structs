func partition(s string) [][]string {

    ret := make([][]string, 0)
    dfs(s, make([]string, 0), &ret)
    return ret
}

func dfs(s string, path[]string, ret *[][]string){

    if len(s) == 0{
        lol := make([]string, 0)
        lol = append(lol, path...)
        *ret = append(*ret, lol)
    }
    for i:=1; i <= len(s); i++{
        if palindrome(s[:i]){
            dfs(s[i:], append(path, s[:i]), ret)
        }
    }

}

func palindrome(s string) bool{
    start := 0
    end := len(s)-1

    for start < end{
        if s[start] != s[end]{
            return false
        }
        start ++
        end --
    }
    return true

}

// 131 // good backtracking problem little unique tree imo
