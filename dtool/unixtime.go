package dtool

import (
	"fmt"
	"io"
	"time"
)

func PrintUnixtime(w io.Writer, unixtime int64) {
	t := time.Unix(unixtime, 0).In(time.Local).Format("2006-01-02 15:04:05")
	_, _ = fmt.Fprintf(w, "%s (%d)", t, unixtime)
}
