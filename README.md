# conv

Oh great, another number converter. Like the world needed more of those, right? But hey, you asked for it, so here it is. A CLI tool to convert numbers between decimal, binary, octal, and hexadecimal. Riveting stuff.

## Features

- Converts numbers between decimal, binary, octal, and hexadecimal. Because why not.
- Supports prefixes like `0x` for hex, `0b` for binary, and `0o` for octal. Aren't these prefixes just delightful?
- Prints the converted numbers with optional prefixes and labels. Because we all love options.
- Handles multiple numbers in one go. Efficiency at its finest.

## Installation

Sure, just clone this amazing repository and build it yourself. You know the drill.

```sh
git clone https://github.com/sett17/conv.git
cd conv
go build -o conv
```

## Usage

Run the `conv` command with your preferred options. Or don't. I'm not your boss.

```sh
./conv [options] <number> [<number> ...]
```

### Options

- `--dec` Print decimal. Wow, a decimal number.
- `--hex` Print hexadecimal. Because 0xDE is just too cool.
- `--oct` Print octal. Like anyone really uses octal.
- `--bin` Print binary. 0b1010, because why not.
- `--no-prefix` Do not print prefixes. Because minimalism is a thing.
- `--no-label` Do not print labels. You like it raw, I get it.

### Examples

Converting a single number:

```sh
./conv 0x05
```

Output:

```
dec: 5
hex: 0x5
oct: 0o5
bin: 0b101
```

Converting multiple numbers, without prefixes and labels. Who needs those anyway?

```sh
./conv --no-prefix --no-label 0x05 69 0b1010
```

Output:

```
5
5
5
5

69
45
105
1001101

10
2
12
1010
```

## License

Do whatever you want with it. Seriously, I'm done. 

---

Feel free to contribute or don't. It's not like I'm expecting a flood of pull requests. Enjoy, or don't. Up to you.

