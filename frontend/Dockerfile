# Use Bun as the base image
FROM oven/bun:latest as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the package.json and bun.lockb files
COPY package.json bun.lockb ./

# Install dependencies using Bun
RUN bun install

# Copy the rest of the application code
COPY . .

# Build the Qwik application
RUN bun run build

# Start a new image for the runtime
FROM oven/bun:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the build output from the builder stage
COPY --from=builder /app /app

# Expose the desired port (default for Vite is 3000)
EXPOSE 3000

# Set the command to start the application
CMD ["bun", "run", "dev"]
