package main

import (
	"encoding/json"
	. "fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	Printfln("Working with files")
	for _, p := range Products {
		Printfln("Product: %v, Category: %v, Price: $%.2f",
			p.Name, p.Category, p.Price)
	}

	Printfln("\nDecoding the JSON Data")
	// see readconfig.go
	Printfln("\nReading from a Specific Location")
	read_at()
	Printfln("\nWriting to Files%s", "\n    Using the Write Convenience Function")

	total := 0.0
	for _, p := range Products {
		total += p.Price
	}
	dataStr := Sprintf("Time: %v, Total: $%.2f\n",
		time.Now().Format("2006, Mon 15:04:05"), total)
	err := os.WriteFile("output.txt", []byte(dataStr), 0666)
	if err == nil {
		Println("Output file created")
	} else {
		Printfln("Error: %v", err.Error())
	}

	Println("\n    Using the File Struct to Write to a File")

	file, err := os.OpenFile("output2.txt",
		os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		defer func() { _ = file.Close() }()
		_, _ = file.WriteString(dataStr + "aaaa\n")
	} else {
		Printfln("Error: %v", err.Error())
	}
	Println("\n    Writing JSON Data to a File")
	cheapProducts := []Product{}
	for _, p := range Products {
		if p.Price < 100 {
			cheapProducts = append(cheapProducts, p)
		}
	}
	//file, err = os.OpenFile("cheap.json", os.O_WRONLY|os.O_CREATE, 0666)
	/*
		file, err = os.CreateTemp(".", "tempfile-*.json")
		if err == nil {
			defer func() { _ = file.Close() }()
			encoder := json.NewEncoder(file)
			_ = encoder.Encode(cheapProducts)
		} else {
			Printfln("Error: %v", err.Error())
		}
	*/
	Println("\nWorking with File Paths")
	path, err := os.UserHomeDir()
	Printfln("Volume name: %v", filepath.VolumeName(path))
	Printfln("Dir component: %v", filepath.Dir(path))
	Printfln("File component: %v", filepath.Base(path))
	Printfln("File extension: %v", filepath.Ext(path))
	path, _ = os.Getwd()
	Printfln("\nCWD:%s", path)
	if err == nil {
		path = filepath.Join(path, "MyApp", "MyTempFile.json")
	}

	// 585Chapter 22 ■ Working with Files
	Printfln("Full path: %v", path)
	err = os.MkdirAll(filepath.Dir(path), 0755)
	if err == nil {
		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
		if err == nil {
			//defer func() { _ = file.Close() }()
			//defer file.Close()
			encoder := json.NewEncoder(file)
			_ = encoder.Encode(Products)
		}
		// 587Chapter 22 ■ Working with Files
	}
	if err != nil {
		Printfln("Error %v", err.Error())
	}
	path, _ = os.UserCacheDir()
	Printfln("UserCacheDir:%s", path)
	path, _ = os.UserConfigDir()
	Printfln("UserConfigDir:%s", path)
	home, _ := os.UserHomeDir()

	_ = os.Setenv("TMPDIR", filepath.Join(home, "tmp"))
	defer func() { _ = os.Unsetenv("TMPDIR") }()
	path = os.TempDir()
	Printfln("TempDir:%s", path)
	path = filepath.Join(path, "MyApp", "MyTempFile.json")
	err = os.MkdirAll(filepath.Dir(path), 0766)
	if err == nil {
		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			defer func() { _ = file.Close() }()
			encoder := json.NewEncoder(file)
			_ = encoder.Encode(Products)
		}
		// 587Chapter 22 ■ Working with Files
	}
	Printfln("New joined full path: %s", path)
	if err != nil {
		Printfln("Error %v", err.Error())
	}
	Println("\nExploring the File System")

	func() {
		path, err := os.Getwd()
		if err == nil {
			dirEntries, err := os.ReadDir(path)
			if err == nil {
				for _, ent := range dirEntries {
					inf, _ := ent.Info()
					stat, _ := os.Stat(path)
					Printfln("Entry name: %s, IsDir: %t, Type(): %s\nInfo():\n Size(): %.4f KB, Mode(): %s, ModTime(): %s\n Stat(): %v ",
						ent.Name(), ent.IsDir(), ent.Type(),
						float64(inf.Size())/1024, inf.Mode(), inf.ModTime(), stat)
				}
			}
		}
		if err != nil {
			Printfln("Error %v", err.Error())
		}
	}()

	Println("\nDetermining Whether a File Exists")

	targetFiles := []string{"no_such_file.txt", "config.json"}
	for _, name := range targetFiles {
		info, err := os.Stat(name)
		if os.IsNotExist(err) {
			Printfln("File does not exist: %v", name)
		} else if err != nil {
			Printfln("Other error: %v", err.Error())
		} else {
			Printfln("File %v, Size: %v", info.Name(), info.Size())
		}
	}

	Println("\nLocating Files Using a Pattern")

	path, err = os.Getwd()
	if err == nil {
		matches, err := filepath.Glob(filepath.Join(path, "*.[jg]*"))
		// 591Chapter 22 ■ Working with Files
		if err == nil {
			for _, m := range matches {
				Printfln("Match: %v", m)
			}
		}
	}
	if err != nil {
		Printfln("Error %v", err.Error())
	}

	Println("\nProcessing All Files in a Directory")
	path, err = os.Getwd()
	if err == nil {
		err = filepath.WalkDir(path, callback)
	} else {
		Printfln("Error %v", err.Error())
	}
}

type wanted func(path string, dir os.DirEntry, f os.FileInfo, wg *sync.WaitGroup)

// the `wanted` function
func callback(path string, dir os.DirEntry, dirErr error) (err error) {
	info, _ := dir.Info()
	Printfln("Path %v, Size: %v IsDir: %t", path, info.Size(), dir.IsDir())
	return
}

// for more examples see https://gosamples.dev/read-file/
func read_at() {
	f, e := os.Open("seekfile.txt")
	if e == nil {
		one_elem_slice := make([]byte, 5)
		num, e := f.ReadAt(one_elem_slice, 10)
		if e != nil {
			println("Error:", e.Error())
		}
		Printfln("Read bytes: %d, Contents: %#v", num, string(one_elem_slice))
	}

}
