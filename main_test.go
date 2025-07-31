package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReplaceBadWords(t *testing.T) {
	// Test Case 1: No bad words present
	chirp1 := &Chirp{Body: "This is a normal chirp."}
	cleanBody := replaceBadWords(chirp1.Body)
	require.Equal(t, "This is a normal chirp.", cleanBody, "Chirp with no bad words should remain unchanged")

	// Test Case 2: One bad word (lowercase)
	chirp2 := &Chirp{Body: "This is a kerfuffle opinion."}
	cleanBody = replaceBadWords(chirp2.Body)
	require.Equal(t, "This is a **** opinion.", cleanBody, "Single lowercase bad word should be replaced")

	// Test Case 3: One bad word (uppercase)
	chirp3 := &Chirp{Body: "I saw a SHARBERT today."}
	cleanBody = replaceBadWords(chirp3.Body)
	require.Equal(t, "I saw a **** today.", cleanBody, "Single uppercase bad word should be replaced")

	// Test Case 4: Multiple different bad words
	chirp4 := &Chirp{Body: "Fornax and kerfuffle are bad."}
	cleanBody = replaceBadWords(chirp4.Body)
	require.Equal(t, "**** and **** are bad.", cleanBody, "Multiple different bad words should be replaced")

	// Test Case 5: Bad word with punctuation (should NOT be replaced)
	chirp5 := &Chirp{Body: "That's a sharbert! What a mess."}
	cleanBody = replaceBadWords(chirp5.Body)
	require.Equal(t, "That's a sharbert! What a mess.", cleanBody, "Bad word with punctuation should not be replaced")

	// Test Case 6: Empty chirp
	chirp6 := &Chirp{Body: ""}
	cleanBody = replaceBadWords(chirp6.Body)
	require.Equal(t, "", cleanBody, "Empty chirp should remain empty")
}
