genlink
---

This is a command line tool that generates the link to the specified URL in various formats.

## Features

This tool can generate a link to the specified URL in the following format.

- Markdown
- HTML a tag
- HTML a tag with `target="_blank" rel="noopener"` attributes
- QR code

## Installation

1. Clone this repository

    ```bash
    $ git clone https://github.com/michimani/genlink.git
    ```

2. Install Go modules

    ```bash
    $ cd genlink
    $ go install
    ```

3. Build

    ```bash
    $ go build
    ```

## Usage

```bash
$ ./genlink -h
Usage:
  genlink [flags] [values]
Flags:
        -u (required)  URL
        -t             Type of link to output
            md:        Markdown (default)
            html:      HTML a tag
            html-bl:   HTML a tag with 'target="_blank"'
            qr:        QR code image
        -o             Absolute path to directory to output QR code
                       Use this flag in combination with '-t qr'
                       Default is current directory

Author:
  michimani <michimani210@gmail.com>
```

### Example

#### Generate link as Markdown format.

```bash
$ ./genlink -u https://github.com/michimani/genlink -t md
Target URL    :  https://github.com/michimani/genlink
Generate Type :  md


[GitHub - michimani/genlink: This is a command line tool that generates links to the specified URL in various formats.](https://github.com/michimani/genlink)
```

#### Generate link as QR code.

```bash
$ ./genlink -u https://github.com/michimani/genlink -t qr -o /path/to/out/dir
Target URL    :  https://github.com/michimani/genlink
Generate Type :  qr


QR code has been created.
/path/to/out/dir/8b448f6e8e9c09185f6f6ca8d143397b85586b8154cb9233cdc62cf2cecd90c6.png
```

![d11209189b7a094eb4b1c08b6ea73b8b8a6feba7bd6154efc2effdc87120c7b7](https://user-images.githubusercontent.com/9986092/106961267-acf2cc00-6780-11eb-9d83-9a9d4664a476.png)
