# R6 Dissect Portable - Build Instructions

## Overview

This is a portable GUI version of r6-dissect that provides a user-friendly interface for analyzing Rainbow Six Siege match replays.

## Build Requirements

### For Windows GUI Build:
1. **Go 1.23+** - https://golang.org/dl/
2. **C Compiler** (for CGO) - Required for Fyne GUI:
   - **Option 1**: Install [TDM-GCC](https://jmeubank.github.io/tdm-gcc/) or [MinGW-w64](https://www.mingw-w64.org/)
   - **Option 2**: Install [MSYS2](https://www.msys2.org/) and then install gcc:
     ```bash
     pacman -S mingw-w64-x86_64-gcc
     ```

### For CLI Build (no CGO required):
```bash
go build -tags cli -o r6-dissect-cli.exe
```

## Building the GUI Version

1. **Install dependencies:**
   ```bash
   go mod tidy
   ```

2. **Build the GUI executable:**
   ```bash
   # Windows
   set CGO_ENABLED=1
   go build -o r6-dissect-portable.exe
   
   # Or using PowerShell
   $env:CGO_ENABLED=1
   go build -o r6-dissect-portable.exe
   ```

## Portable Package Structure

When distributing the portable version, ensure the following structure:

```
r6-dissect-portable/
├── r6-dissect-portable.exe  (GUI executable)
├── r6-dissect.exe           (CLI tool - REQUIRED)
├── r6-maps-images/         (Map images folder - REQUIRED)
│   ├── ModernizedMap_Nighthaven_keyart.png
│   ├── ModernizedMap_Consulate_keyart.png
│   ├── ModernizedMap_Lair_keyart.png
│   └── ... (other map images)
├── matches/                 (Created automatically - stores Excel files)
└── matches.json             (Created automatically - stores metadata)
```

## Usage

1. Ensure `r6-dissect.exe` and `r6-maps-images/` folder are in the same directory as the portable executable
2. Run `r6-dissect-portable.exe`
3. Follow the GUI prompts to analyze matches

## Features

- ✅ Automatic Steam/Ubisoft folder detection
- ✅ Multi-drive search for MatchReplay folders
- ✅ Smart filename generation (Map_Score_Date.xlsx)
- ✅ Duplicate detection and numbering
- ✅ Map image association
- ✅ Built-in Excel viewer
- ✅ Match history tracking
- ✅ Portable (all paths relative to executable)

## Troubleshooting

### Build Errors:
- **"C compiler not found"**: Install a C compiler (TDM-GCC, MinGW-w64, or MSYS2)
- **"missing go.sum entry"**: Run `go mod tidy`
- **"package not found"**: Run `go get fyne.io/fyne/v2` and `go mod tidy`

### Runtime Errors:
- **"r6-dissect.exe not found"**: Ensure r6-dissect.exe is in the same directory as the portable executable
- **"r6-maps-images folder not found"**: Ensure the r6-maps-images folder exists in the same directory
- **MatchReplay folder not found**: The app will prompt you to select manually if auto-detection fails

## Notes

- The GUI version is built by default (use `-tags cli` for CLI version)
- All data files (matches/, matches.json) are stored relative to the executable directory
- The application automatically handles duplicate matches with numbering: `(1)`, `(2)`, etc.

