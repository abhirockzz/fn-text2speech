# Function for Text to Speech conversion

This function converts text to speech using the [alpine `flite`](https://pkgs.alpinelinux.org/package/v3.8/main/x86_64/flite) package. 

It is installed in the runtime image using a custom Dockerfile which is only slightly different compared to the default one i.e. one line has been added to the runtime stage of the Docker build - `RUN apk add --no-cache flite`. This installs the `flite` package into the Docker image itself

The Go function 
- uses `os.Command` to invoke the `flite` binary
- which converts the text into speech and saves it into `/tmp/output.wav` (this is deleted before the function exits)
- the file is then read and the raw bytes are simply returned to the caller


## Pre-requisites

- Start by cloning this repository
- Configure your functions development environment
- Get latest Fn CLI - `curl -LSs https://raw.githubusercontent.com/fnproject/cli/master/install | sh`
- Switch to right context using `fn use context <context name>`

## Create an app

`fn create app text2speech --annotation oracle.com/oci/subnetIds=<SUBNETS>`

## Deploy to Oracle Functions

`fn -v deploy --app text2speech`

## Invoke (and listen..)

`echo -n 'Hope you enjoyed this text to speech example using Oracle Functions' | fn invoke text2speech convert > op.wav && afplay op.wav` 

`echo -n 'That was dope!' | fn invoke text2speech convert > op.wav && afplay op.wav` 

^^ (these work seamlessly on Mac)
