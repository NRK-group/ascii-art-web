# ASCII-ART-WEB

Ascii art web is a web app that convert the input string to a graphical representation. Currently, the graphiacal representation includes [standard](banner/standard.txt), [shadow](banner/shadow.txt) and [thinkertoy](banner/thinkertoy.txt).

## Authors

> Ricky Adriell

> kkamil

> G.Orlandi

## Dockerfile Intructions

1. Download [docker](https://docs.docker.com/get-docker/)
2. Run the command below in the terminal while inside the repository ascii-art-web to build the docker

> ```$ docker build -t ascii-art-web .```

3. Run the command below in the terminal while inside the repository ascii-art-web to run the docker

> ```$ docker run -p 5500:5500 -it ascii-art-web```

> or ```$ docker run -p 5500:5500 -d ascii-art-web``` to run this container in detached mode (run it permanently in the background)

4. Visit the [localhost:5500](http://localhost:5500/)

## Usage

1. Select the font of your choice
1. Type the text you would like to convert
1. Click the "Generate Art" button
1. To clear the text input use the backspace key or click the clear button

![ascii-art-web](./asciiartweb.gif)

## Test Implementation

> Run the command below in the terminal to test the error handling, Post end points, Get end points, and Ascii Art function

> ```$ go test .```

> or ```$ go test -v```
