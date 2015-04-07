CC = go build
CFLAGS  =
TARGET = whats
DEPS = ./whatslib/google/google.go
PREFIX ?= /usr/local

$(TARGET): $(TARGET).go $(DEPS)
	$(CC) $(CFLAGS) -o $(TARGET) $(TARGET).go

clean:
	rm -f $(TARGET)

uninstall:
	rm -f $(PREFIX)/bin/$(TARGET)

install: $(TARGET)
	install -m 0755 $(TARGET) $(PREFIX)/bin/$(TARGET)

all: $(TARGET)
