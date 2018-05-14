package token

import (
	"log"
	"sync"
)

// MultiToken is used to check that we have a limited number of things happening at once
type MultiToken struct {
	sync.Mutex
	name   string
	maxCnt int
	array  map[string]*WkTok
}

// NewTokenChan creates a new token chan tracker
func NewTokenChan(numTokens int, name string) *MultiToken {
	var itm MultiToken
	itm.array = make(map[string]*WkTok)
	itm.name = name
	itm.maxCnt = numTokens
	return &itm
}

func (tc *MultiToken) getWk(basename string) *WkTok {
	tc.Lock()
	val, ok := tc.array[basename]
	if !ok {
		val = NewWkTok(tc.maxCnt)
		tc.array[basename] = val
	}
	tc.Unlock()
	return val
}

// Get get a token for the supplied basename
// blocks until it suceeds
func (tc *MultiToken) Get(basename string) {
	if basename == "" {
		return
	}
	tc.getWk(basename).GetTok()
}

// TryGet attempts to get a token
// return true if it got a token
// returns false if it fails
// does not block
func (tc *MultiToken) TryGet(basename string) bool {
	if basename == "" {
		return true
	}
	// This is a variant of Get that doesn't block
	// it returns true if it can give you a token
	// false otherwise
	//log.Printf("Token requested for \"%v\"\n", basename)
	return tc.getWk(basename).TryGetTok()
}

// Put returns the token after use
func (tc *MultiToken) Put(basename string) {
	if basename == "" {
		return
	}
	_, ok := tc.array[basename]
	if !ok {
		log.Fatal("Tried to put a token that doesn't exist", basename)
	}
	tc.getWk(basename).PutTok()
}

// TryPut alias to allow consustent naming
func (tc *MultiToken) TryPut(basename string) {
	tc.Put(basename)
}

// Exist - Check if the basename exists
func (tc *MultiToken) Exist(urx string) bool {
	tc.Lock()
	_, ok := tc.array[urx]
	tc.Unlock()
	return ok
}
