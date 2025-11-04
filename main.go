//go:build !cli

package main

import (
	_ "embed"
	"embed"
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/redraskal/r6-dissect/dissect"
	"github.com/xuri/excelize/v2"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

//go:embed r6-dissect.exe
var embeddedR6DissectExe []byte

//go:embed r6-maps-images/*
var embeddedMapImages embed.FS

type MatchMetadata struct {
	MapName     string    `json:"mapName"`
	Score       string    `json:"score"`
	Date        time.Time `json:"date"`
	FilePath    string    `json:"filePath"`
	ImagePath   string    `json:"imagePath"`
	ReplayPath  string    `json:"replayPath"`
}

type MatchStore struct {
	Matches []MatchMetadata `json:"matches"`
}

const (
	matchesDir = "matches"
	storeFile  = "matches.json"
)

var (
	tempDir        string
	r6DissectPath  string
	mapImagesDir   string
)

func main() {
	// Extract embedded files to temp directory
	var err error
	tempDir, err = extractEmbeddedFiles()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error extracting embedded files: %v\n", err)
		os.Exit(1)
	}
	defer os.RemoveAll(tempDir) // Clean up on exit

	myApp := app.NewWithID("com.r6dissect.portable")
	myApp.Settings().SetTheme(&darkTheme{})
	
	myWindow := myApp.NewWindow("R6 Dissect Portable")
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.CenterOnScreen()

	// Load existing matches
	store := loadStore()
	
	// Create main content
	mainContent := createMainView(myWindow, store, myApp)
	myWindow.SetContent(mainContent)

	myWindow.ShowAndRun()
}

func extractEmbeddedFiles() (string, error) {
	// Create temp directory
	tempDir, err := os.MkdirTemp("", "r6-dissect-portable-*")
	if err != nil {
		return "", err
	}

	// Extract r6-dissect.exe
	r6DissectPath = filepath.Join(tempDir, "r6-dissect.exe")
	if err := os.WriteFile(r6DissectPath, embeddedR6DissectExe, 0755); err != nil {
		return "", err
	}

	// Extract map images
	mapImagesDir = filepath.Join(tempDir, "r6-maps-images")
	if err := os.MkdirAll(mapImagesDir, 0755); err != nil {
		return "", err
	}

	// Walk embedded filesystem and extract all images
	err = fs.WalkDir(embeddedMapImages, "r6-maps-images", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		
		// Read embedded file
		data, err := embeddedMapImages.ReadFile(path)
		if err != nil {
			return err
		}
		
		// Write to temp directory, preserving relative path
		relPath, err := filepath.Rel("r6-maps-images", path)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(mapImagesDir, relPath)
		
		// Create parent directories if needed
		if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
			return err
		}
		
		return os.WriteFile(targetPath, data, 0644)
	})
	if err != nil {
		return "", err
	}

	return tempDir, nil
}

func createMainView(window fyne.Window, store *MatchStore, app fyne.App) *container.Vertical {
	title := widget.NewLabel("R6 Dissect Portable")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	viewButton := widget.NewButton("View Previously Scanned Games", func() {
		window.SetContent(createMatchesView(window, store, app))
	})

	analyzeButton := widget.NewButton("Analyze New Game", func() {
		window.SetContent(createAnalyzeView(window, store, app))
	})

	content := container.NewVBox(
		title,
		widget.NewSeparator(),
		viewButton,
		analyzeButton,
	)

	return container.NewCenter(content)
}

func createMatchesView(window fyne.Window, store *MatchStore, app fyne.App) *container.Vertical {
	backButton := widget.NewButton("← Back", func() {
		window.SetContent(createMainView(window, store, app))
	})

	if len(store.Matches) == 0 {
		noMatches := widget.NewLabel("No matches scanned yet.")
		return container.NewVBox(backButton, container.NewCenter(noMatches))
	}

	// Create list of matches
	matchList := widget.NewList(
		func() int {
			return len(store.Matches)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(
				widget.NewIcon(nil),
				widget.NewLabel(""),
			)
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			match := store.Matches[id]
			cont := obj.(*container.HBox)
			label := cont.Objects[1].(*widget.Label)
			
			displayText := fmt.Sprintf("%s %s %s", match.MapName, match.Score, match.Date.Format("01-02-2006"))
			label.SetText(displayText)
			
			// Load image if available
			if match.ImagePath != "" && fileExists(match.ImagePath) {
				icon := cont.Objects[0].(*widget.Icon)
				if img, err := fyne.LoadResourceFromPath(match.ImagePath); err == nil {
					icon.SetResource(img)
				}
			}
		},
	)

	matchList.OnSelected = func(id widget.ListItemID) {
		match := store.Matches[id]
		window.SetContent(createSpreadsheetView(window, match, store, app))
	}

	title := widget.NewLabel("Previously Scanned Games")
	title.TextStyle = fyne.TextStyle{Bold: true}

	return container.NewBorder(
		container.NewVBox(title, widget.NewSeparator()),
		backButton,
		nil,
		nil,
		matchList,
	)
}

