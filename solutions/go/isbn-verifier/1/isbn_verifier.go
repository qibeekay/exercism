package isbnverifier

func IsValidISBN(isbn string) bool {
    sum := 0
    count := 0          // how many real characters we have seen

    for i := 0; i < len(isbn); i++ {
        ch := isbn[i]

        if ch == '-' {
            continue
        }

        var value int
        if ch >= '0' && ch <= '9' {
            value = int(ch - '0')
        } else if ch == 'X' && count == 9 {
            value = 10          // X is only allowed as the 10th character
        } else {
            return false
        }

        sum += value * (10 - count)   // weight: 10, 9, 8, ... 1
        count++
    }

    return count == 10 && sum%11 == 0
}