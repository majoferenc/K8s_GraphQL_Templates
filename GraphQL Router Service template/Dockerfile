FROM node:10
ENV PORT 4000
EXPOSE 4000

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
COPY package.json .
RUN npm install
COPY . .

CMD ["npm", "run", "dev"]