func createAnalyzeView(window fyne.Window, store *MatchStore, app fyne.App) *container.Vertical {
	backButton := widget.NewButton("← Back", func() {
		window.SetContent(createMainView(window, store, app))
	})

	statusLabel := widget.NewLabel("")
	
	launcherLabel := widget.NewLabel("Select your launcher:")
	steamButton := widget.NewButton("Steam", func() {
		analyzeMatch(window, store, app, "steam", statusLabel)
	})
	ubisoftButton := widget.NewButton("Ubisoft", func() {
		analyzeMatch(window, store, app, "ubisoft", statusLabel)
	})

	content := container.NewVBox(
		backButton,
		widget.NewSeparator(),
		launcherLabel,
		steamButton,
		ubisoftButton,
		statusLabel,
	)

	return container.NewCenter(content)
}

func analyzeMatch(window fyne.Window, store *MatchStore, app fyne.App, launcher string, statusLabel *widget.Label) {
	statusLabel.SetText("Finding MatchReplay folder...")
	
	replayPath := findMatchReplayFolder(launcher)
	if replayPath == "" {
		statusLabel.SetText("MatchReplay folder not found. Please select manually.")
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil || uri == nil {
				return
			}
			replayPath = uri.Path()
			processMatch(window, store, app, replayPath, statusLabel)
		}, window)
		return
	}

	// Open folder dialog
	statusLabel.SetText("Select a match replay folder...")
	uri, _ := storage.ListerForURI(storage.NewFileURI(replayPath))
	dialog.ShowFolderOpen(func(selectedURI fyne.ListableURI, err error) {
		if err != nil || selectedURI == nil {
			statusLabel.SetText("No folder selected.")
			return
		}
		processMatch(window, store, app, selectedURI.Path(), statusLabel)
	}, window)
}

func processMatch(window fyne.Window, store *MatchStore, app fyne.App, replayPath string, statusLabel *widget.Label) {
	statusLabel.SetText("Analyzing match...")
	
	// Read match to get metadata
	matchFile, err := os.Open(replayPath)
	if err != nil {
		statusLabel.SetText(fmt.Sprintf("Error: %v", err))
		return
	}
	defer matchFile.Close()

	matchReader, err := dissect.NewMatchReader(matchFile)
	if err != nil {
		statusLabel.SetText(fmt.Sprintf("Error reading match: %v", err))
		return
	}

	if err := matchReader.Read(); err != nil {
		statusLabel.SetText(fmt.Sprintf("Error reading match data: %v", err))
		return
	}

	// Get first round for metadata
	firstRound, err := matchReader.FirstRound()
	if err != nil {
		statusLabel.SetText(fmt.Sprintf("Error reading first round: %v", err))
		return
	}

	// Extract metadata
	mapName := firstRound.Header.Map.String()
	score0 := firstRound.Header.Teams[0].Score
	score1 := firstRound.Header.Teams[1].Score
	score := fmt.Sprintf("%d-%d", score0, score1)
	date := firstRound.Header.Timestamp

	// Get final scores from last round
	lastRound, err := matchReader.LastRound()
	if err == nil {
		score0 = lastRound.Header.Teams[0].Score
		score1 = lastRound.Header.Teams[1].Score
		score = fmt.Sprintf("%d-%d", score0, score1)
	}

	// Generate filename
	filename := generateFilename(mapName, score, date, store)
	matchesDirPath := getMatchesDir()
	outputPath := filepath.Join(matchesDirPath, filename)

	// Ensure matches directory exists
	os.MkdirAll(matchesDirPath, 0755)

	statusLabel.SetText("Generating Excel file...")

	// Execute r6-dissect.exe - find it relative to executable
	exePath := findR6DissectExe()
	if exePath == "" {
		statusLabel.SetText("Error: r6-dissect.exe not found in current directory")
		return
	}

	cmd := exec.Command(exePath, replayPath, "-o", outputPath)
	if err := cmd.Run(); err != nil {
		statusLabel.SetText(fmt.Sprintf("Error running r6-dissect: %v", err))
		return
	}

	// Find map image
	imagePath := findMapImage(mapName)

	// Create metadata
	metadata := MatchMetadata{
		MapName:    mapName,
		Score:      score,
		Date:       date,
		FilePath:   outputPath,
		ImagePath:  imagePath,
		ReplayPath: replayPath,
	}

	// Add to store
	store.Matches = append(store.Matches, metadata)
	saveStore(store)

	statusLabel.SetText("Match analyzed successfully!")
	window.SetContent(createSpreadsheetView(window, metadata, store, app))
}

