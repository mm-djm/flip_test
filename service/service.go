package service

import (
	"fmt"
	"strconv"
)

/**
	Name      Price   Inventory
	Google    49.99   10
	MacBook   5399.99 5
	Alexa     109.50  10
	Raspberry 30      2

Promotions:
1.Each sale of a MacBook Pro comes with a free Raspberry Pi B
2.Buy 3 Google Homes for the price of 2
3.Buying more than 3 Alexa Speakers will have a 10% discount on all Alexa
speakers
**/

func checkout(items []string) float64 {
	var (
		result float64
		google float64
		mac    float64
		alexa  float64
		rasp   float64
	)

	// count every kind of item
	for _, item := range items {
		switch item {
		case "Google Home":
			google += 1
		case "MacBook Pro":
			mac += 1
		case "Alexa Speaker":
			alexa += 1
		case "Raspberry Pi B":
			rasp += 1
		default:
			continue
		}
	}

	// count google home
	if google != 0 {
		if google > 10 {
			google = 10
		}
		i := int(google) / 3
		j := int(google) % 3
		result += float64(i)*2*49.99 + float64(j)*49.99
	}

	// count macbook and minus free rasp
	if mac != 0 {
		if mac > 5 {
			mac = 5
		}
		if rasp > 2 {
			rasp = 2
		}
		result += mac * 5399.99
		if mac >= rasp {
			rasp = 0
		} else {
			rasp = rasp - mac
		}
	}

	// count rasp
	if rasp != 0 {
		if rasp > 2 {
			rasp = 2
		}
		result += rasp * 30.00
	}

	// count alexa
	if alexa >= 3 {
		if alexa > 10 {
			alexa = 10
		}
		result += alexa * 109.50 * 0.9
	} else {
		result += alexa * 109.50
	}

	// keep two decimal places
	result, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", result), 64)
	return result
}
