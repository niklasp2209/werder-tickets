# Werder Bremen Ticket Availability Notifier

This project monitors the official Werder Bremen ticketing website for available tickets. The program checks the website every 5 seconds and sends an email notification when tickets become available for purchase.

## Features

- **Real-Time Ticket Monitoring:** Periodically checks for available tickets for Werder Bremen matches.
- **Email Notifications:** Sends an email when tickets are available for purchase.
- **Customizable Configuration:** Easily configure the email sender and recipient for notifications.

## Prerequisites

Before using the project, make sure you have the following:

- [Go](https://golang.org/dl/) (version 1.23 or higher)
- A valid email address for sending notifications (e.g., Gmail, Outlook, etc.)

## Installation

### 1. Clone the Repository

Clone this repository to your local machine using the following command:

```bash
git clone https://github.com/niklasp2209/werder-tickets/.git
cd werder-tickets
