FROM oven/bun:1

WORKDIR /usr/src/app

COPY package.json bun.lockb ./

RUN bun install

COPY . .

CMD bun run index.ts