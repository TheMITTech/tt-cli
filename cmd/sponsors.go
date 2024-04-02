package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sponsorsCmd)
}

type sponsor struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	ImageURL string `json:"imageUrl"`
}

type input struct {
	Sponsors []sponsor `json:"sponsors"`
}

var sponsorsCmd = &cobra.Command{
	Use:   "sponsors [json]",
	Short: "Generates the 'sponsors_content' escaped string",
	Long: `Generates escaped string for 'sponsors_content' dynamic variable.
Input JSON file should be formated:
{
  "sponsors": [
    {
      "name": "The Tech",
      "url": "https://web.mit.edu",
      "imageUrl: "https://brand.mit.edu/sites/default/files/styles/image_text_1x/public/2023-08/MIT-logo-red-textandimage.png?itok=otGzsFBb"
    }
  ]
}
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open(args[0])
		if err != nil {
			fmt.Printf("Unable to open %s\n", args[0])
			os.Exit(1)
			return
		}

		var jsonData input
		if err = json.NewDecoder(file).Decode(&jsonData); err != nil {
			fmt.Printf("Unable to read %s\n", args[0])
			os.Exit(1)
			return
		}

		var output strings.Builder
		output.WriteString("<div class=\"sponsor-list\">\r\n")
		for _, s := range jsonData.Sponsors {
			output.WriteString("  <a class=\"sponsor-item\" href=\"")
			output.WriteString(s.URL)
			output.WriteString("\" target=\"_blank\" title=\"")
			output.WriteString(s.Name)
			output.WriteString("\">\r\n    <img src=\"")
			output.WriteString(s.ImageURL)
			output.WriteString("\"/>\r\n  </a>\r\n")
		}
		output.WriteString("</div>")

		fmt.Println(output.String())
	},
}
