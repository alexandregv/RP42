NAME		= RP42

# Compiler & Preprocessor flags
LDFLAGS		+= 
MAKEFLAGS	+= --no-print-directory

# Colors
C_RESET		= \033[0m
C_PENDING	= \033[0;36m
C_SUCCESS	= \033[0;32m

# Escape Sequences (ANSI/VT100)
ES_ERASE	= "\033[A\033[K\033[A"
ERASE		= $(ECHO) $(ES_ERASE)

# Hide STD/ERR and prevent Make from returning non-zero code
HIDE_STD	= > /dev/null
HIDE_ERR	= 2> /dev/null || true

# Cross platforms
ECHO 		= echo
ifeq ($(shell uname),Linux)
	ECHO	+= -e
endif

all: $(NAME)

$(NAME): linux windows macos
	@$(ECHO) "$(C_SUCCESS)Compilation successful! 👌 (./build/)$(C_RESET)"

linux:
	@$(ECHO) "Linux\t[$(C_PENDING)⏳ $(C_RESET)]"
	@GOOS=linux GOARCH=amd64 go build -o build/linux/$(NAME) -ldflags "$(LDFLAGS)" -tags legacy_appindicator cmd/$(NAME)/main.go
	@$(ERASE)
	@$(ECHO) "Linux\t[$(C_SUCCESS)✅ $(C_RESET)]"

windows:
	@$(ECHO) "Windows\t[$(C_PENDING)⏳ $(C_RESET)]"
	@GOOS=windows GOARCH=amd64 go build -o build/windows/$(NAME).exe -ldflags "-H=windowsgui $(LDFLAGS)" cmd/$(NAME)/main.go
	@$(ERASE)
	@$(ECHO) "Windows\t[$(C_SUCCESS)✅ $(C_RESET)]"

macos:
	@$(ECHO) "MacOS\t[$(C_PENDING)⏳ $(C_RESET)]"
	@env GOOS=darwin GOARCH=amd64 go build -o build/macOS/$(NAME) -ldflags "$(LDFLAGS)" cmd/$(NAME)/main.go
	@cp -R assets/macOS/ build/macOS/
	@cp build/macOS/$(NAME) build/macOS/$(NAME).app/Contents/MacOS/RP42
	@rm build/macOS/$(NAME).app/Contents/MacOS/.gitkeep
	@$(ERASE)
	@$(ECHO) "MacOS\t[$(C_SUCCESS)✅ $(C_RESET)]"

deploy:
	@cp build/macOS/RP42 /sgoinfre/goinfre/Perso/aguiot--/public/RP42
	@cp -R build/macOS/RP42.app /sgoinfre/goinfre/Perso/aguiot--/public

clean:
	@#$(RM) -r build/ $(HIDE_ERR)

fclean: clean
	@$(RM) -rf build/

re: fclean all

.PHONY: clean fclean all re