func generateFilename(mapName, score string, date time.Time, store *MatchStore) string {
	// Format map name (replace spaces with underscores, remove special chars)
	mapNameFormatted := strings.ReplaceAll(mapName, " ", "_")
	mapNameFormatted = strings.ReplaceAll(mapNameFormatted, "'", "")
	
	// Format date as MM-DD-YYYY
	dateFormatted := date.Format("01-02-2006")
	
	// Base filename
	baseName := fmt.Sprintf("%s_%s_%s.xlsx", mapNameFormatted, score, dateFormatted)
	
	// Check for duplicates
	counter := 1
	filename := baseName
	for {
		exists := false
		for _, match := range store.Matches {
			if filepath.Base(match.FilePath) == filename {
				exists = true
				break
			}
		}
		if !exists {
			// Also check if file exists on disk
			if _, err := os.Stat(filepath.Join(getMatchesDir(), filename)); os.IsNotExist(err) {
				break
			}
		}
		counter++
		ext := filepath.Ext(baseName)
		nameWithoutExt := strings.TrimSuffix(baseName, ext)
		filename = fmt.Sprintf("%s_(%d)%s", nameWithoutExt, counter, ext)
	}
	
	return filename
}

func findMatchReplayFolder(launcher string) string {
	var paths []string
	
	if launcher == "steam" {
		paths = []string{
			`C:\Program Files (x86)\Steam\steamapps\common\Tom Clancy's Rainbow Six Siege\MatchReplay`,
			`C:\Program Files\Steam\steamapps\common\Tom Clancy's Rainbow Six Siege\MatchReplay`,
		}
	} else {
		paths = []string{
			`C:\Program Files (x86)\Ubisoft\Ubisoft Game Launcher\games\Tom Clancy's Rainbow Six Siege\MatchReplay`,
			`C:\Program Files\Ubisoft\Ubisoft Game Launcher\games\Tom Clancy's Rainbow Six Siege\MatchReplay`,
		}
	}

	// Check C drive first
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	// Check other drives
	drives := []string{"D:", "E:", "F:", "G:", "H:"}
	for _, drive := range drives {
		var searchPaths []string
		if launcher == "steam" {
			searchPaths = []string{
				filepath.Join(drive, "Program Files (x86)", "Steam", "steamapps", "common", "Tom Clancy's Rainbow Six Siege", "MatchReplay"),
				filepath.Join(drive, "Program Files", "Steam", "steamapps", "common", "Tom Clancy's Rainbow Six Siege", "MatchReplay"),
			}
		} else {
			searchPaths = []string{
				filepath.Join(drive, "Program Files (x86)", "Ubisoft", "Ubisoft Game Launcher", "games", "Tom Clancy's Rainbow Six Siege", "MatchReplay"),
				filepath.Join(drive, "Program Files", "Ubisoft", "Ubisoft Game Launcher", "games", "Tom Clancy's Rainbow Six Siege", "MatchReplay"),
			}
		}
		for _, path := range searchPaths {
			if _, err := os.Stat(path); err == nil {
				return path
			}
		}
	}

	return ""
}

