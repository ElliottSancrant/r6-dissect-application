# R6 Dissect Portable - Distribution Package

## Creating a Distribution Package

This document explains how to create a ready-to-distribute package of R6 Dissect Portable.

## Package Structure

```
r6-dissect-portable-v1.0/
├── r6-dissect-portable.exe    (GUI executable - built from main.go)
├── r6-dissect.exe             (CLI tool - download from releases)
├── r6-maps-images/            (Map images folder)
│   ├── ModernizedMap_Nighthaven_keyart.png
│   ├── ModernizedMap_Consulate_keyart.png
│   ├── ModernizedMap_Lair_keyart.png
│   ├── r6-maps-coastline.png
│   ├── r6-maps-favela__1_.png
│   ├── r6-maps-fortress.png
│   ├── r6-maps-hereford.png
│   ├── r6-maps-house.png
│   ├── r6-maps-kanal.png
│   ├── r6-maps-oregon.png
│   ├── r6-maps-outback.png
│   ├── r6-maps-plane.png
│   ├── r6-maps-skyscraper.png
│   ├── r6-maps-tower.png
│   ├── r6-maps-villa.png
│   ├── r6-maps-yacht.png
│   ├── R6S_Maps_Bank_EXT.png
│   ├── R6S_Maps_Border_EXT.png
│   ├── R6S_Maps_Chalet_EXT.png
│   ├── R6S_Maps_ClubHouse_EXT.png
│   ├── R6S_Maps_RussianCafe_EXT.png
│   ├── r6s_maps_emeraldplains__1_.png
│   ├── rainbow6_maps_theme-park_thumbnail.png
│   ├── StadiumA_keyart.png
│   └── stadiumB_keyart.png
├── setup.bat                  (Optional setup script)
├── USER_GUIDE.md             (User documentation)
├── README.md                  (Main readme)
└── matches/                   (Created automatically on first run)
```

## Steps to Create Distribution Package

### 1. Build the GUI Executable

```bash
# Ensure dependencies are installed
go mod tidy

# Build with CGO enabled
set CGO_ENABLED=1
go build -o r6-dissect-portable.exe

# Or for release build (optimized, no debug info)
set CGO_ENABLED=1
go build -ldflags="-s -w" -o r6-dissect-portable.exe
```

### 2. Get r6-dissect.exe

Download the latest `r6-dissect.exe` from:
- https://github.com/redraskal/r6-dissect/releases

Place it in the same directory as `r6-dissect-portable.exe`.

### 3. Include Map Images

Copy the entire `r6-maps-images` folder to the distribution package.

### 4. Create Package Directory

```bash
mkdir r6-dissect-portable-v1.0
copy r6-dissect-portable.exe r6-dissect-portable-v1.0\
copy r6-dissect.exe r6-dissect-portable-v1.0\
xcopy /E /I r6-maps-images r6-dissect-portable-v1.0\r6-maps-images
copy USER_GUIDE.md r6-dissect-portable-v1.0\
copy setup.bat r6-dissect-portable-v1.0\
```

### 5. Create ZIP Archive

```bash
# Using PowerShell
Compress-Archive -Path r6-dissect-portable-v1.0\* -DestinationPath r6-dissect-portable-v1.0.zip

# Or using 7-Zip (if installed)
7z a -tzip r6-dissect-portable-v1.0.zip r6-dissect-portable-v1.0\*
```

## Distribution Checklist

- [ ] GUI executable built and tested (`r6-dissect-portable.exe`)
- [ ] r6-dissect.exe included (latest version)
- [ ] r6-maps-images folder included (all images)
- [ ] USER_GUIDE.md included
- [ ] setup.bat included (optional)
- [ ] README.md included (optional)
- [ ] Package tested on clean Windows system
- [ ] ZIP archive created

## Alternative: Include Build Script

For users who want to build from source, include:
- `go.mod`
- `go.sum`
- Source files (`main.go`, `dissect/` folder, etc.)
- `BUILD_INSTRUCTIONS.md`

## Testing the Package

Before distributing:

1. **Extract to a new folder** (simulate fresh download)
2. **Run setup.bat** (if included)
3. **Run r6-dissect-portable.exe**
4. **Test analyzing a match**
5. **Test viewing previous matches**
6. **Verify all features work**

## Version Information

Update version in:
- Package folder name: `r6-dissect-portable-v1.0`
- USER_GUIDE.md (if version-specific info)
- README.md

## File Sizes (Approximate)

- `r6-dissect-portable.exe`: ~15-20 MB (depends on build flags)
- `r6-dissect.exe`: ~5-10 MB
- `r6-maps-images/`: ~5-10 MB (depends on image quality)
- **Total package**: ~25-40 MB

## Alternative Distribution Methods

### Option 1: Single ZIP Archive
- Easiest for users
- Extract anywhere and run

### Option 2: Installer (Advanced)
- Use NSIS, Inno Setup, or WiX
- Can create Start Menu shortcuts
- Can add to PATH
- More professional but requires installer creation

### Option 3: Portable App Format
- Follow PortableApps.com format
- Can be included in PortableApps.com launcher
- More structured but requires format compliance

## Recommended: Simple ZIP Distribution

For maximum compatibility and ease of use:
1. Create ZIP archive with all files
2. Users extract to any folder
3. Users run `r6-dissect-portable.exe`
4. That's it!

## Notes

- The application is fully portable - no installation needed
- All data is stored relative to the executable directory
- Users can move the entire folder anywhere
- No registry entries or system modifications required

