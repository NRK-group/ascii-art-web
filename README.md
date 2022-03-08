# ASCII-ART-WEB

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
