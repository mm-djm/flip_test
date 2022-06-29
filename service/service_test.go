package service

import (
	"fmt"
	"testing"
)

func TestGoogleOnly(t *testing.T) {
	cases := []struct {
		desc     string
		items    []string
		expected float64
	}{
		// 3 Google Home
		{
			desc:     "testThreeGoogle",
			items:    []string{"Google Home", "Google Home", "Google Home"},
			expected: 99.98,
		},
		// 11 Google Home
		{
			desc:     "testElevenGoogle",
			items:    []string{"Google Home", "Google Home", "Google Home", "Google Home", "Google Home", "Google Home", "Google Home", "Google Home", "Google Home", "Google Home", "Google Home"},
			expected: 349.93,
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			actual := checkout(c.items)
			if actual != c.expected {
				t.Errorf("Test failed, expected: '%v', got: '%v'", c.expected, actual)
				t.Fail()
			}
		})
	}
}

func TestMacAndRasp(t *testing.T) {
	cases := []struct {
		desc     string
		items    []string
		expected float64
	}{
		// 6 MacBook Pro
		{
			desc:     "testSixMacBook",
			items:    []string{"MacBook Pro", "MacBook Pro", "MacBook Pro", "MacBook Pro", "MacBook Pro", "MacBook Pro"},
			expected: 26999.95,
		},
		// 5 Raspberry
		{
			desc:     "testFiveRasp",
			items:    []string{"Raspberry Pi B", "Raspberry Pi B", "Raspberry Pi B", "Raspberry Pi B", "Raspberry Pi B"},
			expected: 60.00,
		},
		// 1 MacBook Pro + 2 Raspberry
		{
			desc:     "testMacAddRasp",
			items:    []string{"MacBook Pro", "Raspberry Pi B", "Raspberry Pi B"},
			expected: 5429.99,
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			actual := checkout(c.items)
			if actual != c.expected {
				t.Errorf("Test failed, expected: '%v', got: '%v'", c.expected, actual)
				t.Fail()
			}
		})
	}
}

func TestAlexaOnly(t *testing.T) {
	cases := []struct {
		desc     string
		items    []string
		expected float64
	}{
		// 2 Alexa Speaker
		{
			desc:     "testTwoAlexa",
			items:    []string{"Alexa Speaker", "Alexa Speaker"},
			expected: 219,
		},
		// 12 Alexa Speaker
		{
			desc:     "testTwelveAlexa",
			items:    []string{"Alexa Speaker", "Alexa Speaker", "Alexa Speaker", "Alexa Speaker", "Alexa Speaker", "Alexa Speaker", "Alexa Speaker", "Alexa Speaker", "Alexa Speaker", "Alexa Speaker", "Alexa Speaker", "Alexa Speaker"},
			expected: 985.5,
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			actual := checkout(c.items)
			if actual != c.expected {
				t.Errorf("Test failed, expected: '%v', got: '%v'", c.expected, actual)
				t.Fail()
			}
		})
	}
}

func TestCombination(t *testing.T) {
	cases := []struct {
		desc     string
		items    []string
		expected float64
	}{
		{
			desc:     "testNothing",
			items:    []string{},
			expected: 0,
		},
		// 10 Google + 5 MacBook + 10 Alexa + 2 Raspberry
		{
			desc:     "testAllCombination",
			items:    []string{"MacBook Pro", "MacBook Pro", "MacBook Pro", "MacBook Pro", "MacBook Pro", "Raspberry Pi B", "Raspberry Pi B", "Google Home", "Google Home", "Google Home", "Google Home", "Google Home", "Google Home", "Google Home", "Google Home", "Google Home", "Google Home", "Alexa Speaker", "Alexa Speaker", "Alexa Speaker", "Alexa Speaker", "Alexa Speaker", "Alexa Speaker", "Alexa Speaker", "Alexa Speaker", "Alexa Speaker", "Alexa Speaker"},
			expected: 28335.38,
		},
		// 2 MacBook + 3 Raspberry + 5 Google
		{
			desc:     "testCombination",
			items:    []string{"MacBook Pro", "MacBook Pro", "Raspberry Pi B", "Raspberry Pi B", "Raspberry Pi B", "Google Home", "Google Home", "Google Home", "Google Home", "Google Home"},
			expected: 10999.94,
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			actual := checkout(c.items)
			if actual != c.expected {
				t.Errorf("Test failed, expected: '%v', got: '%v'", c.expected, actual)
				t.Fail()
			}
		})
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Tests begin")
	m.Run()
}
