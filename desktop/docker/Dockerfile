FROM node:22 as build

WORKDIR /app
COPY . .

RUN npm i && npm run package:angular-app

FROM nginx:alpine
COPY --from=build /app/workspaces/angular-app/.dist/angular-app /usr/share/nginx/html

COPY docker/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 80

ENTRYPOINT ["/entrypoint.sh"]