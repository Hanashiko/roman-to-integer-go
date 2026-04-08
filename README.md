# Roman to Integer (Go)

A simple Go program that converts Roman numeral strings to their integer (Arabic) representation.

## How it works

The program iterates through the Roman numeral string, comparing each symbol with the next one:
- If the current value is **greater than or equal** to the next, it is added to the result.
- If the current value is **less** than the next (e.g., `IV`, `IX`), the difference is added instead (subtractive notation).

## Usage

```bash
go run main.go
```

By default, it converts `MCMXCIV` (1994) and prints the result.
