# On OSX the PATH variable isn't exported unless "SHELL" is also set, see: http://stackoverflow.com/a/25506676
SHELL = /bin/bash
NODE_BINDIR = ./node_modules/.bin
export PATH := $(PATH):$(NODE_BINDIR)

INPUT_FILES = ./ui

# Where to write the files generated by this makefile.
OUTPUT_DIR = ./l10n

# Template file
TEMPLATE_FILE = ./l10n/template.pot

# Name of the generated .po files for each available locale.
LOCALE_FILES = $(shell find l10n/locale -name '*.po')

.PHONY: l10n-push
l10n-push:
	cd $(OUTPUT_DIR) && tx -d push -s --skip --no-interactive

.PHONY: l10n-pull
l10n-pull:
	cd $(OUTPUT_DIR) && tx -d pull -a --skip --minimum-perc=75

.PHONY: l10n-clean
l10n-clean:
	rm -f $(TEMPLATE_FILE)
	rm -rf $(OUTPUT_DIR)/locale

.PHONY: l10n-read
l10n-read: node_modules $(TEMPLATE_FILE)

.PHONY: l10n-write
l10n-write: node_modules $(OUTPUT_DIR)/translations.json

# Create a main .pot template, then generate .po files for each available language.
# Thanks to Systematic: https://github.com/Polyconseil/systematic/blob/866d5a/mk/main.mk#L167-L183
$(TEMPLATE_FILE):
# Extract gettext strings from each template file and create a POT dictionary template.
# Generate .po files for each available language.
	export GETTEXT_SOURCES=`find $(INPUT_FILES) -name '*.vue' -o -name '*.js'`; \
	node ./node_modules/easygettext/src/extract-cli.js --attribute v-translate --output $(OUTPUT_DIR)/template.pot $$GETTEXT_SOURCES;

# Generate translations.json file from .pot template.
.PHONY: $(OUTPUT_DIR)/translations.json
$(OUTPUT_DIR)/translations.json:
	rm -rf $(OUTPUT_DIR)/translations.json
	gettext-compile --output $(OUTPUT_DIR)/translations.json $(LOCALE_FILES);

