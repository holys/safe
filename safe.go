// Copyright 2015 David Chen <chendahui007@gmail.com>.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package safe implements a password safety checker.
package safe

import (
	"regexp"
	"strings"
)

type Level uint8

const (
	Terrible Level = iota
	Simple
	Medium
	Strong
)

// qwertyuiop
//  asdfghjkl
//   zxcvbnm
var (
	asdf    = "qwertyuiopasdfghjklzxcvbnm"
	revAsdf = reverse(asdf)
)

var (
	lower  = regexp.MustCompile(`[a-z]`)
	upper  = regexp.MustCompile(`[A-Z]`)
	number = regexp.MustCompile(`[0-9]`)
	marks  = regexp.MustCompile(`[^0-9a-zA-Z]`)
)

func Check(s string) Level {
	return safe.Check(s)
}

var safe = new()

type safety struct{}

func new() *safety {
	return &safety{}
}

func (s *safety) Check(raw string) Level {

	// TODO
	return Terrible
}

// If the password is in the order on keyboard.
func (s *safety) isAsdf(raw string) bool {
	// s in asdf , or reverse in asdf
	rev := reverse(raw)
	if strings.Contains(asdf, raw) || strings.Contains(asdf, rev) {
		return true
	}

	// s in reverse(asdf),  or reverse in reverse(asdf)
	if strings.Contains(revAsdf, raw) || strings.Contains(revAsdf, rev) {
		return true
	}

	return false
}

// If the password is alphabet step by step.
func (s *safety) isByStep(raw string) bool {
	r := []rune(raw)
	delta := r[1] - r[0]

	for i, _ := range r {
		if i == 0 {
			continue
		}
		if r[i]-r[i-1] != delta {
			return false
		}
	}

	return true
}

// If the password is common used
// 10k top passwords: https://xato.net/passwords/more-top-worst-passwords/
func (s *safety) isCommonPassword(raw string) bool {
	return false
}

//TODO: 优化查找,  对比不同的查找算法

func loadWords() {

}

func reverse(raw string) string {
	r := []rune(raw)

	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
