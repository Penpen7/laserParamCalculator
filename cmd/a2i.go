package cmd

import (
	"math"
	"strconv"

	"github.com/spf13/cobra"
)

var a2iCmd = &cobra.Command{
	Use:   "a2i",
	Short: "規格化強度a0からレーザー強度I0を計算します",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

    // 引数を受け取る
		a0, err := strconv.ParseFloat(args[0], 64)

		if err != nil {
			return
		}

    // 角周波数, 電場, 強度, 磁場を計算
		omega := 2 * math.Pi * c / (float64(lambda) * 1e-9)
		E := m * c * omega * a0 / e
		I := 0.5 * epsilon * E * E * c * 1e-4
		B := E / c

		print_result(I, E, B, a0, 35, 36)
	},
}
