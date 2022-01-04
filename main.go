package main

import (
	"flag"
	"fmt"
	"github.com/anatolethien/go-notification"
	"time"
)

type Pomodoro struct {
	work int
	pause int
	cycle int
}

func main()  {
	pomodoro := new(Pomodoro)
	flags(pomodoro)
	fmt.Printf("Starting a Pomodoro...\n")
	for i := 0; i < pomodoro.cycle; i++ {
		notify(fmt.Sprintf("It's time to work for %d minutes.", pomodoro.work))
		time.Sleep(time.Duration(pomodoro.work) * time.Minute)
		notify(fmt.Sprintf("You can take a %d minutes break.", pomodoro.pause))
		time.Sleep(time.Duration(pomodoro.pause) * time.Minute)
	}
}

func flags(pomodoro *Pomodoro) {
	flag.IntVar(&pomodoro.work, "work", 25, "set work time")
	flag.IntVar(&pomodoro.pause, "pause", 5, "set pause time")
	flag.IntVar(&pomodoro.cycle, "cycle", 4, "set cycle number")
	flag.Parse()
}

func notify(body string) {
	notification := new(notification.Notification)
	notification.Title = "Pomodoro"
	notification.Body = body
	notification.Icon = "alarm"
	fmt.Printf("[%s] %s\n", notification.Title, notification.Body)
	notification.Push()
}