func findMapImage(mapName string) string {
	// Use extracted embedded images directory
	imageDir := mapImagesDir
	if imageDir == "" {
		// Fallback to executable directory (for development)
		imageDir = filepath.Join(getExecutableDir(), "r6-maps-images")
	}
	
	// Normalize map name for matching
	normalizedName := strings.ToLower(strings.ReplaceAll(mapName, " ", ""))
	
	// Map name to image filename mapping (case-insensitive base matching)
	mapImageMap := map[string]string{
		"nighthavenlabs":     "ModernizedMap_Nighthaven_keyart.png",
		"nighthaven":         "ModernizedMap_Nighthaven_keyart.png",
		"consulate":           "ModernizedMap_Consulate_keyart.png",
		"lair":                "ModernizedMap_Lair_keyart.png",
		"coastline":           "r6-maps-coastline.png",
		"favela":              "r6-maps-favela__1_.png",
		"fortress":            "r6-maps-fortress.png",
		"herefordbase":        "r6-maps-hereford.png",
		"hereford":            "r6-maps-hereford.png",
		"house":               "r6-maps-house.png",
		"kanal":               "r6-maps-kanal.png",
		"oregon":              "r6-maps-oregon.png",
		"outback":             "r6-maps-outback.png",
		"presidentialplane":   "r6-maps-plane.png",
		"plane":               "r6-maps-plane.png",
		"skyscraper":          "r6-maps-skyscraper.png",
		"tower":               "r6-maps-tower.png",
		"villa":               "r6-maps-villa.png",
		"yacht":               "r6-maps-yacht.png",
		"bank":                "R6S_Maps_Bank_EXT.png",
		"border":              "R6S_Maps_Border_EXT.png",
		"chalet":              "R6S_Maps_Chalet_EXT.png",
		"clubhouse":           "R6S_Maps_ClubHouse_EXT.png",
		"club":                "R6S_Maps_ClubHouse_EXT.png",
		"kafedostoyevsky":     "R6S_Maps_RussianCafe_EXT.png",
		"kafe":                "R6S_Maps_RussianCafe_EXT.png",
		"emeraldplains":       "r6s_maps_emeraldplains__1_.png",
		"emerald":             "r6s_maps_emeraldplains__1_.png",
		"themepark":           "rainbow6_maps_theme-park_thumbnail.png",
		"theme":               "rainbow6_maps_theme-park_thumbnail.png",
		"stadiumbravo":        "stadiumB_keyart.png",
		"stadiumb":            "stadiumB_keyart.png",
		"stadium2020":         "StadiumA_keyart.png",
		"stadiuma":            "StadiumA_keyart.png",
		"stadium":             "StadiumA_keyart.png",
	}

	// Try exact normalized match first
	if img, ok := mapImageMap[normalizedName]; ok {
		path := filepath.Join(imageDir, img)
		if fileExists(path) {
			return path
		}
	}

	// Try Y10 variants (remove Y10 suffix)
	if strings.HasSuffix(normalizedName, "y10") {
		baseName := strings.TrimSuffix(normalizedName, "y10")
		if img, ok := mapImageMap[baseName]; ok {
			path := filepath.Join(imageDir, img)
			if fileExists(path) {
				return path
			}
		}
	}

	// Try partial matches (for maps like "ClubHouseY10" -> "clubhouse")
	for key, img := range mapImageMap {
		if strings.Contains(normalizedName, key) {
			path := filepath.Join(imageDir, img)
			if fileExists(path) {
				return path
			}
		}
	}

	// Default: return empty string if no match found
	return ""
}

func createSpreadsheetView(window fyne.Window, match MatchMetadata, store *MatchStore, app fyne.App) *container.Vertical {
	backButton := widget.NewButton("← Back", func() {
		window.SetContent(createMatchesView(window, store, app))
	})

	// Format date as MM-DD-YYYY
	dateStr := match.Date.Format("01-02-2006")
	title := widget.NewLabel(fmt.Sprintf("%s %s %s", match.MapName, match.Score, dateStr))
	title.TextStyle = fyne.TextStyle{Bold: true}

	// Load and display Excel file
	content := container.NewVBox(
		container.NewHBox(backButton, title),
		widget.NewSeparator(),
	)

	if !fileExists(match.FilePath) {
		errorLabel := widget.NewLabel(fmt.Sprintf("File not found: %s", match.FilePath))
		content.Add(errorLabel)
		return content
	}

	// Read Excel file and display
	excelViewer := createExcelViewer(match.FilePath)
	content.Add(excelViewer)

	return content
}

