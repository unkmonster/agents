FROM golang AS builder

COPY . /src
WORKDIR /src/app/stats/service

# 合并服务配置和公共配置
RUN cp -r configs/* /src/configs 

RUN make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates  \
    netbase \
    && rm -rf /var/lib/apt/lists/ \
    && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/app/stats/service/bin /app
COPY --from=builder /src/configs /data/conf
COPY --from=builder /src/db/migrations/stats /migrations
WORKDIR /app

EXPOSE 8000
EXPOSE 9000

CMD sleep 10 && ./service -conf /data/conf
