FROM alpine:latest
RUN echo "This is a docker image for landers1037/blog"
# start mv binary file
COPY ./app_blog /app/
# default port is 5000
COPY ./example /app/example
COPY ./dist/* /app/html/
EXPOSE 5000
WORKDIR /app

# workdir is /app conf is a path where app.ini must exist.
# data is used to save database and log file
VOLUME ["/app/data", "/app/conf"]
# just start in foreground
CMD /app/app_blog -p 5000