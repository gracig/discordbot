FROM golang:latest as builder
LABEL Author="Gerson Graciani <15052330+gracig@users.noreply.github.com>"
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o discordbot .

FROM scratch
LABEL Author="Gerson Graciani <15052330+gracig@users.noreply.github.com>"
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /src/discordbot /
CMD ["/discordbot"]