FROM node:20-alpine3.18

RUN apk add bash

WORKDIR /frontend

COPY . .

RUN npm install

CMD [ "npm", "run", "serve" ]