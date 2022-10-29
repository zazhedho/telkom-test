FROM golang
(untuk menggunakan base image golang, instruksi ini bakal menginstal Golang environment)

ADD . /go/src/github.com/telkomdev/indihome/backend
(menambahkan semua file yang ada di dalam folder saat ini ke dalam "/go/src/github.com/telkomdev/indihome/backend")

WORKDIR /go/src/github.com/telkomdev/indihome
(instruksi ini akan membuat folder "/go/src/github.com/telkomdev/indihome" jika belum ada dan akan mengatur working directory ke folder tersebut)

RUN go get github.com/tools/godep
(akan menjalankan perintah "go get" untuk menginstall dependensi ke dalam project)

RUN godep restore
(akan menjalankan perintah untuk me-restore godep)

RUN go install github.com/telkomdev/indihome
(akan menjalankan perintah "go install" untuk menginstall dependensi ke dalam container)

ENTRYPOINT /go/bin/indihome
(entrypoint mirip seperti CMD yang digunakan untuk menentukan command yang akan dieksekusi ketika container di jalankan, namun ENTRYPOINT tidak memungkinkan untuk mengganti command)

LISTEN 80
