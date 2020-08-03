package cmd

import (
	"math"
	"strconv"

	"github.com/spf13/cobra"
)

var i2aCmd = &cobra.Command{
	Use:   "i2a",
	Short: "レーザー強度から規格化強度a0を計算します",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

    // 引数から強度を得る
		I, err := strconv.ParseFloat(args[0], 64)

		if err != nil {
			return
		}

    // 角周波数, 電場, 磁場, a0を計算する
		omega := 2 * math.Pi * c / (float64(lambda) * 1e-9)
		E := math.Sqrt(2*I/(epsilon*c)) * 100.0
		B := E / c
		a0 := e * E / (m * c * omega)
		print_result(I, E, B, a0, 36, 35)
	},
}
