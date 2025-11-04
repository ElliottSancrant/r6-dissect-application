@echo off
echo R6 Dissect Portable - Setup Script
echo ===================================
echo.

REM Check if r6-dissect.exe exists
if not exist "r6-dissect.exe" (
    echo ERROR: r6-dissect.exe not found!
    echo Please ensure r6-dissect.exe is in the same directory as this script.
    echo.
    pause
    exit /b 1
)

REM Check if r6-maps-images folder exists
if not exist "r6-maps-images" (
    echo WARNING: r6-maps-images folder not found!
    echo The application will work, but map images won't display.
    echo.
)

REM Create matches directory if it doesn't exist
if not exist "matches" (
    mkdir matches
    echo Created matches directory.
)

REM Check if main executable exists
if not exist "r6-dissect-portable.exe" (
    echo ERROR: r6-dissect-portable.exe not found!
    echo Please ensure the executable is in the same directory as this script.
    echo.
    pause
    exit /b 1
)

echo Setup complete!
echo.
echo You can now run r6-dissect-portable.exe
echo.
pause

