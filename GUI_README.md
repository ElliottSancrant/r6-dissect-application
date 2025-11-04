# R6 Dissect Portable - GUI Version

This is a portable GUI version of r6-dissect that allows non-developers to easily analyze Rainbow Six Siege match replays.

## Features

- **User-friendly GUI**: Simple interface for analyzing match replays
- **Automatic folder detection**: Automatically finds your MatchReplay folder (Steam or Ubisoft)
- **Match history**: View previously analyzed matches with map images
- **Built-in spreadsheet viewer**: View Excel files directly in the application
- **Smart naming**: Automatically names files based on map, score, and date

## Building

### GUI Version (Default)
```bash
go build -o r6-dissect-portable.exe
```

### CLI Version
```bash
go build -tags cli -o r6-dissect-cli.exe
```

## Requirements

- `r6-dissect.exe` must be in the same directory as the portable executable
- `r6-maps-images` folder must be in the same directory (for map images)

## Usage

1. Run `r6-dissect-portable.exe`
2. Click "Analyze New Game"
3. Select your launcher (Steam or Ubisoft)
4. Select a match replay folder
5. The application will automatically:
   - Extract map name, score, and date
   - Generate an Excel file with statistics
   - Save it with a descriptive filename
   - Display it in the built-in viewer

## File Structure

```
dissect-portable/
├── r6-dissect-portable.exe  (GUI version)
├── r6-dissect.exe           (CLI tool - must be present)
├── r6-maps-images/          (Map images folder)
├── matches/                 (Generated Excel files)
└── matches.json             (Match metadata)
```

## Notes

- Matches are stored in the `matches/` directory
- Match metadata is stored in `matches.json`
- Duplicate matches (same map, score, date) are automatically numbered: `(1)`, `(2)`, etc.

