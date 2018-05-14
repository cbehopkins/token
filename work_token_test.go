package token

import (
	"log"
	"testing"
	"time"
)

func TestWorkToken0(t *testing.T) {
	wt := NewWkTok(10)
	wt.Get()
	log.Println("Got")
	wt.Put()
	log.Println("Put")
	log.Println("Waiting")
	wt.Wait()

}
func TestWorkToken1(t *testing.T) {
	wt := NewWkTok(10)
	wt.Get()
	wt.Put()
	log.Println("Waiting")
	wt.Wait()
}

func TestWorkToken2(t *testing.T) {
	wt := NewWkTok(10)
	wt.Get()
	go func() {
		<-time.After(1 * time.Second)
		wt.Put()
	}()
	log.Println("Waiting")
	wt.Wait()
}
func TestWorkToken3(t *testing.T) {
	wt := NewWkTok(2)
	wt.Get()
	log.Println("Got")
	wt.Get()
	log.Println("Got")
	wt.Put()
	log.Println("Put")
	wt.Get()
	log.Println("Got")
	go func() {
		<-time.After(1 * time.Second)
		wt.Put()
		log.Println("Put")
	}()
	go func() {
		<-time.After(2 * time.Second)
		wt.Put()
		log.Println("Put")
	}()
	log.Println("Waiting")
	wt.Wait()
}
func TestWorkToken4(t *testing.T) {
	wt := NewWkTok(2)
	wt.Get()
	log.Println("Got")
	wt.Get()
	log.Println("Got")
	go func() {
		<-time.After(1 * time.Second)
		wt.Put()
		log.Println("Put")
	}()
	go func() {
		<-time.After(2 * time.Second)
		wt.Put()
		log.Println("Put")
	}()
	wt.Get()
	log.Println("Got")
	wt.Put()
	log.Println("Put")
	log.Println("Waiting")
	wt.Wait()
}
