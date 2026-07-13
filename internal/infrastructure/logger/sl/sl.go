package sl

import (
	"fmt"
	"log/slog"
)

func Err(err error) slog.Attr {
	fmt.Println(err.Error())
	fmt.Println(slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	})
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
