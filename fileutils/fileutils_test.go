package hashing

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var tExe io.Reader
var tExeSum = []byte("a0f1984f7fbd6a41d10b4f134caf24a6c4bffb105ad1aa79ca7d5a61ce58e10d0b792a0829b72bce74a1259950eaddf37f688e46a1213f51d7b89cc58fd93a23")
var tExe2 io.Reader
var tExe2Sum = []byte("a14dfb5cc18d917b7916632fa08f7179f8f3375afd979f6b6bfe1ec074052fd3dec44983c0ff5bdb646eaecc4a57c9c5ea8162ffed37893c7a08ebb2967dc762")
var tExeDiff io.Reader

func TestMain(m *testing.M) {
	var err error

	tExe, err = os.Open("../fixtures/test_files/lyyti.exe")
	if err != nil {
		log.Fatalln("couldn't load test binary")
	}

	tExe2, err = os.Open("../fixtures/test_files/lyyti.2.exe")
	if err != nil {
		log.Fatalln("couldn't load test binary 2")
	}

	d, err := ioutil.ReadFile("../fixtures/test_files/lyyti.exe.bsdiff")
	tExeDiff = bytes.NewBuffer(d)
	if err != nil {
		log.Fatalln("couldn't load bsdiff")
	}

	os.Exit(m.Run())
}

func TestCalcHash(t *testing.T) {
	o, err := CalcHash(tExe)
	if err != nil {
		t.Error(err)
		return
	}

	if !bytes.Equal(o, tExeSum) {
		t.Errorf("mismatch\n>> %s\n>> %s", o, tExeSum)
	}
}

func TestCalcDiff(t *testing.T) {
	d, err := CalcDiff(tExe, tExe2)
	if err != nil {
		t.Error(err)
		return
	}

	var ourDiff []byte
	io.ReadFull(d, ourDiff)

	var knownDiff []byte
	io.ReadFull(tExe2, knownDiff)

	if !bytes.Equal(ourDiff, knownDiff) {
		t.Error("bsdiff output differed")
	}
}

func TestApplyPatch(t *testing.T) {
	patchedExeBuf, err := ApplyPatch(tExe, tExeDiff)
	if err != nil {
		t.Error(err)
		return
	}

	var patch []byte
	io.ReadFull(patchedExeBuf, patch)

	var exe2 []byte
	io.ReadFull(tExe2, exe2)

	if !bytes.Equal(patch, exe2) {
		t.Error("lyyti.2.exe output differed")
		h1, _ := CalcHash(tExe2)
		h2, _ := CalcHash(patchedExeBuf)
		log.Printf("HASH DIFFERENCE:\noriginal >> %s\npatched  >> %s\n", h1, h2)
	}
}
