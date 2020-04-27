// Package hamming provides a method to calculate the distance between 2 DNS strands
package hamming

import (
	"errors"
)

// Distance returns the number of different nucleotides between 2 homologous DNS strands
func Distance(a, b string) (int, error) {
	ra, rb := []rune(a), []rune(b)
	if len(a) != len(b) {
		return 0, errors.New("DNA strands are different lengths")
	}

	d := 0
	for i := 0; i < len(a); i++ {
		if ra[i] != rb[i] {
			d++
		}

	}
	return d, nil
}
