package raindrops

import "strconv"

var rules = []struct {
    divisor int
    sound string
}{
    {3, "Pling"},
    {5, "Plang"},
    {7, "Plong"},
}

func Convert(number int) string {
    var result string
	// if number % 3 == 0 {
	//     result += "Pling"
	// }
 //    if number % 5 == 0 {
 //        result += "Plang"
 //    } 
 //    if number % 7 == 0{
 //        result += "Plong"
 //    }

    for _, r := range rules {
        if number%r.divisor == 0 {
            result += r.sound
        }
    }
    if result == "" {
        result = strconv.Itoa(number)
    }
    return result
    
}
