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
	@mkdir -p $(OUT)/assets
	@mkdir -p $(OUT)/tmpl
	@cp -R ./tmpl/*.html $(OUT)/tmpl
	@cp -R ./tmpl/static_assets/* $(OUT)/assets

clean:
	@rm -rf $(OUT)

test:
	$(GOX) test ./...

.PHONY: clean
