package dioxyribonucleic

import "strings"

func DNAStrand(dna string) string {
	// your code here
	comp := make([]rune, 0)
	for _, v := range dna {
		switch v {
		case 'A':
			comp = append(comp, 'T')

		case 'T':
			comp = append(comp, 'A')
		case 'G':
			comp = append(comp, 'C')
		case 'C':
			comp = append(comp, 'G')
		default:
			return ""
		}
	}
	return string(comp)
}

var dnaReplacer *strings.Replacer = strings.NewReplacer(
	"A", "T",
	"T", "A",
	"C", "G",
	"G", "C",
)

// DNAStrandBest - this best practice Codewars
func DNAStrandBest(dna string) string {
	return dnaReplacer.Replace(dna)
}
