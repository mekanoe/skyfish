package hashing

import (
	"bytes"
	"crypto/sha512"
	"encoding/hex"
	"io"

	"github.com/kr/binarydist"
)

// CalcHash spits out a hex-encoded SHA-512 sum of it's io.Reader
func CalcHash(r io.Reader) (b []byte, err error) {
	hash := sha512.New()
	_, err = io.Copy(hash, r)
	if err != nil {
		return
	}

	sum := hash.Sum(b)
	b = make([]byte, hex.EncodedLen(len(sum)))
	hex.Encode(b, sum)

	return
}

// CalcDiff calculates the bsdiff of each reader, and spits out a reader with the diff.
func CalcDiff(cur, next io.Reader) (io.Reader, error) {
	buf := &bytes.Buffer{}
	binarydist.Diff(cur, next, buf)

	return buf, nil
}

// ApplyPatch to the current file, and will return an io.Reader with the patched binary.
func ApplyPatch(cur, patch io.Reader) (io.Reader, error) {
	buf := &bytes.Buffer{}
	err := binarydist.Patch(cur, buf, patch)

	return buf, err
}
