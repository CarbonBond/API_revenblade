# Bundle static assets with nginx
FROM golang:bullseye as production
# Copy built assets from builder
WORKDIR /app

COPY . ./

RUN go generate ./... 

COPY ./server.go.PRODUCTION ./server.go

ENV PORT=80
# Expose port
EXPOSE 80

# Start sever
CMD ["go", "run", "./server.go"]

