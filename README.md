# tickers
tickers is a cli tool to detect arbitrage opportunities among supported exchanges.

## Usage
`tickers [flags]` <br> `tickers [command]` <br>

## Available Commands
`arb`                     get arbitrage opportunities <br>
`common`                  list common symbols for 2 or more exchanges <br>
`help`                    help about any command <br>
`exchanges`               list supported exchanges <br>
`tickers`                 get tickers for 1 or more excanges <br>

<br>

## Flags
`-h, --help`              help for tickers <br>
`-v, --verbose`           verbose output <br>

<br>

Use `tickers [command] --help` for more information about a command.

## Build
To build depending of the OS.

`make build-linux`        for linux build.
`make build-mac`          for mac build.
`make build-win`          for windows build.
