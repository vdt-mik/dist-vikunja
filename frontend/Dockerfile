FROM node:22.13.1-alpine AS builder

WORKDIR /app

ENV PNPM_CACHE_FOLDER=.cache/pnpm/
ENV PUPPETEER_SKIP_DOWNLOAD=true
ENV CYPRESS_INSTALL_BINARY=0

COPY frontend/ ./
RUN npm install -g corepack && \
    corepack enable && \
    pnpm install && \
    pnpm run build

FROM nginx:alpine AS app

COPY frontend/nginx.conf /etc/nginx/conf.d/default.conf

COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"] 