package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var serverName string

func main() {
	if len(os.Args) > 1 {
		serverName = os.Args[1]
	} else {
		println("serverName is required as argument")
		os.Exit(1)
	}

	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/servers", listServers)
	http.HandleFunc("/files/list", listFiles)
	http.HandleFunc("/files/read", readFile)
	http.HandleFunc("/files/save", saveFile)
	http.HandleFunc("/mods/list", listMods)
	http.HandleFunc("/mods/add", addMod)
	http.HandleFunc("/plugins/list", listPlugins)
	http.HandleFunc("/plugins/add", addPlugin)
	http.ListenAndServe(":3005", nil)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "web/src/index.html")
}

func listServers(w http.ResponseWriter, r *http.Request) {
	serversDir := filepath.Join(string(os.PathSeparator), "servers")
	files, err := os.ReadDir(serversDir)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		return
	}
	var servers []string
	for _, f := range files {
		if f.IsDir() {
			servers = append(servers, f.Name())
		}
	}
	json.NewEncoder(w).Encode(map[string]any{"servers": servers})
}

func listFiles(w http.ResponseWriter, r *http.Request) {
	srv := serverNameFromQuery(r)
	if srv == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]any{"error": "missing server"})
		return
	}
	dir := filepath.Join(string(os.PathSeparator), "servers", srv)
	files, err := os.ReadDir(dir)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		return
	}
	var out []string
	for _, f := range files {
		if !f.IsDir() {
			out = append(out, f.Name())
		}
	}
	json.NewEncoder(w).Encode(map[string]any{"files": out})
}

func readFile(w http.ResponseWriter, r *http.Request) {
	srv := serverNameFromQuery(r)
	file := r.URL.Query().Get("file")
	if srv == "" || file == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]any{"error": "missing server or file"})
		return
	}
	path := filepath.Join(string(os.PathSeparator), "servers", srv, file)
	data, err := os.ReadFile(path)
	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Write(data)
}

func saveFile(w http.ResponseWriter, r *http.Request) {
	srv := serverNameFromQuery(r)
	file := r.URL.Query().Get("file")
	if srv == "" || file == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]any{"error": "missing server or file"})
		return
	}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		return
	}
	path := filepath.Join(string(os.PathSeparator), "servers", srv, file)
	if strings.Contains(filepath.Base(file), "..") {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]any{"error": "invalid file name"})
		return
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		return
	}
	w.Write([]byte("ok"))
}

func listMods(w http.ResponseWriter, r *http.Request) {
	dir := filepath.Join(string(os.PathSeparator), "servers", serverName, "mods")
	files, _ := os.ReadDir(dir)
	var mods []string
	for _, f := range files {
		if !f.IsDir() {
			mods = append(mods, f.Name())
		}
	}
	json.NewEncoder(w).Encode(map[string]any{"mods": mods})
}

func addMod(w http.ResponseWriter, r *http.Request) {
	dir := filepath.Join(string(os.PathSeparator), "servers", serverName, "mods")
	if err := os.MkdirAll(dir, 0755); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "missing url", 400)
		return
	}
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer resp.Body.Close()
	fn := filepath.Base(url)
	dest := filepath.Join(dir, fn)
	out, err := os.Create(dest)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte("ok"))
}

func listPlugins(w http.ResponseWriter, r *http.Request) {
	dir := filepath.Join(string(os.PathSeparator), "servers", serverName, "plugins")
	files, _ := os.ReadDir(dir)
	var plugins []string
	for _, f := range files {
		if !f.IsDir() {
			plugins = append(plugins, f.Name())
		}
	}
	json.NewEncoder(w).Encode(map[string]any{"plugins": plugins})
}

func addPlugin(w http.ResponseWriter, r *http.Request) {
	dir := filepath.Join(string(os.PathSeparator), "servers", serverName, "plugins")
	if err := os.MkdirAll(dir, 0755); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "missing url", 400)
		return
	}
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer resp.Body.Close()
	fn := filepath.Base(url)
	dest := filepath.Join(dir, fn)
	out, err := os.Create(dest)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte("ok"))
}

func serverNameFromQuery(r *http.Request) string {
	srv := r.URL.Query().Get("server")
	if srv != "" {
		return srv
	}
	return serverName
}
