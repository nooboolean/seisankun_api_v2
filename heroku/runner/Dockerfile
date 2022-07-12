FROM alpine:3.9

COPY --from=builder /main .

ENV PORT=${PORT}
ENTRYPOINT ["/main"]