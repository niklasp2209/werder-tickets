package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"net/smtp"
	"time"
)

// sendEmail sends a notification email.
func sendEmail(subject, body string) {
	from := "sender@example.com"
	to := "target@example.com"
	password := "password"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	auth := smtp.PlainAuth("", from, password, "smtp.mail.com")
	err := smtp.SendMail("smtp.mail.com:587", auth, from, []string{to}, []byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}

// checkTickets monitors the ticketing page and triggers email if tickets are available.
func checkTickets() {
	c := colly.NewCollector()

	// Check for "Tickets Kaufen" button that is not disabled
	c.OnHTML("a.button", func(e *colly.HTMLElement) {
		buttonText := e.Text
		if buttonText == "Tickets Kaufen" && !e.Attr("class").contains("button--disabled") {
			sendEmail("Neue Tickets verfügbar!", "Es sind neue Tickets auf der Werder Bremen Seite verfügbar!")
		} else {
			fmt.Println("Keine neuen Tickets gefunden.")
		}
	})

	err := c.Visit("https://www.ticket-onlineshop.com/ols/werderbremen/de/einzelkarten/channel/shop/index")
	if err != nil {
		log.Fatal(err)
	}
}

// main runs the ticket check every 5 seconds.
func main() {
	for {
		checkTickets()
		time.Sleep(5 * time.Second)
	}
}
