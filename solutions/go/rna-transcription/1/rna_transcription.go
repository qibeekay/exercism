package rnatranscription

func ToRNA(dna string) string {
	dnaRune := make([]rune, 0, len(dna))

    for _, char := range dna {
        switch char {
        case 'G':
            dnaRune = append(dnaRune, 'C')
        case 'C':
            dnaRune = append(dnaRune, 'G')
        case 'T':
            dnaRune = append(dnaRune, 'A')
        case 'A':
            dnaRune = append(dnaRune, 'U')
        default:
            dnaRune = append(dnaRune, char)
        }
    }
    return string(dnaRune)
}
