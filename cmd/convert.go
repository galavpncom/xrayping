package cmd

import (
	"fmt"

	"github.com/dev7dev/uri-to-json/pkgs/outbound"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/spf13/cobra"
)

// ShowOutboundStr prints the JSON string in an indented format
func ShowOutboundStr(oStr string) {
	j := gjson.New(oStr)
	fmt.Println(j.MustToJsonIndentString())
}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert VPN URI to outbound JSON for Sing-Box or Xray",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Usage: xrayping convert <sing|xray> <vpn-uri>")
			return
		}

		subCmd := args[0]
		rawUri := args[1]

		// Depending on the sub-command, handle Sing-Box or Xray conversion
		switch subCmd {
		case "sing":
			ob := outbound.GetOutbound(outbound.SingBox, rawUri)
			ob.Parse(rawUri)
			ShowOutboundStr(ob.GetOutboundStr())
		case "xray":
			ob := outbound.GetOutbound(outbound.XrayCore, rawUri)
			ob.Parse(rawUri)
			ShowOutboundStr(ob.GetOutboundStr())
		default:
			fmt.Println("Invalid sub-command. Use 'sing' or 'xray'.")
		}
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
}
