FROM node:18-alpine AS build

WORKDIR /app

# Copy only the necessary Node.js application files
COPY package.json package-lock.json ./
RUN npm ci --production

COPY . .

# Create a non-root user
RUN adduser --disabled-password -u 10001 --shell /bin/sh appuser && \
    chown -R appuser /app

USER appuser

EXPOSE 3000

# Run the Node.js application when the container starts
CMD ["npm", "start"]
