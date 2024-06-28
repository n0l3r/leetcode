func replaceDigits(s string) string {
    res := []byte(s)

    for i := 1; i < len(s); i += 2 {
        res[i] = res[i-1] + res[i] - '0'
    }

    return string(res)
}