GOX := $(shell which go)
BIN := wikara
OUT := dist

dist:
	@mkdir -p $(OUT)
	$(GOX) build \
		-v \
		-x \
		-o ./$(OUT)/$(BIN) \
		./cmd/$(BIN)
	@cp config.example.yml $(OUT)/config.yml
	@cp -R ./tmpl $(OUT)/tmpl

clean:
	@rm -rf $(OUT)

test:
	$(GOX) test ./...

.PHONY: clean
