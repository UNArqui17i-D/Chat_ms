docker build -t chat_ms
docker run -it --rm --name chat_ms chat_ms

docker run --rm -v "$PWD":/usr/src/chat_ms -w /usr/src/chat_ms golang:1.6 go build -v
docker run --rm -v "$PWD":/usr/src/chat_ms -w /usr/src/chat_ms golang:1.6 bash -c make