# R6 Dissect Portable - User Guide

## Quick Start

1. **Download** the R6 Dissect Portable package
2. **Extract** all files to a folder (e.g., `C:\R6DissectPortable\`)
3. **Run** `r6-dissect-portable.exe`
4. **Select** your launcher (Steam or Ubisoft)
5. **Choose** a match replay folder to analyze
6. **View** your match statistics!

## What's Included

```
r6-dissect-portable/
├── r6-dissect-portable.exe    ← Main application (double-click to run)
├── r6-dissect.exe              ← Required: CLI tool (must be present)
├── r6-maps-images/            ← Required: Map images folder
│   ├── ModernizedMap_Nighthaven_keyart.png
│   ├── ModernizedMap_Consulate_keyart.png
│   └── ... (other map images)
├── README.md                   ← This file
└── matches/                   ← Created automatically (stores your analyzed matches)
```

## Requirements

- **Windows 10/11** (64-bit)
- **Rainbow Six Siege** installed via Steam or Ubisoft
- **No additional software needed** - everything is included!

## How to Use

### Analyzing a New Match

1. Click **"Analyze New Game"**
2. Select **Steam** or **Ubisoft** (depending on where you installed Rainbow Six Siege)
3. The app will automatically find your MatchReplay folder
4. Select the match folder you want to analyze (e.g., `Match-2024-01-15_14-30-00-123`)
5. Wait for the analysis to complete
6. The Excel file will open automatically in the built-in viewer

### Viewing Previous Matches

1. Click **"View Previously Scanned Games"**
2. Select any match from the list
3. View the detailed statistics in the built-in spreadsheet viewer
4. Switch between different sheets (Match overview, Round 1, Round 2, etc.)

## File Locations

### Where are my analyzed matches stored?

All analyzed matches are saved in the `matches/` folder, automatically created next to the executable.

Files are named like: `MapName_Score_Date.xlsx`
- Example: `Nighthaven_Labs_7-1_10-25-2024.xlsx`

### Match metadata

Match information is stored in `matches.json` (created automatically).

## Troubleshooting

### "r6-dissect.exe not found"
- **Solution**: Make sure `r6-dissect.exe` is in the same folder as `r6-dissect-portable.exe`
- If missing, download it from the r6-dissect releases page

### "MatchReplay folder not found"
- **Solution**: The app will ask you to browse manually
- Typical locations:
  - Steam: `C:\Program Files (x86)\Steam\steamapps\common\Tom Clancy's Rainbow Six Siege\MatchReplay`
  - Ubisoft: `C:\Program Files (x86)\Ubisoft\Ubisoft Game Launcher\games\Tom Clancy's Rainbow Six Siege\MatchReplay`
- If installed on a different drive, browse to it manually

### "Map image not found"
- **Solution**: Make sure the `r6-maps-images` folder is present
- The app will still work, but match images won't display

### Application won't start
- **Solution**: Make sure you're running Windows 10/11 (64-bit)
- Try running as Administrator
- Check Windows Defender/Antivirus isn't blocking it

### Analysis fails
- **Solution**: Make sure you selected a valid match folder (contains .rec files)
- Check that r6-dissect.exe is working by running it from command line:
  ```
  r6-dissect.exe "path\to\match\folder" -o test.xlsx
  ```

## Features

- ✅ **Automatic folder detection** - Finds your MatchReplay folder automatically
- ✅ **Smart naming** - Files automatically named by map, score, and date
- ✅ **Duplicate handling** - Automatically numbers duplicate matches
- ✅ **Built-in viewer** - View Excel files without opening Excel
- ✅ **Match history** - Keep track of all analyzed matches
- ✅ **Portable** - No installation needed, just extract and run!

## Support

If you encounter issues:
1. Check this README for common solutions
2. Ensure all files are in the correct locations
3. Make sure you have the latest version of r6-dissect.exe

## License

This portable version uses the same license as r6-dissect.

---

**Enjoy analyzing your Rainbow Six Siege matches!**

