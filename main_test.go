package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReplaceBadWords(t *testing.T) {
	// Test Case 1: No bad words present
	chirp1 := &Chirp{Content: "This is a normal chirp."}
	chirp1.replaceBadWords()
	require.Equal(t, "This is a normal chirp.", chirp1.Content, "Chirp with no bad words should remain unchanged")

	// Test Case 2: One bad word (lowercase)
	chirp2 := &Chirp{Content: "This is a kerfuffle opinion."}
	chirp2.replaceBadWords()
	require.Equal(t, "This is a **** opinion.", chirp2.Content, "Single lowercase bad word should be replaced")

	// Test Case 3: One bad word (uppercase)
	chirp3 := &Chirp{Content: "I saw a SHARBERT today."}
	chirp3.replaceBadWords()
	require.Equal(t, "I saw a **** today.", chirp3.Content, "Single uppercase bad word should be replaced")

	// Test Case 4: Multiple different bad words
	chirp4 := &Chirp{Content: "Fornax and kerfuffle are bad."}
	chirp4.replaceBadWords()
	require.Equal(t, "**** and **** are bad.", chirp4.Content, "Multiple different bad words should be replaced")

	// Test Case 5: Bad word with punctuation (should NOT be replaced)
	chirp5 := &Chirp{Content: "That's a sharbert! What a mess."}
	chirp5.replaceBadWords()
	require.Equal(t, "That's a sharbert! What a mess.", chirp5.Content, "Bad word with punctuation should not be replaced")

	// Test Case 6: Empty chirp
	chirp6 := &Chirp{Content: ""}
	chirp6.replaceBadWords()
	require.Equal(t, "", chirp6.Content, "Empty chirp should remain empty")
}
