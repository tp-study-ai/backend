FROM golang
COPY main /project
WORKDIR /project
EXPOSE 8000
CMD ./main