package hashing

import (
	"bytes"
	"io"
	"log"
	"os"
	"testing"
)

var tExe io.Reader
var tExeSum = []byte("f0d1afb614741b8f532bb25b36076691872767f7495a837b447a985e1825d68eeffa47165db6d7274078aef4c5773cfbce2d7aaa25115d9c87350a0aabba65f0")
var tExe2 io.Reader
var tExe2Sum = []byte("021720fc32b214e6e3c52457f965ce125e47250a6c82a8a750e08e3f53468bd0190d437e10068464aa393e03f556f71c23f430e69f6d0c2b00e2e8bbd88e11bf")
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

	tExeDiff, err = os.Open("../fixtures/test_files/lyyti.exe.bsdiff")
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
