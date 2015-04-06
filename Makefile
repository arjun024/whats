CC = go build

CFLAGS  =

TARGET = whats

DEPS = ./whatslib/google/google.go

all: $(TARGET)

$(TARGET): $(TARGET).go $(DEPS)
	$(CC) $(CFLAGS) -o $(TARGET) $(TARGET).go
