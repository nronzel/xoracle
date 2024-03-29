![ci-status](https://img.shields.io/github/actions/workflow/status/nronzel/xoracle/ci.yml?label=ci)
![cd-status](https://img.shields.io/github/actions/workflow/status/nronzel/xoracle/cd.yml?label=cd)
![top-lang](https://img.shields.io/github/languages/top/nronzel/xoracle?logo=go)
![docker-pulls](https://img.shields.io/docker/pulls/sutats/xoracle?logo=docker)
![license](https://img.shields.io/github/license/nronzel/xoracle)
![tag](https://img.shields.io/github/v/tag/nronzel/xoracle)

# XORacle

XORacle is a simple tool aimed at decrypting data that's been encrypted using
a repeating-key XOR cipher. It combines a brute-force approach with transposition
and frequency analysis to try and figure out the encryption key's size, the
key itself, and attempts to decrypt the data with the derived key(s).

> This project is dockerized and gets deployed to Google Cloud Run.
>
> [View Hosted Site](https://xoracle-uzphfx7uwa-ue.a.run.app)
>
> _See the [Usage](#usage) section for examples to test it out._

## Features

- **Automatic Key Size Detection**: Uses statistical analysis to guess the
  most probable key sizes.
- **Frequency Analysis**: Utilizes English language frequency analysis to
  suggest the most likely keys.
- **Base64 and Hex support**: Automatically detects and processes input data
  encoded in Base64 and Hexadecimal formats.
- **User-Friendly Interface**: Front-end written with HTMX to provide a basic
  interface.

### Table of Contents

- [Endpoints](#endpoints)
- [Installation](#installation)
- [Docker](#docker)
- [Usage](#usage)
- [About](#about)
- [How it Works](#how-it-works)
- [Testing](#testing)
- [Contributing](#contributing)

## Endpoints

### Homepage

- **URL:** `/`
- **Method:** `GET`
- **Description:** Retrieves the homepage of the Xoracle web application. Includes
  a form where users can input encoded data for decryption.
- **Auth Required:** No
- **Parameters:** None
- **Success Responses:**
  - **Code:** `200 OK`
  - **Content-Type:** `text/html`
  - **Content:** HTML content of homepage

### Decrypt Data

- **URL:** `/decrypt`
- **Method:** `POST`
- **Description:** Accepts encoded (either Base64 or Hex format) via form submission
  and attempts to decrypt it using various key sizes. Returns the decrypted data
  alongside the key used for decryption and the key size.
- **Auth Required:** No
- **Parameters:**
  - **inputData** (formData): The encoded data to be decrypted
- **Success Response:**
  - **Code:** `200 OK`
  - **Content-Type:** `text/html`
  - **Content:** HTML containing the decryption results, including the key size,
    the key itself, and the decrypted data.
- **Error Responses:**
  - **Code: `400 Bad Request`**
    - **Content-Type:** `text/plain`
    - **Reason:** Data or key sizes are missing, or the form failed to parse.
  - **Code: `500 Internal Server Error`**
    - **Content-Type:** `text/plain`
    - **Reason:** An error occurred during data decoding or key size guessing.

## Installation

### Prerequisites

Go 1.22

### Steps

**1. Clone the repository:**

```bash
git clone https://github.com/nronzel/xoracle.git
```

**2. Navigate to the project directory:**

```bash
cd xoracle
```

**3. Install dependencies:**

- golang.org/x/time

Install dependencies with the command:

```bash
go mod tidy
```

**4. Build and run the project:**

Linux & MacOS:

```bash
go build -o xoracle && ./xoracle
```

Windows:

```bash
go build -o xoracle.exe && .\xoracle.exe
```

**5. Open your browser and navigate to:**

```bash
localhost:8080/
```

### Docker

If you'd like to run this in a Docker container:

**Using docker compose:**

```bash
# Builds the image and runs the container in the background
docker compose up --build -d
```

#### Alternatively (Build Image Yourself From Dockerfile)

You can build a docker image with the included Dockerfile yourself, and run
the image in a container.

While in the root of the project directory:

**Build the image:**

```bash
docker build . -t xoracle
```

**Run the image in a container:**

```bash
docker run -p 8080:8080 xoracle
```

You should then be able to navigate to `localhost:8080` in your browser to
see the hosted Docker application.

## Usage

Enter the encoded text in the input box and click the "Process" button. The
processed data will show in the output box, or an error will display on the screen.

Feel free to use the small example below, or the text from the included [example.txt](./example.txt)
file in the repository to test it out.

```bash
# Base64 encoded
MiciMCQ2YTYkOjViJTclJyQ=
```

or

```bash
# Hex encoded
3227223024366136243a35622537252724
```

## About

This project was created while going through the [CryptoPals](https://cryptopals.com/)
challenge to get more familiar with cryptography; specifically,
[Set 1 - Project 6](https://cryptopals.com/sets/1/challenges/6). I had a lot of
fun making this and saw it as a good chance to make a basic frontend
to get a little familiar with HTMX. I also learned quite a bit about XOR ciphers.

## How it Works

### Identifying Key Sizes

Begin by identifying potential key sizes. XORacle employs a heuristic based on
the Hamming distance (the number of differing bits) between the blocks of
ciphertext. By analyzing the distances between blocks of various sizes, we can
make educated guesses about the most probable key sizes. The assumption is that
the correct key size will result in the smallest average normalized Hamming
distance because correctly sized blocks aligned with the repeating key will have
more similar bit patterns.

### Key Size Validation

With a set of potential key sizes, XORacle then divides the ciphertext into blocks
of each guessed key size. For each key size, the blocks are transposed to align
with the `nth` byte of each block into the new blocks. This effectively groups
together all bytes encrypted with the same byte of the key, turning the problem into
multiple single-byte XOR ciphers.

The chart below helps visualize this transposition.

```mermaid
graph TD
    A[Block 1: byte1, byte2, ..., byteN] --> B((Transpose))
    C[Block 2: byte1, byte2, ..., byteN] --> B
    D[Block 3: byte1, byte2, ..., byteN] --> B
    B --> E[New Block 1: Block1.byte1, Block2.byte1, Block3.byte1, ...]
    B --> F[New Block 2: Block1.byte2, Block2.byte2, Block3.byte2, ...]
    B --> G[...]
    B --> H[New Block N: Block1.byteN, Block2.byteN, Block3.byteN, ...]

    classDef block fill:#fcc7f,stroke:#333,stroke-width:2px;
    class A,C,D block;
```

### Frequency Analysis

For each transposed block, XORacle applies frequncy analysis. Presuming the data
is in English plaintext, the frequency of characters in these transposed blocks
is compared against known English language frequency statistics. Each byte in the
range of 0x00 and 0xFF (0-255) is tried as the key for the single-byte XOR, and
the output is scored based on how closely it matches the expected English text
character frequencies.

### Determining Best Key

After scoring each potential key byte for each position in the key, XORacle combines
the highest-scoring bytes to form the keys for each guessed key size. It will then
attempt to decrypt the ciphertext using these keys and scores the resulting
plaintexts, using the same frequency analysis function from above. The key
that produces output that most closely resembles English is selected as the most
likely key used for encryption.

### Limitations

Only works on English language text. If the ciphertext is numbers, coordinates
of some kind, or any other language, XORacle will not be able to decrypt the data.
If you'd like other languages, feel free to submit a pull request with a frequency
map for the language, and a function that can detect the language.

## Testing

Run the included test suite with the following command:

```sh
go test ./... -v
```

## Contributing

Contributions to XORacle are welcome! If you have suggestions for improvements
or bug fixes, please fork this repo and create a pull request.
