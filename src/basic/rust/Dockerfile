FROM rust:1.46 as builder
WORKDIR /usr/src/myapp
COPY src/basic/client.rs src/basic/client.rs
COPY src/basic/server.rs src/basic/server.rs
COPY build.rs .
COPY Cargo.toml .
COPY Basic.proto .

RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v3.13.0/protoc-3.13.0-linux-x86_64.zip && unzip protoc-3.13.0-linux-x86_64.zip && ls -lha && cp bin/protoc . && ls -lha
RUN rustup component add rustfmt --toolchain 1.46.0-x86_64-unknown-linux-gnu
RUN cargo install --path .

FROM debian:buster-slim
RUN rm -rf /var/lib/apt/lists/*
COPY --from=builder /usr/local/cargo/bin/server /usr/local/bin/rust_server
COPY --from=builder /usr/local/cargo/bin/client /usr/local/bin/rust_client

CMD ["rust_server"]