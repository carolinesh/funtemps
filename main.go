package main

import (
	"flag"
	"fmt"
	"github.com/carolinesh/funtemps/conv"
)

var fahr float64
var out string
var funfacts string
var cels float64
var kelv float64

func init() {
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&cels, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelv, "K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&out, "out", "C, K, F", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
}

func main() {
	flag.Parse()

	if isFlagPassed("out") {
		if out == "C" && isFlagPassed("F") {
			fmt.Println(conv.FahrenheitToCelsius(fahr))
		} else if out == "F" && isFlagPassed("C") {
			fmt.Println(conv.CelsiusToFahrenheit(cels))
		}
	}

	if isFlagPassed("out") {
		if out == "K" && isFlagPassed("C") {
			fmt.Println(conv.KelvinToCelsius(cels))
		} else if out == "C" && isFlagPassed("K") {
			fmt.Println(conv.CelsiusToKelvin(kelv))
		}
	}

	if isFlagPassed("out") {
		if out == "K" && isFlagPassed("F") {
			fmt.Println(conv.FahrenheitToKelvin(kelv))
		} else if out == "F" && isFlagPassed("K") {
			fmt.Println(conv.KelvinToFahrenheit(fahr))
		}
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
