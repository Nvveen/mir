package storage

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestMain(m *testing.M) {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT)
	go func() {
		<-sigc
		TestDB.StopMongoTesting()
		log.Fatal("signal interrupt caught")
	}()
	defer func() {
		if r := recover(); r != nil {
			log.Println("recovering mongo error")
			TestDB.StopMongoTesting()
			log.Print(r)
		}
	}()
	TestDB = &testDB{&MongoDB{
		Host: "127.0.0.1",
		Port: "40001",
	}}
	err := TestDB.StartMongoTesting()
	if err != nil {
		panic(err)
	}
	// run tests
	ret := m.Run()
	err = TestDB.StopMongoTesting()
	if err != nil {
		panic(err)
	}
	os.Exit(ret)
}
