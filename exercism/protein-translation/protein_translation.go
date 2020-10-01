package protein

import (
	"errors"
)
var ErrInvalidBase = errors.New("error invalid base")
var ErrStop = errors.New("error stop")

func FromCodon(input string) (string, error){
	switch input {
	case "AUG":
		return "Methionine", nil
	case "UUU" , "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

func FromRNA(input string) ([]string, error){
	i:=0
	var rna []string
	for i + 3 <= len(input){
		substr := input[i:i+3]
		rtn, err := FromCodon(substr)
		if err == ErrInvalidBase{
			return rna, ErrInvalidBase
		}
		if err == ErrStop{
			return rna, nil
		}
		rna = append(rna, rtn)
		i += 3
	}
	return rna, nil
}