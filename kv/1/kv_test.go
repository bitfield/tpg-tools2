package kv_test

import (
	"os"
	"testing"

	"github.com/bitfield/kv"
	"github.com/google/go-cmp/cmp"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"kv": kv.Main,
	}))
}

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

func TestSetCreatesNewPairIfKeyDoesNotExist(t *testing.T) {
	t.Parallel()
	s := kv.NewStore()
	if _, ok := s.Get("test key"); ok {
		t.Fatal("new empty store shouldn't contain test key")
	}
	s.Set("test key", "test value")
	v, ok := s.Get("test key")
	if !ok {
		t.Fatal("test key not found after set")
	}
	if v != "test value" {
		t.Errorf("want value \"test value\", got %q", v)
	}
}

func TestSetUpdatesExistingKeyToNewValue(t *testing.T) {
	t.Parallel()
	s := kv.NewStore()
	s.Set("test key", "original value")
	s.Set("test key", "updated value")
	v, ok := s.Get("test key")
	if !ok {
		t.Fatal("test key not found after set")
	}
	if v != "updated value" {
		t.Errorf("want value \"updated value\", got %q", v)
	}
}

func TestGetReturnsValueAndOKIfKeyExists(t *testing.T) {
	t.Parallel()
	s := kv.NewStore()
	s.Set("test key", "test value")
	v, ok := s.Get("test key")
	if !ok {
		t.Fatal("not ok")
	}
	if v != "test value" {
		t.Errorf("want value \"test value\", got %q", v)
	}
}

func TestGetReturnsNotOKIfKeyDoesNotExist(t *testing.T) {
	t.Parallel()
	s := kv.NewStore()
	_, ok := s.Get("test key")
	if ok {
		t.Fatal("unexpected ok")
	}
}

func TestAllReturnsExpectedMap(t *testing.T) {
	t.Parallel()
	s := kv.NewStore()
	s.Set("A", "1")
	s.Set("B", "2")
	s.Set("C", "3")
	want := map[string]string{
		"A": "1",
		"B": "2",
		"C": "3",
	}
	got := s.All()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSaveCreatesExpectedFile(t *testing.T) {
	t.Parallel()
	s := kv.NewStore()
	s.Set("A", "1")
	s.Set("B", "2")
	s.Set("C", "3")
	path := t.TempDir() + "/kvtest.store"
	err := s.Save(path)
	if err != nil {
		t.Fatal(err)
	}
	want, err := os.ReadFile("testdata/golden.store")
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSaveReturnsErrorWhenFileCannotBeCreated(t *testing.T) {
	t.Parallel()
	s := kv.NewStore()
	err := s.Save("boguspath/doesntexist")
	if err == nil {
		t.Fatal("no error")
	}
}

func TestOpenStore_ReadsCorrectDataFromGivenStoreFile(t *testing.T) {
	t.Parallel()
	s, err := kv.OpenStore("testdata/golden.store")
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{
		"A": "1",
		"B": "2",
		"C": "3",
	}
	got := s.All()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestOpenStore_ReturnsErrorOnInvalidJSON(t *testing.T) {
	t.Parallel()
	_, err := kv.OpenStore("testdata/invalid.store")
	if err == nil {
		t.Fatal("no error")
	}
}
