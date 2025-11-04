# Single Executable Distribution

## ✅ Single .exe File Solution

The application now embeds **everything** into a single executable file:
- ✅ r6-dissect.exe (embedded)
- ✅ All map images (embedded)
- ✅ No external files needed!

## How It Works

When the application starts:
1. Extracts `r6-dissect.exe` to a temporary directory
2. Extracts all map images to a temporary directory
3. Uses these extracted files during runtime
4. Temporary files are cleaned up automatically by Windows

## Building the Single Executable

### Prerequisites

1. **r6-dissect.exe** must be in the project root directory
2. **r6-maps-images/** folder must exist in the project root

### Build Steps

```powershell
# 1. Ensure r6-dissect.exe is in the root directory
# (Download from GitHub releases if needed)

# 2. Ensure r6-maps-images folder exists
# (Should already be present)

# 3. Build the single executable
$env:CGO_ENABLED=1
go build -ldflags="-s -w" -o r6-dissect-portable.exe
```

### What Gets Embedded

- `r6-dissect.exe` - Entire CLI tool embedded as binary data
- `r6-maps-images/*` - All map image files embedded

### Final Executable Size

- **Base GUI**: ~15-20 MB
- **Embedded r6-dissect.exe**: ~5-10 MB
- **Embedded map images**: ~5-10 MB
- **Total**: ~25-40 MB single executable

## Distribution

### For Users

Users just need to:
1. **Download** `r6-dissect-portable.exe`
2. **Run** it (double-click)
3. **Done!** No setup, no folders, no configuration

### For Developers

When building:
1. Place `r6-dissect.exe` in project root
2. Ensure `r6-maps-images/` folder exists
3. Build with `go build`
4. The resulting `.exe` contains everything!

## Benefits

✅ **Single file** - Just one .exe to download and run
✅ **No setup** - No installation or configuration needed
✅ **Portable** - Works from anywhere (USB drive, desktop, etc.)
✅ **Self-contained** - No external dependencies required
✅ **User-friendly** - No technical knowledge needed

## Notes

- Temporary files are extracted to Windows temp directory
- Match data (matches/ folder) is still stored relative to executable
- Embedded files are automatically extracted on first run
- No manual extraction or setup required

## Troubleshooting

### "r6-dissect.exe not found" during build
- **Solution**: Download `r6-dissect.exe` from GitHub releases and place in project root

### Build fails with embed errors
- **Solution**: Ensure `r6-dissect.exe` and `r6-maps-images/` folder exist in project root before building

### Executable is very large
- **Normal**: The executable contains everything embedded. ~25-40 MB is expected.

### Temp directory permissions
- **Solution**: Run as Administrator if temp directory creation fails (rare)

