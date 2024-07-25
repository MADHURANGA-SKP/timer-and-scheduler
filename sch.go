package main

import (
	"log"
	"time"
)

func runScheduler(tm string, freq time.Duration) {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
	   log.Fatalf("ERROR: %v", err)
	}
	tm_, _ := time.ParseInLocation("15:04:05", tm, loc)
	stm := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), tm_.Hour(), tm_.Minute(), tm_.Second(), 0, tm_.Location())
   
	duration := stm.Sub(time.Now().In(loc))
   
	var timer *time.Timer
   
	timer = time.AfterFunc(duration, func() {
	   timer.Reset(freq)
	 
	   // Run scheduled functions
	   log.Println("Running Scheduled Functions..")
   
	})
 }
 