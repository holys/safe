// Copyright 2015 David Chen <chendahui007@gmail.com>.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package safe

import (
	"testing"
)

var s = New(8, 0, 3, Strong)

func TestCheck(t *testing.T) {

}

func TestIsAsdf(t *testing.T) {
	for _, c := range []struct {
		in   string
		want bool
	}{
		{"qwer", true},
		{"tyuio", true},
		{"asdf", true},
		{"lkjhg", true},
		{"zxcvb", true},
		{"mnbvc", true},
		{"Asdf", false},
		{"qwrty", false},
	} {
		got := s.isAsdf(c.in)
		if got != c.want {
			t.Errorf("got %t want %t", got, c.want)
		}
	}
}

func TestIsByStep(t *testing.T) {
	for _, c := range []struct {
		in   string
		want bool
	}{
		{"abc", true},
		{"hijklmn", true},
		{"aceg", true},
		{"asdf", false},
		{"123456", true},
		{"13579", true},
		{"123567", false},
	} {
		got := s.isByStep(c.in)
		if got != c.want {
			t.Errorf("got %t want %t", got, c.want)
		}
	}

}

func TestIsCommonPassword(t *testing.T) {

}

func BenchmarkIsAsdf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.isAsdf("asdfghjkl")
	}
}

func BenchmarkIsByStep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.isByStep("abcdefg")
	}
}

func BenchmarkIsCommonPassword(b *testing.B) {

}
