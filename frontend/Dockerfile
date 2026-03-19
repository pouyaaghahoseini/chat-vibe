FROM node:22-bookworm-slim

WORKDIR /app

COPY package.json ./
COPY tsconfig.json ./
COPY next-env.d.ts ./
COPY next.config.ts ./
COPY app ./app

RUN npm install
RUN npm run build

EXPOSE 3000

CMD ["npm", "run", "start"]