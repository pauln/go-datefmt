# go-datefmt
Simple CLI date formatter written in Go

## Usage
`go-datefmt` accepts the following CLI parameters:

| Parameter | Default value       | Description                                         |
| --------- | ------------------- | --------------------------------------------------- |
| `-ts`     | `time.Now().Unix()` | Input date/time, as a UNIX timestamp                |
| `-in`     | [none]              | Input date/time, as an arbitrarily-formatted string |
| `-infmt`  | [none]              | Input date/time format (Go-style format string)     |
| `-fmt`    | `RFC3339`           | Output date/time format (Go-style format string)    |
| `-utc`    | `false`             | Output date/time in UTC                             |

For convenience, Go's [time format constants](https://golang.org/pkg/time/#pkg-constants)
can be used in both `-infmt` and `-fmt` parameters - Simply specify the name of the
constant instead of a [format string](https://golang.org/pkg/time/#Parse).

If both `-in` and `-infmt` parameters are specified, they will override the `-ts`
parameter.


## Examples
```
# Current date/time, in RFC3339 format
>go-datefmt
2017-07-17T16:21:35+12:00

# Current time, RCG3339 format, UTC
>go-datefmt -utc
2017-07-17T04:21:39Z

# Specified timestamp, in RFC822 format
>go-datefmt -ts 1465183975 -fmt RFC822
06 Jun 16 15:32 NZST

# Specified date/time string, converting from RFC822Z to RubyDate
>go-datefmt -in "02 Jan 06 15:04 -0700" -infmt RFC822Z -fmt RubyDate
Mon Jan 02 15:04:00 -0700 2006

# Specified date/time string, converting from RCG822Z to custom date format
>go-datefmt -in "31 Dec 99 23:59 +1200" -infmt RFC822Z -fmt "2006/01/02 3:04PM"
1999/12/31 11:59PM
```
