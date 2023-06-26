FROM node:18.16.0-alpine as build-stage

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the entire app directory
COPY . .

# Build the Vue.js application
RUN npm run build

# Production-ready image
FROM nginx:latest

# Copy the built artifacts from the build-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html

# Expose the container port
EXPOSE 80

# Start Nginx server
CMD ["nginx", "-g", "daemon off;"]