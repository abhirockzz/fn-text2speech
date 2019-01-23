package main

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	fdk "github.com/fnproject/fdk-go"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(text2speech))
}

func text2speech(ctx context.Context, in io.Reader, out io.Writer) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(in)
	text := buf.String()
	log.Println("will convert text", text)

	outputLocation := "/tmp/output.wav"
	flite := exec.Command("flite", "-t", text, "-o", outputLocation)
	err := flite.Run()
	if err != nil {
		errMsg := "failed due to " + err.Error()
		log.Println(errMsg)
		out.Write([]byte(errMsg))
		return
	}

	log.Println("Converted file written to " + outputLocation)

	//delete converted file at the end
	defer func() {
		fileErr := os.Remove(outputLocation)
		if fileErr == nil {
			log.Println("Deleted temp file", outputLocation)
		} else {
			log.Println("Error removing output file", fileErr.Error())
		}
	}()

	//convert the file (.wav containing speech) to raw bytes
	speech, err := ioutil.ReadFile(outputLocation)

	if err != nil {
		errMsg := "could not read from .wav file " + err.Error()
		log.Println(errMsg)
		out.Write([]byte(errMsg))
		return
	}
	log.Println("Returning .wav bytes")

	out.Write(speech)
}
