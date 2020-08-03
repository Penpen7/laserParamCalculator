package cmd

import (
	"fmt"
)

func print_result(I float64, E float64, B float64, a0 float64, hilightI int, hilighta0 int) {
	lambda_micron := float64(lambda) * 1e-3
	fmt.Printf("波長           λ = %d nm (%G um)\n", lambda, lambda_micron)
	fmt.Printf("\x1b[%dmレーザー強度  I0 = %G W/cm^2\x1b[0m\n", hilightI, I)
	fmt.Printf("電場          E0 = %G V/m \n", E)
	fmt.Printf("磁場          B0 = %G T\n", B)
	fmt.Printf("\x1b[%dm規格化強度    a0 = %G\x1b[0m\n", hilighta0, a0)
}
