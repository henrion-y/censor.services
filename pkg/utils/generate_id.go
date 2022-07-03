package utils

import (
	"regexp"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

const (
	Str_Base_Alphabet    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NumBer_Base_Alphabet = "1234567890"
	All_Base_Alphabet    = Str_Base_Alphabet + NumBer_Base_Alphabet
)

func GenerateEntityID(entityID string) (string, error) {
	if len(entityID) == 16 {
		if regexp.MustCompile(`^[a-z0-9A-Z]+$`).MatchString(entityID) {
			return entityID, nil
		}
	}
	return gonanoid.Generate(All_Base_Alphabet, 16)
}
