package main
import (
 "log"
 "net/http"
 "path/filepath"
 "strings"
)


// Define the MIME types for HLS
const (
 MimeTypeM3U8 = "application/vnd.apple.mpegurl"
 MimeTypeTS   = "video/mp2t"
)

// Base directory for HLS files
const baseDir = "C:/hlsco"
// Handler for serving HLS files


func hlsHandler(w http.ResponseWriter, r *http.Request) {
 // Log the incoming request URL
 log.Printf("Received request for: %s\n", r.URL.Path)
 // Ensure the request path starts with "/hls/"
 if !strings.HasPrefix(r.URL.Path, "/hls/") {
  http.Error(w, "Invalid request path", http.StatusBadRequest)
  return
 }
 // Get the file path from the request and map it to the base directory
 filePath := filepath.Join(baseDir, r.URL.Path[len("/hls/"):])
 // Log the resolved file path
 log.Printf("Resolved file path: %s\n", filePath)
 // Determine the MIME type based on the file extension
 var mimeType string
 switch filepath.Ext(filePath) {
 case ".m3u8":
  mimeType = MimeTypeM3U8
 case ".ts":
  mimeType = MimeTypeTS
  http.Error(w, "Unsupported file type", http.StatusUnsupportedMediaType)
  return
 }
 // Set the appropriate content type
 w.Header().Set("Content-Type", mimeType)
 // Serve the file
 http.ServeFile(w, r, filePath)
}


func main() {
 // Define the HLS route
 http.HandleFunc("/hls/", hlsHandler)
 // Start the server
 port := ":8080"
 log.Printf("Starting server on %s\n", port)
 if err := http.ListenAndServe(port, nil); err != nil {
  log.Fatalf("Could not start server: %s\n", err)
 }
}