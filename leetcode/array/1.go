package array

func Filter(n int) int {
    total := 0
    checked := make([]bool, n)
    for i := 2; i < n; i++ {
        if !checked[i] {
            for j := i * i; j < n; j += i {
                checked[j] = true
            }
            total++
        }
    }
    return total
}
