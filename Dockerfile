FROM moby/buildkit:v0.9.3
WORKDIR /ara
COPY ara README.md /ara/
ENV PATH=/ara:$PATH
ENTRYPOINT [ "/ara/ara" ]
