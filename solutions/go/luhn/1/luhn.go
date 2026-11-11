package luhn

func Valid(id string) bool {
    sum := 0
    digits := 0
    double := false 
	for i := len(id) - 1; i >= 0; i-- {
        ch := id[i]
        if ch == ' ' {
            continue
        }
        if ch < '0' || ch > '9' {
            return false
        }

        n := int (ch - '0')

        if double {
            n *= 2
            if n > 9 {
                n -= 9
            }
        } 

        sum += n
		digits++
		double = !double
    }
    return digits >= 2 && sum%10 == 0
}