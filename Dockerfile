FROM alpine:latest as build
RUN echo "This is a docker image for landers1037/blog"
RUN cd html/ && npm install && npm run build
COPY ./app_blog /app/
COPY ./example /app/example
COPY ./html/dist /app/html
COPY ./conf /app/conf
VOLUME ["/app/data", "/app/conf"]

EXPOSE 5000
WORKDIR /app
CMD ./app_blog web start 
