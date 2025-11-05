@echo off
setlocal EnableExtensions EnableDelayedExpansion
echo Building R6 Dissect Portable...
echo.

REM Check for required files
if not exist "r6-dissect.exe" (
    echo ERROR: r6-dissect.exe not found!
    echo Please download it from: https://github.com/redraskal/r6-dissect/releases
    echo Place it in the project root directory.
    pause
    exit /b 1
)

if not exist "r6-maps-images" (
    echo ERROR: r6-maps-images folder not found!
    pause
    exit /b 1
)

echo Checking for C compiler...
where gcc >nul 2>&1
IF ERRORLEVEL 1 (
    echo.
    echo ERROR: C compiler (gcc) not found!
    echo.
    echo Fyne GUI requires a C compiler to build.
    echo.
    echo Please install one of the following:
    echo   1. TDM-GCC: https://jmeubank.github.io/tdm-gcc/
    echo   2. MinGW-w64: https://www.mingw-w64.org/
    echo   3. MSYS2: https://www.msys2.org/ (then install: pacman -S mingw-w64-x86_64-gcc)
    echo.
    echo After installing, restart your terminal and try again.
    echo.
    pause
    exit /b 1
)

echo C compiler found!
echo.

echo Checking for Go toolchain...
where go >nul 2>&1
IF ERRORLEVEL 1 (
    echo ERROR: Go not found. Install with: choco install golang
    pause
    exit /b 1
)

echo Building executable...
echo.

set "CGO_ENABLED=1"
set "CC=gcc"
go build -ldflags="-s -w" -o r6-dissect-portable.exe

IF ERRORLEVEL 0 (
    echo.
    echo SUCCESS! Built: r6-dissect-portable.exe
    for %%A in (r6-dissect-portable.exe) do echo File size: %%~zA bytes
) ELSE (
    echo.
    echo BUILD FAILED!
    echo Check the error messages above.
)

pause
