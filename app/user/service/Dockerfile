FROM golang AS builder

COPY . /src
WORKDIR /src/app/user/service

# 合并服务配置和公共配置
RUN cp -r configs/* /src/deploy/configs 

RUN make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates  \
    netbase \
    && rm -rf /var/lib/apt/lists/ \
    && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/app/user/service/bin /app
COPY --from=builder /src/deploy/configs /data/conf
COPY --from=builder /src/db/migrations/user /migrations
WORKDIR /app

EXPOSE 8000
EXPOSE 9000

CMD sleep 10 && ./service -conf /data/conf
