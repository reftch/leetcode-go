// Package fullJustify contains a stub for LeetCode problem 0068. fullJustify.
//
// Link: https://leetcode.com/problems/fullJustify/
//
// Given an array of strings words and a width maxWidth,
// format the text such that each line has exactly maxWidth characters and is fully (left and right) justified.
//
// You should pack your words in a greedy approach; that is,
// pack as many words as you can in each line.
// Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.
//
// Extra spaces between words should be distributed as evenly as possible.
// If the number of spaces on a line does not divide evenly between words, the empty slots on the left will be assigned more spaces than the slots on the right.
//
// For the last line of text, it should be left-justified, and no extra space is inserted between words.
//
// Note:
//
// A word is defined as a character sequence consisting of non-space characters only.
// Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
// The input array words contains at least one word.
//
// Example 1:
//
// Input: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
// Output:
// [
//    "This    is    an",
//    "example  of text",
//    "justification.  "
// ]
//
// Example 2:
//
// Input: words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
// Output:
// [
//   "What   must   be",
//   "acknowledgment  ",
//   "shall be        "
// ]
// Explanation: Note that the last line is "shall be    " instead of "shall     be", because the last line must be left-justified instead of fully-justified.
// Note that the second line is also left-justified because it contains only one word.
//
// Example 3:
//
// Input: words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"], maxWidth = 20
// Output:
// [
// "Science  is  what we",
// "understand      well",
// "enough to explain to",
// "a  computer.  Art is",
// "everything  else  we",
// "do                  "
// ]

package fullJustify

import (
	"fmt"
	"math"
)

func addLine(words_acc []string, maxWidth int, w int) string {
	accumulate := ""
	number_of_spaces := len(words_acc)
	if number_of_spaces > 1 {
		spaces_w := int(math.Round(float64(maxWidth-w) / float64(number_of_spaces-1)))
		// fmt.Printf("Spaces for %d %d %d\n", spaces_w, number_of_spaces, maxWidth-w)

		for i := range words_acc {
			accumulate += words_acc[i]
			if i < len(words_acc)-1 {
				for range spaces_w {
					if w < maxWidth {
						accumulate += " "
						w++
					}
				}
			}
		}
	} else {
		accumulate = words_acc[0]
		for range maxWidth - len(words_acc[0]) {
			accumulate += " "
		}
	}

	return accumulate
}

func addLastLine(words_acc []string, maxWidth int) string {
	accumulate := ""
	for _, word := range words_acc {
		accumulate += word + " "
	}
	if len(accumulate) < maxWidth {
		for range maxWidth - len(accumulate) {
			accumulate += " "
		}
	}
	return accumulate
}

func fullJustify(words []string, maxWidth int) []string {
	new := []string{}

	words_acc := []string{}
	w := 0
	for _, word := range words {
		len_w := len(word)
		if w+len_w+len(words_acc) <= maxWidth {
			words_acc = append(words_acc, word)
			w += len_w
		} else {
			accumulate := addLine(words_acc, maxWidth, w)
			new = append(new, accumulate)
			fmt.Println(accumulate)

			words_acc = []string{}
			words_acc = append(words_acc, word)
			w = len_w
		}
	}

	if w != 0 {
		new = append(new, addLastLine(words_acc, maxWidth))
	}

	fmt.Printf("Result %d\n", len(new))
	for _, line := range new {
		fmt.Printf("[%s]\n", line)
	}

	return new
}
