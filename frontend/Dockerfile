FROM node:16-alpine

RUN mkdir /frontend
COPY . /frontend/
WORKDIR /frontend
RUN npm install
EXPOSE 3000
ENTRYPOINT ["npm", "run", "dev", "--", "--host"]