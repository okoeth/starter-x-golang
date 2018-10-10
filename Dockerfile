FROM scratch
EXPOSE 8080
ENTRYPOINT ["/starter-x-golang"]
COPY ./bin/ /