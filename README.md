Panduan install : 
1. Install  [Go](https://golang.org/doc/install) golang pada server dan atur [Go Workspace](https://golang.org/doc/code.html#Workspaces) pada server
2. Masuk ke $GOPAYH/src dan jalankan `go get https://github.com/seregant/tcc_workshop.git` 
3. Install [Govendor](https://github.com/kardianos/govendor) `go get github.com/kardianos/govendor`
4. Masuk ke folder $GOPATH/src/github.com/seregant/tugas-workshop
5. Jalankan `govendor sync` untuk menginstall library yang dibutuhkan
6. Atur konfigurasi port dan listen ip pada file config.json
7. Jalankan perintah `go run main.go`.

Dengan konfigurasi default aplikasi berjalan pada port 8000.


Panduan install (Docker) : 

1. Masuk ke folder $GOPATH/src/github.com/seregant/tugas-workshop
2. Build docker image dengan menjalankan command `docker build .`

aplikasi akan berjalan pada port 8000 pada container


Route list:
* `v1/schedule` : list of all tv show
* `v1/search?q=[string]` : search tv show

