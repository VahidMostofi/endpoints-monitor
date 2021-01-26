FROM nginx:1.19-alpine

COPY ./nginx.conf.template /nginx.conf.template

CMD ["/bin/sh" , "-c" , "envsubst '${TELEGRAF_SYSLOG_SERVER},${PROXY_PASS_URL},${PORT}' < /nginx.conf.template > /etc/nginx/nginx.conf && exec nginx -g 'daemon off;'"]

# image name: nginx-gateway