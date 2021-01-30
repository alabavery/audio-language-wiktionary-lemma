FROM golang:alpine3.13

RUN mkdir tgt_dir_mt && mkdir toks_dir_mt && mkdir wds_dir_mt
COPY ./src /src
WORKDIR /src
ENTRYPOINT go run main.go -words='/wds_dir_mt/words.json' -tokens=/toks_dir_mt -target=/tgt_dir_mt
