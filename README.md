# Blaze calculator

I wanted to build a Spotlight alternative for Linux and decided to start
implementing the calculator.

I skipped all the GUI, as I wanted to get a working prototype as soon as
possible.

This turned out very useful even without the rest of Spotlight's features, so
here it is.

> Absolutely experimental, use at your own risk.

## Built with AI

This project has been built in about 4 hours. The bulk of tokenization, parsing
and evaluation has been iteratively built by ChatGPT with my (human)
interventions where it couldn't figure out the right way to proceed.

GitHub Copilot was also quite helpful in suggesting code snippets, especially
around boilerplate setup and error handling.

## Example

![demo](demo.gif?)

## How to use

1. Clone the repository
2. Compile the code: `go build`
3. Run the code: `./blaze`
4. Type in your math expression, the program will print the result as you type.

## Features

1. Real-time evaluation of mathematical expressions
2. Basic arithmetic operations: `+`, `-`, `*`, `/`, `^` (or `**`)
3. Nested expressions with parentheses
4. Percentages (e.g., `10% * 200`, `10% of 200`)
5. Modulo operation (e.g., `10 % 3`, `10 mod 3`)
6. Constants: `pi`, `e`, `phi`
7. Trigonometric functions: `sin`, `cos`, `tan`, `asin`, `acos`, `atan`
8. Logarithmic functions: `log`, `ln`
9. Square root: `sqrt`
10. Copy the result to clipboard by pressing `Enter`
11. Unit conversions:
    - weight (`10 kg to oz`)
    - length (`10 m to yd`)
    - area (`10 sqm to sqft`)
    - volume (`10 L to cup`)
    - time (`1d to s`)
    - temperature (`0c to f`)
12. Currency conversion: `10 usd to eur`, `10 eur to usd`, `1 btc to usd`
13. All of the above can be combined in a single expression:
    - `10 usd to eur * 5 + 10%`
    - `sqrt(10 kg to g)`
    - `cos(2*pi) * 2 btc to usd`

**Note**: currency rates are obtained from
[ExchangeRate-API](https://www.exchangerate-api.com/) and
[Coinbase API](https://developers.coinbase.com/api/v2).

Currency rates are cached for 1 hour (1 minute for cryptocurrencies) to avoid
hitting the API too frequently. The cache is stored in an appropriate location 
based on the OS (see `appDataPath`)

