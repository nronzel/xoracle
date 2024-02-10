# XORacle

XORacle is a tool designed to decrypt data encoded with a repeating-key XOR
cipher. Utilizing brute-force methods alongside transposition and frequency
analysis techniques, XORacle will attempt to deduce the key size and the key
itself.

## ToDo

- [ ] Clear the output on each request
- [x] Finish README

## Features

- **Automatic Key Size Detection**: Uses statistical analysis to guess the
  most probable key sizes.
- **Frequency Analysis**: Utilizes English language frequency analysis to
  suggest the most likely keys.
- **Base64 and Hex support**: Automatically detects and processes input data
  encoded in Base64 and Hexadecimal formats.
- **User-Friendly Interface**: Front-end written with HTMX to provide a basic
  interface.

## Installation

### Prerequisites

Go 1.15 or later

### Steps

1. Clone the repository:

```sh
git clone https://github.com/nronzel/xoracle.git
```

2. Navigate to the project directory:

```sh
cd xoracle
```

3. Install dependencies:

```sh
go mod tidy
```

4. Build and run the project:

```sh
go build -o xoracle && ./xoracle
```

5. Open your browser and navigate to:

```text
localhost:8080/
```

## About

This project was created while going through the [CryptoPals](https://cryptopals.com/)
challenge to get more familiar with cryptography. Specifically,
[Set 1 Project 6](https://cryptopals.com/sets/1/challenges/6).

## How it Works

### Identifying Key Sizes

Begin by identifying potential key sizes. XORacle employs a heuristic based on
the Hamming distance (the number of differing bits) between the blocks of
ciphertext. By analyzing the distances between blocks of various sizes, we can
make educated guesses about the most probable key sizes. The assumption is that
the correct key size with result in the smallest average normalized Hamming
distance because correctly sizes blocks aligned witht the repeating key will have
more similar bit patterns.

### Key Size Validation

With a set of potential key sizes, XORacle then divides the ciphertext into blocks
of each guessed key size. For each key size, the blocks are transposed to align
with the `nth` byte of each block into the new blocks. This effectively groups
together all bytes encrypted with the same byte of the turning the problem into
multiple single-byte XOR cipher problems.

### Frequency Analysis

For each transposed block, XORacle applies frequncy analysis. Presuming the data
is in English plaintext, the frequency of characters in these transposed blocks
is compared against known English language frequency statistic. Each byte in the
range of 0x00 and 0xFF is tried as the key for the single-byte XOR, and the output
is scored based on how closely it matches the expected English text character
frequencies.

### Determining Best Key

After scoring each potential key byte for each position in the key, XORacle combines
the highest-scoring bytes to form the keys for each guessed key size. It will then
attempt to decrypt the ciphertext using these keys and scores the resulting
plaintexts, using the same frequency analysis function from above. The key
that produces the most closely resembling English is selected as the most
likely key used for encryption.

## Testing

Run the included test suite with the following command:

```sh
go test ./... -v
```

## Contributing

Contributions to XORacle are welcome! If you have suggestions for improvements
or bug fixes, please fork this repo and create a pull request.
