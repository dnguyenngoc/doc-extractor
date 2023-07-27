# Stage 1: Build the Vue.js app
FROM node:18.16.0 as build-stage

WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the app source code
COPY . .

# Build the Vue.js app for production
RUN npm run build

# Stage 2: Serve the Vue.js app with Nginx
FROM nginx:1.21-alpine as production-stage

# Copy the built Vue.js app from the previous stage
COPY --from=build-stage /app/dist /usr/share/nginx/html

# Copy the custom Nginx configuration file (optional)
# COPY nginx.conf /etc/nginx/nginx.conf

# Expose port 80
EXPOSE 80

# Start Nginx server
CMD ["nginx", "-g", "daemon off;"]
