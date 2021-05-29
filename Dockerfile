FROM alpine:latest as build
RUN echo "This is a docker image for landers1037/blog"
# start mv binary file
COPY ./app_blog /app/
# default port is 5000
COPY ./example /app/example
COPY ./dist /app/html

# workdir is /app conf is a path where app.ini must exist.
# data is used to save database and log file
FROM alpine:latest
VOLUME ["/app/data", "/app/conf"]
EXPOSE 5000
WORKDIR /app
COPY --from=build /app /app
# just start in foreground
CMD ["/app/app_blog"]