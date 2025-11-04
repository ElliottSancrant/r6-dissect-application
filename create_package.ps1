# R6 Dissect Portable - Package Creation Script (PowerShell)
# This script creates a distribution-ready package

$VERSION = "1.0"
$PACKAGE_NAME = "r6-dissect-portable-v$VERSION"
$PACKAGE_DIR = $PACKAGE_NAME

Write-Host "Creating R6 Dissect Portable distribution package..."
Write-Host "Version: $VERSION"
Write-Host ""

# Remove existing package directory
if (Test-Path $PACKAGE_DIR) {
    Write-Host "Removing existing package directory..."
    Remove-Item -Recurse -Force $PACKAGE_DIR
}

# Create package directory
New-Item -ItemType Directory -Path $PACKAGE_DIR | Out-Null
New-Item -ItemType Directory -Path "$PACKAGE_DIR\r6-maps-images" | Out-Null
New-Item -ItemType Directory -Path "$PACKAGE_DIR\matches" | Out-Null

Write-Host "Copying files..."

# Copy executable (if built)
if (Test-Path "r6-dissect-portable.exe") {
    Copy-Item "r6-dissect-portable.exe" "$PACKAGE_DIR\"
    Write-Host "  ✓ Copied r6-dissect-portable.exe"
} else {
    Write-Host "  ⚠ WARNING: r6-dissect-portable.exe not found (build it first!)"
}

# Copy r6-dissect.exe (if present)
if (Test-Path "r6-dissect.exe") {
    Copy-Item "r6-dissect.exe" "$PACKAGE_DIR\"
    Write-Host "  ✓ Copied r6-dissect.exe"
} else {
    Write-Host "  ⚠ WARNING: r6-dissect.exe not found (download from releases!)"
}

# Copy map images
if (Test-Path "r6-maps-images") {
    Copy-Item -Recurse "r6-maps-images\*" "$PACKAGE_DIR\r6-maps-images\"
    Write-Host "  ✓ Copied r6-maps-images/"
} else {
    Write-Host "  ⚠ WARNING: r6-maps-images folder not found"
}

# Copy documentation
if (Test-Path "USER_GUIDE.md") {
    Copy-Item "USER_GUIDE.md" "$PACKAGE_DIR\"
    Write-Host "  ✓ Copied USER_GUIDE.md"
}
if (Test-Path "README.md") {
    Copy-Item "README.md" "$PACKAGE_DIR\"
    Write-Host "  ✓ Copied README.md"
}
if (Test-Path "setup.bat") {
    Copy-Item "setup.bat" "$PACKAGE_DIR\"
    Write-Host "  ✓ Copied setup.bat"
}

# Create matches directory placeholder
Set-Content -Path "$PACKAGE_DIR\matches\README.txt" -Value "# Matches will be stored here"
Write-Host "  ✓ Created matches directory"

Write-Host ""
Write-Host "Package created: $PACKAGE_DIR\"
Write-Host ""
Write-Host "Next steps:"
Write-Host "1. Ensure r6-dissect-portable.exe is built"
Write-Host "2. Ensure r6-dissect.exe is included"
Write-Host "3. Test the package"
Write-Host "4. Create ZIP archive:"
Write-Host "   Compress-Archive -Path $PACKAGE_DIR\* -DestinationPath $PACKAGE_NAME.zip"
Write-Host ""

