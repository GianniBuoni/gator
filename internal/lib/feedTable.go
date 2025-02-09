package lib

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func NewFeedTable() *tabwriter.Writer {
	padding := 4
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	fmt.Fprintln(w, "NAME\tURL\tUSER\t")
  return w
}
