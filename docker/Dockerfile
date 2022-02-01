FROM node:16-alpine

WORKDIR /srv/app
EXPOSE 80

RUN set -ex; \
    mkdir -p /srv/app; \
    cd /srv/app; \
    yarn init -y; \
    yarn add express nodemon;

COPY ./app.js /srv/app/app.js

CMD ["node", "/srv/app/node_modules/.bin/nodemon", "/srv/app/app.js"]
