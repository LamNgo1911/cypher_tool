# Cypher Tool

A simple command-line cipher tool for encrypting and decrypting messages using three different cipher methods.

## Table of Contents

- [About the Project](#about-the-project)
- [Tool Introduction](#tool-introduction)
- [Tool Usage](#tool-usage)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [How Ciphers Work](#how-ciphers-work)
  - [ROT13](#rot13)
  - [Reverse](#reverse)
  - [Vigenère Cipher](#vigenère-cipher)

## About the Project

This project was developed as a collaborative group assignment at **Hive**.

We created a command-line cipher tool that allows users to encrypt and decrypt messages using different cipher techniques. The goal of the project was to practice programming concepts such as string manipulation, algorithms, user input, and modular arithmetic.

## Tool Introduction

The Cypher Tool supports three different cipher methods:

- **ROT13** — shifts each letter by 13 positions in the alphabet.
- **Reverse** — replaces each letter with its corresponding letter from the opposite end of the alphabet.
- **Vigenere Cipher** — uses a keyword to determine the shift applied to each letter.

The tool supports both **encryption** and **decryption** operations.

## Tool Usage

### Choose an Operation

When you start the program, you can choose whether to encrypt or decrypt a message:

```text
Select operation (1/2):
1. Encrypt
2. Decrypt
```

### Choose a Cipher

Next, select one of the available cipher methods:

```text
Select cipher (1/3):
1. ROT13
2. Reverse
3. Vigenere Cipher
```

### Enter Your Message

Enter the message you want to encrypt or decrypt:

```text
Enter the message:
hello
```

For example, using ROT13, the program will output:

```text
Encrypted message using ROT13:
uryyb
```

## Getting Started

### Prerequisites

You need to have **Go** installed on your system.

Go is available for:

- Linux
- macOS
- Windows

You can download and install Go from the [official Go website](https://go.dev/doc/install).

### Installation

Clone the repository and navigate into the project directory:

```bash
git clone [https://gitea.kood.tech/quanglamngo/cypher_tool.git](https://gitea.kood.tech/quanglamngo/cypher_tool.git)
cd cypher_tool
```

Run the program with:

```bash
go run .
```

Or, compile it into an executable:

```bash
go build
```

Then, you can run the compiled program with:

```bash
./cypher_tool
```

## How Ciphers Work

### ROT13

ROT13 (Rotate by 13 places) shifts every letter in the alphabet by 13 positions.

Because the alphabet contains 26 letters, applying ROT13 twice returns the original message.

**Encryption:**

```text
abcd → nopq
```

**Decryption:**

```text
nopq → abcd
```

ROT13 is symmetric, meaning the same operation is used for both encryption and decryption.

### Reverse

The Reverse cipher replaces each letter with the corresponding letter from the opposite end of the alphabet.

For example:

```text
a → z
b → y
c → x
d → w
```

Therefore:

**Encryption:**

```text
abcd → zyxw
```

**Decryption:**

```text
zyxw → abcd
```

Like ROT13, the Reverse cipher uses the same operation for encryption and decryption.

### Vigenere Cipher

The Vigenere cipher uses a **keyword** to determine how far each letter in the message should be shifted.

Each letter of the keyword represents a number from **0 to 25**, corresponding to its position in the alphabet:

```text
a = 0
b = 1
c = 2
...
z = 25
```

For example, if the keyword is:

```text
lam
```

The letters represent the following shifts:

```text
l = 11
a = 0
m = 12
```

The keyword is repeated as necessary to match the length of the message.

For example:

```text
Message:  abcd
Keyword:  lama
```

The corresponding shifts are then applied to each letter to produce the encrypted message.

The same keyword is required to decrypt the message back to its original form.