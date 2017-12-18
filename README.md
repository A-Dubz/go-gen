# Go-gen

Command to generate a random or multiple alphanumberic strings with optional symbols

## Usage

To run the script with default args, run:
```bash
gen [num]
```
With [num] being any integer between 1-100.

The following args can be passed to config the script:

| Arg   | Type       | Value                      | Description                           |
| ----- | ---------- | ---------------------------| ------------------------------------- |
| `-s`  | optional   | `<bool>`, default: `false` | Add symbols to generated string       |
| `-r`  | optional   | `<integer>`, default: `1`  | Number of generated strings to output |
| [num] | positional | `<integer>`                | Length of generated string (1-100)    |

Example usage to generate a random string w/ symbols, repeated 5 times with length of 20 characters each:
```bash
~$ gen -s -r 5 20
MkGGi(FR(^jCa*nsN7F0
cL+Ej)3)heirB}oS^bF@
KSRZPX%Y?oiGD?b{SGD!
1J@g9m%y6?CA+}WaA3LX
3+MfN5iPX+Ssnqyn562E
```
