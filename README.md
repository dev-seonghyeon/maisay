# maisay




<img width="1920" height="1080" alt="Image" src="https://github.com/user-attachments/assets/f56ab611-6009-4f30-93c0-1fa1c8644495" />




# maisay

> A terminal companion inspired by [cowsay](https://en.wikipedia.org/wiki/Cowsay) and [ponysay](https://github.com/erkin/ponysay) — but starring Mai. Because every terminal deserves a little senpai. 🐰

## Table of Contents

- [About](#about)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)
- [Credits](#credits)

## About

`maisay` is a small CLI tool inspired by [cowsay](https://en.wikipedia.org/wiki/Cowsay) and [ponysay](https://github.com/erkin/ponysay). It was born from a simple wish: to meet Sakurajima Mai in the terminal.

If you've ever wanted your shell to feel a little more senpai-approved, this one's for you.

## Features

- 🗨️ **`echo` command** — Mai prints whatever you give her, rendered alongside her ASCII portrait.




## Tech Stack

- **Language:** [Go](https://go.dev/)
- **CLI framework:** [Cobra](https://github.com/spf13/cobra)
- **Character art pipeline:** ChatGPT image generation + online image-to-ASCII art tools

## Installation

Build from source:

```bash
go build -o maisay maisay.go
```

This produces a single `maisay` binary you can drop anywhere on your `PATH`.


## Usage

```bash
./maisay echo "Senpai, did you forget about me?"
```



## Project Structure

```
maisay
├──.gitignore
├──.git
├── go.mod
├── go.sum
├── mai14.txt
├── mai.go
├── maisay
└── mai.txt
```


## Roadmap


## Contributing

Contributions are very welcome! To contribute:

1. **Open an issue** describing the bug or feature you'd like to work on.
2. **Fork** the repository and create a feature branch.
3. **Submit a pull request** referencing the issue.

Please keep changes focused and include a brief description of what and why.

## License

Released under the [GNU General Public License v3.0](./LICENSE). See the `LICENSE` file for the full text.

> ⚠️ **Note on character IP:** Mai is a character from *Seishun Buta Yarou wa Bunny Girl Senpai no Yume wo Minai* (Rascal Does Not Dream of Bunny Girl Senpai) and is owned by its respective rights holders. This project is a non-commercial fan work — the **code** is GPL3.0-licensed, but the **character likeness** is not. Please use respectfully and avoid commercial redistribution.

## Credits

- Inspired by [cowsay](https://en.wikipedia.org/wiki/Cowsay) and [ponysay](https://github.com/erkin/ponysay)
- Character: **Mai** from *Rascal Does Not Dream of Bunny Girl Senpai*
- Contributors: _(add as the project grows)_

Thank you for stopping by! 🙇
