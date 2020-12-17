FROM telegraf:1.14.5-alpine

COPY --from=golang:1.15-alpine /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"

COPY ./get_logs.go /get_logs.go

RUN export CGO_ENABLED=0; go build -o get_logs get_logs.go

COPY ./telegraf.conf /etc/telegraf/telegraf.conf

EXPOSE 6514

ENTRYPOINT ["/entrypoint.sh"]
CMD ["telegraf"]