FROM node:22-alpine as build
LABEL authors="calapout"

WORKDIR /app

COPY package.json .
COPY yarn.lock .

RUN ["yarn", "install", "--frozen-lockfile"]

COPY . .

RUN ["yarn", "build"]

FROM nginx as prod

WORKDIR /usr/share/nginx/html

COPY --from=build /app/dist/ .