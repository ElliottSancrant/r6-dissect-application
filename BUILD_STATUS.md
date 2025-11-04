# Quick Build Guide

## ⚠️ C Compiler Required

Building locally requires a C compiler (gcc) because Fyne uses CGO. Your system doesn't have one installed.

## ✅ Solution: Use GitHub Actions (Recommended)

The easiest way to build is using GitHub Actions - it will build automatically!

### Option 1: Trigger Build Manually

1. Push your code to GitHub
2. Go to **Actions** tab
3. Select **"Build Portable Executable"** workflow
4. Click **"Run workflow"**
5. Wait for build to complete (~5-10 minutes)
6. Download `r6-dissect-portable.exe` from **Artifacts**

### Option 2: Create Release Tag

1. Create a git tag: `git tag v1.0`
2. Push tag: `git push origin v1.0`
3. GitHub Actions will build and create a release automatically
4. Download from **Releases** page

## Option 2: Install C Compiler Locally

If you want to build locally:

### Quick Install: TDM-GCC

1. Download: https://jmeubank.github.io/tdm-gcc/
2. Install with default settings
3. Restart terminal
4. Run: `.\build.bat`

## Current Status

✅ Code is ready to build
✅ GitHub Actions workflow created
✅ Build script created (`build.bat`)
❌ C compiler not installed locally

## Next Steps

**Recommended**: Use GitHub Actions to build (no local setup needed)

1. Commit and push your code
2. Trigger GitHub Actions build
3. Download the built executable
4. Distribute the single `.exe` file!

The built executable will be ~25-40 MB and contain everything embedded - users just download and run!

