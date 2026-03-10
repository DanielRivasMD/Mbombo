# mbombo, forge your own path

[![License](https://img.shields.io/badge/license-GPLv3-blue.svg)](LICENSE)

## Overview

Forge unified products from templates with intelligent text replacement

    ┌─────────────────────────────────────────────────────────────────────────────┐
    │                                 mbombo                                      │
    │                   Concatenate files with replacements                       │
    └─────────────────────────────────────────────────────────────────────────────┘
                                        │
                        ┌───────────────┴───────────────┐
                        │               │               │
                        ▼               ▼               ▼
    ┌───────────────────────────┐ ┌───────────────────────────┐ ┌───────────────────────────┐
    │         INPUT             │ │        RULES              │ │         OUTPUT            │
    ├───────────────────────────┤ ├───────────────────────────┤ ├───────────────────────────┤
    │ • Source files (--files)  │ │ • Token replacements      │ │ • Single output file      │
    │ • Input directory (--in)  │ │   ("old=new")             │ │   (--out)                 │
    │                           │ │                           │ │                           │
    │                           │ │ • Line replacements       │ │ • Can overwrite source    │
    │                           │ │   ("old=new:line")        │ │   if single file          │
    │                           │ │                           │ │                           │
    │                           │ │ • Multiple pairs allowed  │ │ • Preserves directory     │
    │                           │ │   (comma-separated)       │ │   structure               │
    └───────────────────────────┘ └───────────────────────────┘ └───────────────────────────┘
                                        │
                        ┌───────────────┴───────────────┐
                        │                               │
                        ▼                               ▼
    ┌───────────────────────────┐                 ┌───────────────────────────┐
    │      TOKEN MODE           │                 │       LINE MODE           │
    ├───────────────────────────┤                 ├───────────────────────────┤
    │ Replace occurrences       │                 │ Replace entire lines      │
    │ within lines              │                 │ containing text           │
    │                           │                 │                           │
    │ Example:                  │                 │ Example:                  │
    │ {{YEAR}} → 2026           │                 │ console.log() → (empty)   │
    └───────────────────────────┘                 └───────────────────────────┘

    ┌─────────────────────────────────────────────────────────────────────────────┐
    │                         COMMAND FLOW                                        │
    ├─────────────────────────────────────────────────────────────────────────────┤
    │                                                                             │
    │  ┌──────┐    ┌──────────┐    ┌───────────┐    ┌──────────┐    ┌────────┐    │
    │  │ USER │───▶│  PARSE   │───▶│ NORMALIZE │───▶│ PROCESS  │───▶│ OUTPUT │    │
    │  │      │    │  FLAGS   │    │  PATHS    │    │  FILES   │    │        │    │
    │  └──────┘    └──────────┘    └───────────┘    └──────────┘    └────────┘    │
    │                   │               │                │               │        │
    │                   ▼               ▼                ▼               ▼        │
    │              --files         extract dir/      read each       write to     │
    │              --out           base names          file          --out        │
    │              --replace                             │                        │
    │                                                    ▼                        │
    │                                               apply rules                   │
    │                                                                             │
    └─────────────────────────────────────────────────────────────────────────────┘

    ┌─────────────────────────────────────────────────────────────────────────────┐
    │                          USE CASES                                          │
    ├─────────────────────────────────────────────────────────────────────────────┤
    │                                                                             │
    │  ╔═══════════════════════════════════════════════════════════════════════╗  │
    │  ║  Template Processing                                                  ║  │
    │  ║  mbombo forge --files template.tmpl --out config.txt \                ║  │
    │  ║      --replace VERSION=1.0.0,API_URL=https://api.example.com          ║  │
    │  ╚═══════════════════════════════════════════════════════════════════════╝  │
    │                                                                             │
    │  ╔═══════════════════════════════════════════════════════════════════════╗  │
    │  ║  File Concatenation                                                   ║  │
    │  ║  mbombo forge --files header.html body.html footer.html \             ║  │
    │  ║      --out website/page.html                                          ║  │
    │  ╚═══════════════════════════════════════════════════════════════════════╝  │
    │                                                                             │
    │  ╔═══════════════════════════════════════════════════════════════════════╗  │
    │  ║  Code Cleanup                                                         ║  │
    │  ║  mbombo forge --files script.js --out script.min.js \                 ║  │
    │  ║      --replace console.log=:line,DEBUG=false                          ║  │
    │  ╚═══════════════════════════════════════════════════════════════════════╝  │
    │                                                                             │
    │  ╔═══════════════════════════════════════════════════════════════════════╗  │
    │  ║  Version Bumping                                                      ║  │
    │  ║  mbombo forge --files README.md --out README.md \                     ║  │
    │  ║      --replace v0.0.0=v1.2.3                                          ║  │
    │  ╚═══════════════════════════════════════════════════════════════════════╝  │
    │                                                                             │
    └─────────────────────────────────────────────────────────────────────────────┘

## Installation

### Language-Specific

    Go:  go install github.com/DanielRivasMD/Mbombo@latest

## License

Copyright (c) 2024

See the [LICENSE](LICENSE) file for license details
