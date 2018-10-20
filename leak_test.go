package scraper

import (
	"github.com/NavyaZaveri/scraper/testdata"
	"github.com/fortytw2/leaktest"
	"testing"
)

func TestGoRoutineLeak(t *testing.T) {

	//should throw a panic error if there's a leak
	leaktest.Check(t)()
	_ = NewWorkerPool(50).Fetch(&testdata.XkcdIterator{})

	t.Logf("No leaks")
}