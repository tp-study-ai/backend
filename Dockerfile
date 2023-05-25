FROM golang
WORKDIR /project
COPY main /project
EXPOSE 8000
CMD ./main