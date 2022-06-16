FROM node:18-alpine3.15

WORKDIR /code

COPY . /code

RUN npm i

CMD ["node", "--test", "--trace-warnings"]
