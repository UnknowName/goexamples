package point

func FindPeaks(mountain []int) []int {
    if len(mountain) < 3 {
        return nil
    }
    result := make([]int, 0)
    pre := 0
    mid := 1
    last := 2
    for last < len(mountain) {
        if mountain[mid] > mountain[pre] && mountain[mid] > mountain[last] {
            result = append(result, mid)
        }
        last++
        pre++
        mid++
    }
    return result
}
