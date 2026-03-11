package main

import (
	"go-reloaded/processor"
	//"reflect"
	"testing"
)

func equalSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range b {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestFixArticles(t *testing.T) {
	want := []string{"an", "apple", "bits", "makes", "man"}
	got := processor.FixArticles([]string{"a", "apple", "bits", "makes", "man"})

	if !equalSlice(got, want) {
		t.Errorf("expected: %v got: %v", want, got)
	}
}

func TestTokenize(t *testing.T) {
	want := []string{"hello", "...", "'", "world", "'"}
	got := processor.Tokenize("hello ... ' world '")

	if !equalSlice(got, want) {
		t.Errorf("got: %v want: %v", got, want)
	}
}

func TestFixApplyModifiers(t *testing.T) {
	input := []string{"FF", "(hex)", "this", "man", "(up, 2)"}
	want := []string{"255", "THIS", "MAN"}
	got := processor.ApplyModifiers(input)
	if !equalSlice(want, got) {
		t.Errorf("expected: %v  got: %v", want, got)
	}
}

func TestRebuild(t *testing.T) {
	input := []string{"that", "man", "...", "'", "is", "a", "kind", "'"}
	want := "that man... 'is a kind'"
	got := processor.Rebuild(input)
	if got != want {
		t.Errorf("Rebuild(%v) got: %v want: %v ", input, got, want)
	}
}
