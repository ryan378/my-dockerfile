# Menggunakan gambar golang sebagai base image
FROM golang:1.20-alpine

# Set working directory
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Download dependensi
RUN go mod init wise-word
RUN go mod tidy

# Build aplikasi
RUN go build -o wise-word

# Jalankan aplikasi
CMD ["./wise-word"]

# Expose port 80
EXPOSE 80