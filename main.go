package main

func main() {
	nextDate := getCountDownDate()

	screen := NewScreen()
	ch := make(chan Date)

	go handleKey(screen)
	go countdown(nextDate, ch)
	Run(screen, nextDate, ch)
}
