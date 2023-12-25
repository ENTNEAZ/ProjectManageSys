FROM ubuntu:latest

COPY ./Database_Homework /Database_Homework
COPY ./html /html

RUN chmod +x /Database_Homework

ENTRYPOINT ["/Database_Homework"]
