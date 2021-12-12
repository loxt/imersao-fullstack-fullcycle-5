FROM node:16.13.1-alpine3.14

RUN apk add --no-cache bash curl

RUN curl https://raw.githubusercontent.com/eficode/wait-for/v2.1.3/wait-for --output /usr/bin/wait-for && \
  chmod +x /usr/bin/wait-for

RUN npm install -g @nestjs/cli

USER node
WORKDIR /home/node/app
