package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Pipe struct {
	Cert string `json:"ssl_file_cert,omitempty" xml:"ssl_file_cert,omitempty"`
	Key  string `json:"ssl_file_key,omitempty" xml:"ssl_file_key,omitempty"`
	Port uint   `json:"web_server_port,omitempty" xml:"web_server_port,omitempty"`
	Path string `json:"web_server_storage,omitempty" xml:"web_server_storage,omitempty"`
}

func (p *Pipe) check() error {
	if p.Port == 0 || p.Port > 65535 {
		p.Port = 80
	}
	if p.Path == "" {
		p.Path = "."
	}
	return nil
}

func (p *Pipe) Process() error {
	log.Printf("pipe server run, port: %v, path: %v.\n", p.Port, p.Path)
	defer log.Println("pipe server close")
	if err := p.check(); err != nil {
		return err
	}
	http.Handle("/", http.FileServer(http.Dir(p.Path)))
	if len(p.Cert) > 4 && len(p.Key) > 4 {
		if _, err := os.Stat(p.Cert); err == nil {
			if _, err := os.Stat(p.Key); err == nil {
				log.Printf("run pipe server https://<server>:%d", p.Port)
				return http.ListenAndServeTLS(fmt.Sprintf(":%d", p.Port), p.Cert, p.Key, nil)
			}
		}
	}
	log.Printf("run pipe  http://<server>:%d", p.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", p.Port), nil)
}

var p Pipe

func init() {
	flag.StringVar(&p.Cert, "cert", "", "ssl cert [certificate.crt]")
	flag.StringVar(&p.Key, "key", "", "ssl key [private.key]")
	flag.UintVar(&p.Port, "port", 80, "web server port [1..65535]\n")
	flag.StringVar(&p.Path, "path", ".", "path to storage [./ or ./static/ or ./storage/ ...]\n")
	flag.Parse()
}

func main() { log.Fatalln(p.Process()) }
