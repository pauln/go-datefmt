package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	inTimestamp := flag.Int64("ts", time.Now().Unix(), "Date/time to convert, as a Unix timestamp")
	inStr := flag.String("in", "", "Date/time to convert, as a string")
	inFmt := flag.String("infmt", "", "Format of input string (-in)")

	outFmt := flag.String("fmt", "RFC3339", "Desired output format")
	outUTC := flag.Bool("utc", false, "Output in UTC")

	flag.Parse()

	in, err := getDate(*inTimestamp, *inStr, *inFmt)
	if err != nil {
		panic("Error parsing input date: " + err.Error())
	}

	outputFmt := getFormat(*outFmt)

	if *outUTC {
		fmt.Println(in.UTC().Format(outputFmt))
	} else {
		fmt.Println(in.Format(outputFmt))
	}
}

func getDate(ts int64, inStr string, inFmt string) (*time.Time, error) {
	if len(inStr) > 0 && len(inFmt) > 0 {
		inFmt = getFormat(inFmt)
		in, err := time.Parse(inFmt, inStr)
		if err != nil {
			return nil, err
		}
		return &in, nil
	}

	in := time.Unix(ts, 0)

	return &in, nil
}

func getFormat(inFmt string) string {
	switch inFmt {
	case "ANSIC":
		return time.ANSIC
	case "UnixDate":
		return time.UnixDate
	case "RubyDate":
		return time.RubyDate
	case "RFC822":
		return time.RFC822
	case "RFC822Z":
		return time.RFC822Z
	case "RFC850":
		return time.RFC850
	case "RFC1123":
		return time.RFC1123
	case "RFC1123Z":
		return time.RFC1123Z
	case "RFC3339":
		return time.RFC3339
	case "RFC3339Nano":
		return time.RFC3339Nano
	case "Kitchen":
		return time.Kitchen
	case "Stamp":
		return time.Stamp
	case "StampMilli":
		return time.StampMilli
	case "StampMicro":
		return time.StampMicro
	case "StampNano":
		return time.StampNano
	}

	return inFmt
}
