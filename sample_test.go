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
	for i, _ := range b {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestFixArticles(t *testing.T) {}

func TestTokenize(t *testing.T) {
	want := []string{"hello", "...", "'", "world", "'"}
	got := processor.Tokenize("hello ... ' world '")

	if !equalSlice(got, want) {
		t.Errorf("got: %v want: %v", got, want)
	}
}

func TestFixApplyModifiers(t *testing.T) {}

func TestRebuild(t *testing.T) {}
