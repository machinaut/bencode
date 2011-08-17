// Alex Ray 2011 <ajray@ncsu.edu>

// Reference:
// http://en.wikipedia.org/wiki/Bencode

package bencode

import "testing"

type intTest struct {
    i int
    b []byte
}
var intTests = []intTest {
    intTest{42, []byte("i42e") },
    intTest{0, []byte("i0e") },
    intTest{-42, []byte("i-42e") },
}

type byteTest struct {
    a, b []byte
}
var byteTests = []byteTest {
    byteTest{[]byte("spam"), []byte("4:spam")},
    byteTest{[]byte(""), []byte("0:")},
    byteTest{[]byte{}, []byte("0:")},
    byteTest{[]byte{0,9}, []byte{'2',':',0,9}},
}

type listTest struct {
    l [][]byte
    b []byte
}
var listTests = []listTest{
    listTest{[][]byte{[]byte("4:spam"),[]byte("i42e")},
        []byte("l4:spami42ee")},
    listTest{[][]byte{[]byte("3:foo")},
        []byte("l3:fooe")},
    listTest{[][]byte{},
        []byte("le")},
}

// Test if two []byte's are equal
func equal(a, b []byte) bool {
    if len(a) != len(b) {
        return false
    }
    for i :=0 ; i < len(a) ; i++ {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func TestEncInc(t *testing.T) {
    for _, it := range intTests {
        b := EncInt(it.i)
        if !equal(b, it.b) {
            t.Errorf("EncInt(%d) = %s, want %s.",it.i,b,it.b)
        }
    }
}

func TestEncBytes(t *testing.T) {
    for _, it := range byteTests {
        b := EncBytes(it.a)
        if !equal(b, it.b) {
            t.Errorf("EncBytes(%q) = %q, want %q.",it.a,b,it.b)
        }
    }
}

func TestEncList(t *testing.T) {
    for _, it := range listTests {
        b := EncList(it.l)
        if !equal(b, it.b) {
            t.Errorf("EncList(%v) = %q, want %q.",it.l,b,it.b)
        }
    }
}

