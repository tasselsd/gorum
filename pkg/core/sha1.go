package core

import (
	"crypto/sha1"
	"fmt"
	"io"
)

type Sha1Object struct {
	sha1Sum string
}

func NewSha1Object(value string) *Sha1Object {
	sha1Obj := Sha1Object{}
	h := sha1.New()
	io.WriteString(h, value)
	sha1Obj.sha1Sum = fmt.Sprintf("%x", h.Sum(nil))
	return &sha1Obj
}

func (obj *Sha1Object) Sha1() string {
	return obj.sha1Sum
}

func (obj *Sha1Object) ShortSha1() string {
	return obj.sha1Sum[:8]
}
