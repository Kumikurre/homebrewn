#############
### build ###
#############

# base image
FROM node:15.2.1 as build

# add app
COPY . /src

# set working directory
WORKDIR /src

# add `/app/node_modules/.bin` to $PATH
ENV PATH /node_modules/.bin:$PATH

# install and cache app dependencies
COPY package.json /src/package.json
RUN npm install


# generate build
CMD ["./build.sh"]