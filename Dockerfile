FROM golang:1.22-alpine AS backend
WORKDIR /app
COPY backend/ ./backend/
WORKDIR /app/backend
RUN go mod download && go build -o server ./cmd/server

FROM node:20-alpine AS frontend
WORKDIR /app
COPY frontend/package*.json ./
RUN npm ci --only=production
COPY frontend/ ./
RUN npm run build

FROM alpine:3.19
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=backend /app/backend/server .
COPY --from=frontend /app/dist ./static

EXPOSE 8080
CMD ["./server"]
