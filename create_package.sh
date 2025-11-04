#!/bin/bash
# R6 Dissect Portable - Package Creation Script
# This script creates a distribution-ready package

set -e

VERSION="1.0"
PACKAGE_NAME="r6-dissect-portable-v${VERSION}"
PACKAGE_DIR="${PACKAGE_NAME}"

echo "Creating R6 Dissect Portable distribution package..."
echo "Version: ${VERSION}"
echo ""

# Create package directory
if [ -d "${PACKAGE_DIR}" ]; then
    echo "Removing existing package directory..."
    rm -rf "${PACKAGE_DIR}"
fi

mkdir -p "${PACKAGE_DIR}"
mkdir -p "${PACKAGE_DIR}/r6-maps-images"
mkdir -p "${PACKAGE_DIR}/matches"

echo "Copying files..."

# Copy executable (if built)
if [ -f "r6-dissect-portable.exe" ]; then
    cp "r6-dissect-portable.exe" "${PACKAGE_DIR}/"
    echo "  ✓ Copied r6-dissect-portable.exe"
else
    echo "  ⚠ WARNING: r6-dissect-portable.exe not found (build it first!)"
fi

# Copy r6-dissect.exe (if present)
if [ -f "r6-dissect.exe" ]; then
    cp "r6-dissect.exe" "${PACKAGE_DIR}/"
    echo "  ✓ Copied r6-dissect.exe"
else
    echo "  ⚠ WARNING: r6-dissect.exe not found (download from releases!)"
fi

# Copy map images
if [ -d "r6-maps-images" ]; then
    cp -r r6-maps-images/* "${PACKAGE_DIR}/r6-maps-images/"
    echo "  ✓ Copied r6-maps-images/"
else
    echo "  ⚠ WARNING: r6-maps-images folder not found"
fi

# Copy documentation
cp USER_GUIDE.md "${PACKAGE_DIR}/" 2>/dev/null && echo "  ✓ Copied USER_GUIDE.md" || echo "  ⚠ USER_GUIDE.md not found"
cp README.md "${PACKAGE_DIR}/" 2>/dev/null && echo "  ✓ Copied README.md" || echo "  ⚠ README.md not found"
cp setup.bat "${PACKAGE_DIR}/" 2>/dev/null && echo "  ✓ Copied setup.bat" || echo "  ⚠ setup.bat not found"

# Create matches directory placeholder
echo "# Matches will be stored here" > "${PACKAGE_DIR}/matches/README.txt"
echo "  ✓ Created matches directory"

echo ""
echo "Package created: ${PACKAGE_DIR}/"
echo ""
echo "Next steps:"
echo "1. Ensure r6-dissect-portable.exe is built"
echo "2. Ensure r6-dissect.exe is included"
echo "3. Test the package"
echo "4. Create ZIP archive:"
echo "   zip -r ${PACKAGE_NAME}.zip ${PACKAGE_DIR}/"
echo ""

