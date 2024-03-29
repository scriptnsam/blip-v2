package tools

import (
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

// The `DownloadTool` function downloads a file from a specified URL and saves it to the user's
// Downloads folder with a specific name based on the operating system.
func DownloadTool(url string,toolName string) (string, error) {
	fileURL:=strings.TrimSpace(url)

	// Get the user's home directory
	usr,err:=user.Current()
	if err!=nil{
		log.Fatal("Error:",err)
		return "", err
	}

	// constructy the path to the downloads folder
	downloadFolder:=filepath.Join(usr.HomeDir,"Downloads")

	// create a new folder in the downloads folder
	newFolder:=filepath.Join(downloadFolder, "blip_tools")
	if err:=os.MkdirAll(newFolder,0755)	;err!=nil{
		log.Fatal("Error creating folder:",err)
		return "", err
	}
	
	// name of the file to save the downloaded content as
	var fileName string
	switch os :=runtime.GOOS; os{
	case "windows":
		fileName = filepath.Join(newFolder,toolName+".exe")
	case "darwin":
		fileName = filepath.Join(newFolder, toolName+".dmg")
	case "linux":
		fileName = filepath.Join(newFolder, toolName+".tar.gz")
	default:
		log.Fatal("Unsupported operating system:", os)
		return "", err
		
	}

	// Create the file to save the downloaded content
	file,err:=os.Create(fileName)
	if err!=nil{
		log.Fatal("Error creating file:", err)
		return "", err
	}
	defer file.Close()

	// Perform HTTP request to download the content
	resp,err:=http.Get(fileURL)
	if err!=nil{
		log.Fatal("Error downloading file:", err)
		return "", err
	}

	resp.Body.Close()

	// cherckl if the request was successful
	if resp.StatusCode!=http.StatusOK{
		log.Fatal("Error downloading file:", resp.Status)
		return "", err
	}

	// copy the content from the response to the file
	_,err=file.ReadFrom(resp.Body)
	if err!=nil{
		log.Fatal("Error copying file:", err)
		return "", err
	}

	log.Println("Tool downloaded successfully")
	return fileName, nil	
}