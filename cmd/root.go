package cmd

import (
	"fmt"
	"os"
	"sync"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/whs-dot-hk/go-rs3-gem-bag-calculator/item"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gemcal",
		Short: "A gem bag calculator",
		Long:  "Calculate the total price of the gems inside a gem bag.",
		Run: func(cmd *cobra.Command, args []string) {
			var wg sync.WaitGroup

			if totalUncutSapphires > 0 {
				wg.Add(1)
				go func() {
					uncutSapphirePrice = uncutSapphire.GetPrice()
					wg.Done()
				}()
			}

			if totalUncutEmeralds > 0 {
				wg.Add(1)
				go func() {
					uncutEmeraldPrice = uncutEmerald.GetPrice()
					wg.Done()
				}()
			}

			if totalUncutRubies > 0 {
				wg.Add(1)
				go func() {
					uncutRubyPrice = uncutRuby.GetPrice()
					wg.Done()
				}()
			}

			if totalUncutDiamonds > 0 {
				wg.Add(1)
				go func() {
					uncutDiamondPrice = uncutDiamond.GetPrice()
					wg.Done()
				}()
			}

			if totalUncutDragonstones > 0 {
				wg.Add(1)
				go func() {
					uncutDragonstonePrice = uncutDragonstone.GetPrice()
					wg.Done()
				}()
			}

			wg.Wait()

			totalUncutSapphiresPrice := uncutSapphirePrice * totalUncutSapphires
			totalUncutEmeraldsPrice := uncutEmeraldPrice * totalUncutEmeralds
			totalUncutRubiesPrice := uncutRubyPrice * totalUncutRubies
			totalUncutDiamondsPrice := uncutDiamondPrice * totalUncutDiamonds
			totalUncutDragonstonesPrice := uncutDragonstonePrice * totalUncutDragonstones

			totalPrice := totalUncutSapphiresPrice + totalUncutEmeraldsPrice + totalUncutRubiesPrice + totalUncutDiamondsPrice + totalUncutDragonstonesPrice

			var data [][]string

			if totalUncutSapphires > 0 {
				row := []string{
					fmt.Sprintf("%d", item.UncutSapphireCode),
					"Uncut sapphire",
					fmt.Sprintf("%d", totalUncutSapphires),
					fmt.Sprintf("%d", uncutSapphirePrice),
					fmt.Sprintf("%d", totalUncutSapphiresPrice),
				}
				data = append(data, row)
			}

			if totalUncutEmeralds > 0 {
				row := []string{
					fmt.Sprintf("%d", item.UncutEmeraldCode),
					"Uncut emerald",
					fmt.Sprintf("%d", totalUncutEmeralds),
					fmt.Sprintf("%d", uncutEmeraldPrice),
					fmt.Sprintf("%d", totalUncutEmeraldsPrice),
				}
				data = append(data, row)
			}

			if totalUncutRubies > 0 {
				row := []string{
					fmt.Sprintf("%d", item.UncutRubyCode),
					"Uncut ruby",
					fmt.Sprintf("%d", totalUncutRubies),
					fmt.Sprintf("%d", uncutRubyPrice),
					fmt.Sprintf("%d", totalUncutRubiesPrice),
				}
				data = append(data, row)
			}

			if totalUncutDiamonds > 0 {
				row := []string{
					fmt.Sprintf("%d", item.UncutDiamondCode),
					"Uncut diamond",
					fmt.Sprintf("%d", totalUncutDiamonds),
					fmt.Sprintf("%d", uncutDiamondPrice),
					fmt.Sprintf("%d", totalUncutDiamondsPrice),
				}
				data = append(data, row)
			}

			if totalUncutDragonstones > 0 {
				row := []string{
					fmt.Sprintf("%d", item.UncutDragonstoneCode),
					"Uncut dragonstone",
					fmt.Sprintf("%d", totalUncutDragonstones),
					fmt.Sprintf("%d", uncutDragonstonePrice),
					fmt.Sprintf("%d", totalUncutDragonstonesPrice),
				}
				data = append(data, row)
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Item code", "Name", "Quantity", "Price", "Amount"})
			table.SetFooter([]string{"", "", "", "Total", fmt.Sprintf("%d", totalPrice)}) // Add Footer
			table.SetBorder(false)                                                        // Set Border to false
			table.AppendBulk(data)                                                        // Add Bulk Data
			table.Render()
		},
	}

	uncutSapphire    *item.Item = item.NewUncutSapphire()
	uncutEmerald     *item.Item = item.NewUncutEmerald()
	uncutRuby        *item.Item = item.NewUncutRuby()
	uncutDiamond     *item.Item = item.NewUncutDiamond()
	uncutDragonstone *item.Item = item.NewUncutDragonstone()

	uncutSapphirePrice    int = 0
	uncutEmeraldPrice     int = 0
	uncutRubyPrice        int = 0
	uncutDiamondPrice     int = 0
	uncutDragonstonePrice int = 0

	totalUncutSapphires    int
	totalUncutEmeralds     int
	totalUncutRubies       int
	totalUncutDiamonds     int
	totalUncutDragonstones int
)

func init() {
	rootCmd.PersistentFlags().IntVarP(&totalUncutSapphires, "sapphire", "b", 0, "Number of uncut sapphires")
	rootCmd.PersistentFlags().IntVarP(&totalUncutEmeralds, "emerald", "g", 0, "Number of uncut emeralds")
	rootCmd.PersistentFlags().IntVarP(&totalUncutRubies, "ruby", "r", 0, "Number of uncut rubys")
	rootCmd.PersistentFlags().IntVarP(&totalUncutDiamonds, "diamond", "w", 0, "Number of uncut diamonds")
	rootCmd.PersistentFlags().IntVarP(&totalUncutDragonstones, "dragonstone", "p", 0, "Number of uncut dragonstones")
}

func Execute() error {
	return rootCmd.Execute()
}
