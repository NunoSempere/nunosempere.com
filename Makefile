# nunosempere.com
#
#   make            build the generator and produce the static site in build/
#   make serve      run the development server on http://127.0.0.1:8080
#   make deploy     upload build/ to Cloudflare Pages (needs wrangler, see server/README.md)
#   make test       go vet + unit tests
#   make clean      remove build/ and the binary

GO      ?= go
BIN      = server/webserver
SOURCE   = source
BUILD    = build
ADDR    ?= 127.0.0.1:8080
PROJECT ?= nunosempere-com

.PHONY: all site bin serve deploy test clean

all: site

bin: $(BIN)

$(BIN): $(wildcard server/*.go server/templates/* server/pub/*/* server/pub/*/*/*) server/go.mod server/go.sum
	cd server && $(GO) build -o webserver .

site: $(BIN)
	$(BIN) -root $(SOURCE) -build $(BUILD)

serve: $(BIN)
	$(BIN) -root $(SOURCE) -addr $(ADDR)

deploy: site
	wrangler pages deploy $(BUILD) --project-name $(PROJECT)

test:
	cd server && $(GO) vet ./... && $(GO) test ./...

clean:
	rm -rf $(BUILD) $(BIN)
