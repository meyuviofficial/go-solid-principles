package dip

import "fmt"

// this package demonstrates the Dependency Inversion Principle (DIP)

// bad example: High-level depends on low-level concretes (violates DIP)
/*
// High-level module tightly coupled to concrete senders
type BadNotificationService struct{}

func (BadNotificationService) NotifyEmail(addr, message string) {
	sender := &BadEmailSender{EmailAddress: addr} // direct concrete dependency
	sender.Send(message)
}

func (BadNotificationService) NotifySMS(phone, message string) {
	sender := &BadSMSSender{PhoneNumber: phone} // direct concrete dependency
	sender.Send(message)
}

// Low-level concrete modules without abstraction
type BadEmailSender struct{ EmailAddress string }

func (b *BadEmailSender) Send(message string) {
	fmt.Printf("[BAD] Email to %s: %s\n", b.EmailAddress, message)
}

type BadSMSSender struct{ PhoneNumber string }

func (b *BadSMSSender) Send(message string) {
	fmt.Printf("[BAD] SMS to %s: %s\n", b.PhoneNumber, message)
}

// Problems:
// - High-level policy (notifications) depends on concrete details (email/SMS structs)
// - Adding Push/Slack/etc. requires modifying BadNotificationService (closed to extension)
// - Hard to test: you can't substitute a mock without changing code
*/

// good example: High-level depends on abstractions (respects DIP)
type Sender interface {
	Send(message string)
}

type NotificationSender struct {
	sender Sender // depends on abstraction
}

func NewNotificationSender(s Sender) *NotificationSender {
	return &NotificationSender{sender: s}
}

func (n *NotificationSender) Notify(message string) {
	n.sender.Send(message)
}

// Low-level concrete modules implement the Sender interface
type EmailSender struct {
	EmailAddress string
}

func (e *EmailSender) Send(message string) {
	fmt.Printf("Email to %s: %s\n", e.EmailAddress, message)
}

type SMSSender struct {
	PhoneNumber string
}

func (s *SMSSender) Send(message string) {
	fmt.Printf("SMS to %s: %s\n", s.PhoneNumber, message)
}

// Example usage
func Example() {
	var (
		emailSender NotificationSender
		smsSender   NotificationSender
	)

	emailSender = *NewNotificationSender(&EmailSender{EmailAddress: "user@example.com"})
	smsSender = *NewNotificationSender(&SMSSender{PhoneNumber: "+1234567890"})

	emailSender.Notify("Hello via Email!")
	smsSender.Notify("Hello via SMS!")
}
