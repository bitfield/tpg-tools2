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

func TestSetUpdatesExistingKeyToNewValue(t *testing.T) {
	t.Parallel()
	s, err := kv.OpenStore("dummy path")
	if err != nil {
		t.Fatal(err)
	}
	s.Set("key", "original")
	s.Set("key", "updated")
	v, ok := s.Get("key")
	if !ok {
		t.Fatal("key not found")
	}
	if v != "updated" {
		t.Errorf("want 'updated', got %q", v)
	}
}

func TestGetReturnsValueAndOKIfKeyExists(t *testing.T) {
	t.Parallel()
	s, err := kv.OpenStore("dummy path")
	if err != nil {
		t.Fatal(err)
	}
	s.Set("key", "value")
	v, ok := s.Get("key")
	if !ok {
		t.Fatal("not ok")
	}
	if v != "value" {
		t.Errorf("want 'value', got %q", v)
	}
}

func TestGetReturnsNotOKIfKeyDoesNotExist(t *testing.T) {
	t.Parallel()
	s, err := kv.OpenStore("dummy path")
	if err != nil {
		t.Fatal(err)
	}
	_, ok := s.Get("key")
	if ok {
		t.Fatal("unexpected ok")
	}
}

func TestSave_ErrorsWhenPathUnwritable(t *testing.T) {
	t.Parallel()
	s, err := kv.OpenStore("bogus/unwritable.store")
	if err != nil {
		t.Fatal("no error")
	}
	err = s.Save()
	if err == nil {
		t.Fatal("no error")
	}
}

func TestSaveSavesDataPersistently(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/kvtest.store"
	s, err := kv.OpenStore(path)
	if err != nil {
		t.Fatal(err)
	}
	s.Set("A", "1")
	s.Set("B", "2")
	s.Set("C", "3")
	err = s.Save()
	if err != nil {
		t.Fatal(err)
	}
	s2, err := kv.OpenStore(path)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{
		"A": "1",
		"B": "2",
		"C": "3",
	}
	got := s2.All()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestOpenStore_ErrorsWhenPathUnreadable(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/unreadable.store"
	if _, err := os.Create(path); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(path, 0o000); err != nil {
		t.Fatal(err)
	}
	_, err := kv.OpenStore(path)
	if err == nil {
		t.Fatal("no error")
	}
}

func TestOpenStore_ReturnsErrorOnInvalidData(t *testing.T) {
	t.Parallel()
	_, err := kv.OpenStore("testdata/invalid.store")
	if err == nil {
		t.Fatal("no error")
	}
}

func TestAllReturnsExpectedMap(t *testing.T) {
	t.Parallel()
	s, err := kv.OpenStore("dummy path")
	if err != nil {
		t.Fatal(err)
	}
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
