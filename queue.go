package main

import (
	"errors"
	"strings"
)

type Gender string

var (
	Boy  Gender = "B"
	Girl Gender = "G"
)

type Queue []Gender

func (c Queue) arrangeOne() {
	for i := 0; i <= len(c)-2; i++ {
		if c[i] == Boy && c[i+1] == Girl {
			c[i], c[i+1] = c[i+1], c[i]
			i += 2
			continue
		}
	}
}

func (c Queue) ArrangeUntil(i int) {
	if i < 1 {
		return
	}
	for index := 1; index <= i; index++ {
		c.arrangeOne()
	}
}

func NewQueue(s string) (Queue, error) {
	q := make([]Gender, 0)
	for _, c := range strings.ToUpper(s) {
		if Gender(c) != Boy && Gender(c) != Girl {
			return nil, errors.New("sequence contains invalid gender")
		}
		q = append(q, Gender(c))
	}
	return q, nil
}
