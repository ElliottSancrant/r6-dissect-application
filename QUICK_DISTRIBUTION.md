# Creating a Distribution Package

## Automated Method

### Windows (PowerShell)

```powershell
.\create_package.ps1
```

This will create a `r6-dissect-portable-v1.0` folder with all necessary files.

### Linux/Mac (Bash)

```bash
chmod +x create_package.sh
./create_package.sh
```

## Manual Method

### Step 1: Build the Executable

```powershell
# Install dependencies
go mod tidy

# Build optimized executable
$env:CGO_ENABLED=1
go build -ldflags="-s -w" -o r6-dissect-portable.exe
```

### Step 2: Download r6-dissect.exe

Download from: https://github.com/redraskal/r6-dissect/releases

### Step 3: Create Package Structure

```
r6-dissect-portable-v1.0/
├── r6-dissect-portable.exe
├── r6-dissect.exe
├── r6-maps-images/
│   └── (all map images)
├── setup.bat
├── USER_GUIDE.md
└── README.md
```

### Step 4: Create ZIP Archive

```powershell
Compress-Archive -Path r6-dissect-portable-v1.0\* -DestinationPath r6-dissect-portable-v1.0.zip
```

## Distribution Checklist

- [ ] GUI executable built (`r6-dissect-portable.exe`)
- [ ] r6-dissect.exe included (latest version)
- [ ] r6-maps-images folder included (all images)
- [ ] Documentation included (USER_GUIDE.md, README.md)
- [ ] Package tested on clean Windows system
- [ ] ZIP archive created

## Testing

Before distributing:

1. Extract package to a new folder
2. Run `setup.bat` (if included)
3. Run `r6-dissect-portable.exe`
4. Test analyzing a match
5. Test viewing previous matches
6. Verify all features work

## File Sizes

- `r6-dissect-portable.exe`: ~15-20 MB
- `r6-dissect.exe`: ~5-10 MB
- `r6-maps-images/`: ~5-10 MB
- **Total package**: ~25-40 MB

## Ready to Distribute!

Once the ZIP is created, users can:
1. Download the ZIP
2. Extract anywhere
3. Run `r6-dissect-portable.exe`
4. Start analyzing matches!

