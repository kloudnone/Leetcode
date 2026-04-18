func mirrorDistance(n int) int {
    ans := n - reverse(n)
    return max(ans, -ans)
}

func reverse(x int) int {
    y := 0
    for x > 0 {
        y = y * 10 + x%10
        x /= 10
    }
    return y
}