func createExcelViewer(filePath string) fyne.CanvasObject {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return widget.NewLabel(fmt.Sprintf("Error opening file: %v", err))
	}
	defer f.Close()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return widget.NewLabel("No sheets found in file")
	}

	// Read all sheet data before closing file
	sheetDataMap := make(map[string][][]string)
	for _, sheetName := range sheets {
		sheetDataMap[sheetName] = readSheetData(f, sheetName)
	}

	// Create tabs for each sheet
	tabs := container.NewAppTabs()
	
	for _, sheetName := range sheets {
		sheetData := sheetDataMap[sheetName]
		table := createSheetTable(sheetData)
		tabs.AppendTab(&container.TabItem{
			Text: sheetName,
			Content: container.NewScroll(table),
		})
	}

	return tabs
}

func readSheetData(f *excelize.File, sheetName string) [][]string {
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return [][]string{}
	}
	return rows
}

func createSheetTable(data [][]string) *widget.Table {
	if len(data) == 0 {
		return widget.NewTable(
			func() (int, int) { return 0, 0 },
			func() fyne.CanvasObject { return widget.NewLabel("") },
			nil,
		)
	}

	maxCols := 0
	for _, row := range data {
		if len(row) > maxCols {
			maxCols = len(row)
		}
	}

	table := widget.NewTable(
		func() (int, int) {
			return len(data), maxCols
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)
			if id.Row < len(data) && id.Col < len(data[id.Row]) {
				label.SetText(data[id.Row][id.Col])
			} else {
				label.SetText("")
			}
		},
	)

	// Set column widths (adjust based on content)
	for i := 0; i < maxCols && i < 10; i++ {
		width := 120.0
		if i == 0 {
			width = 150
		}
		table.SetColumnWidth(i, width)
	}
	return table
}

func getExecutableDir() string {
	if exe, err := os.Executable(); err == nil {
		return filepath.Dir(exe)
	}
	return "."
}

func getMatchesDir() string {
	return filepath.Join(getExecutableDir(), matchesDir)
}

func getStoreFile() string {
	return filepath.Join(getExecutableDir(), storeFile)
}

func loadStore() *MatchStore {
	data, err := os.ReadFile(getStoreFile())
	if err != nil {
		return &MatchStore{Matches: []MatchMetadata{}}
	}

	var store MatchStore
	if err := json.Unmarshal(data, &store); err != nil {
		return &MatchStore{Matches: []MatchMetadata{}}
	}

	return &store
}

func saveStore(store *MatchStore) {
	data, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return
	}
	os.WriteFile(getStoreFile(), data, 0644)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func findR6DissectExe() string {
	// Use extracted embedded executable
	if r6DissectPath != "" && fileExists(r6DissectPath) {
		return r6DissectPath
	}
	
	// Fallback to checking executable directory (for development)
	exeDir := getExecutableDir()
	paths := []string{
		filepath.Join(exeDir, "r6-dissect.exe"),
		"./r6-dissect.exe",
		"r6-dissect.exe",
	}
	
	for _, path := range paths {
		if fileExists(path) {
			return path
		}
	}
	
	return ""
}

// Dark theme for better UI
type darkTheme struct{}

func (t *darkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case fyne.ThemeColorNameBackground:
		return color.RGBA{R: 0x1e, G: 0x1e, B: 0x1e, A: 0xff}
	case fyne.ThemeColorNameButton:
		return color.RGBA{R: 0x3a, G: 0x3a, B: 0x3a, A: 0xff}
	case fyne.ThemeColorNameInputBackground:
		return color.RGBA{R: 0x2a, G: 0x2a, B: 0x2a, A: 0xff}
	case fyne.ThemeColorNameText:
		return color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	default:
		return color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	}
}

func (t *darkTheme) Font(style fyne.TextStyle) fyne.Resource {
	return fyne.DefaultTheme().Font(style)
}

func (t *darkTheme) Size(name fyne.ThemeSizeName) float32 {
	return fyne.DefaultTheme().Size(name)
}

