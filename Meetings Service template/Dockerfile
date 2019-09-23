FROM node:10
ENV PORT 8092
EXPOSE 8092

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
COPY package.json .
RUN npm install
COPY . .

CMD ["npm", "run", "prod"]
