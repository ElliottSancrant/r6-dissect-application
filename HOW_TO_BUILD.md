# Building R6 Dissect Portable

## Option 1: Automatic Build (Recommended)

### Using GitHub Actions

1. Push your code to GitHub
2. Create a new release tag (e.g., `v1.0`)
3. GitHub Actions will automatically build the executable
4. Download from the Releases page

### Manual GitHub Actions Trigger

1. Go to Actions tab in your repository
2. Select "Build Portable Executable"
3. Click "Run workflow"
4. Download the built executable from artifacts

## Option 2: Local Build (Requires C Compiler)

### Step 1: Install C Compiler

**Option A: TDM-GCC (Easiest)**
1. Download from: https://jmeubank.github.io/tdm-gcc/
2. Install with default settings
3. Restart your terminal

**Option B: MinGW-w64**
1. Download from: https://www.mingw-w64.org/downloads/
2. Install and add to PATH

**Option C: MSYS2**
```bash
# Install MSYS2 from https://www.msys2.org/
# Then install gcc:
pacman -S mingw-w64-x86_64-gcc
```

### Step 2: Build

```powershell
# Run the build script
.\build.bat

# Or manually:
$env:CGO_ENABLED=1
go build -ldflags="-s -w" -o r6-dissect-portable.exe
```

## Option 3: Pre-built Binary

If you have access to a machine with gcc already installed, you can build there and distribute the resulting `.exe` file.

## Verification

After building, verify:
- ✅ File exists: `r6-dissect-portable.exe`
- ✅ File size: ~25-40 MB
- ✅ Can run: Double-click and GUI opens
- ✅ No external files needed: Works standalone

## Distribution

Once built, the single `.exe` file can be distributed:
- Upload to GitHub Releases
- Upload to file sharing service
- Share directly with users

Users just need to download and run - no setup required!

