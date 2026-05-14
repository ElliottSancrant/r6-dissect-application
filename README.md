# R6 Dissect Portable

A user-friendly GUI application for analyzing Rainbow Six Siege match replays.

## Quick Start

1. **Download** the latest release
2. **Extract** all files to a folder
3. **Run** `r6-dissect-portable.exe`
4. **Analyze** your matches!

## Features

-  **Easy Match Analysis** - Analyze R6 Siege matches with a simple GUI
-  **Auto-Detection** - Automatically finds your MatchReplay folder
- **Match History** - View all previously analyzed matches
-  **Built-in Viewer** - View Excel statistics without opening Excel
-  **Smart Naming** - Files automatically named by map, score, and date
- **Map Images** - Visual identification of matches
-  **Portable** - No installation required, just extract and run!

##  Requirements

- Windows 10/11 (64-bit)
- Rainbow Six Siege (Steam or Ubisoft)
- All files included in the package

##  Documentation

- **[USER_GUIDE.md](USER_GUIDE.md)** - Complete user guide
- **[BUILD_INSTRUCTIONS.md](BUILD_INSTRUCTIONS.md)** - Build from source
- **[DISTRIBUTION.md](DISTRIBUTION.md)** - Create distribution package

## How to Use

1. Click **"Analyze New Game"**
2. Select your launcher (Steam or Ubisoft)
3. Choose a match replay folder
4. View your statistics!

## Package Contents

```
r6-dissect-portable/
├── r6-dissect-portable.exe    ← Main application
├── r6-dissect.exe             ← Required CLI tool
├── r6-maps-images/            ← Map images
├── setup.bat                   ← Optional setup script
├── USER_GUIDE.md              ← User documentation
└── matches/                   ← Created automatically
```

## Troubleshooting

See [USER_GUIDE.md](USER_GUIDE.md) for common issues and solutions.

## License

Same license as r6-dissect.

## Credits

Built on top of [r6-dissect](https://github.com/redraskal/r6-dissect) by redraskal.
