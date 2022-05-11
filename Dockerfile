FROM scratch
COPY cwc /usr/bin/cwc
ENTRYPOINT ["/usr/bin/cwc"]