package main

import "testing"
import "fmt"
import "log"
import "testing/quick"

func TestRomanNumerals(t *testing.T) {
		for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases[:4] {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic int) bool {
   if arabic < 0 || arabic > 3999 {
   	log.Println(arabic)
   	return true
   }
   roman := ConvertToRoman(arabic)
   fromRoman := ConvertToArabic(roman)
   return fromRoman == arabic
}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}
