FROM nginx:1.21.3-alpine

LABEL maintainer="https://github.com/nekochans"

ADD ./docker/nginx/config/default.conf.template /etc/nginx/conf.d/default.conf.template
ADD ./docker/nginx/config/nginx.conf /etc/nginx/nginx.conf

CMD /bin/sh -c 'sed "s/\${BACKEND_HOST}/$BACKEND_HOST/" /etc/nginx/conf.d/default.conf.template  > /etc/nginx/conf.d/default.conf && nginx -g "daemon off;"'
