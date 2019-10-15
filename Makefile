NAME		= RP42

# Compiler & Preprocessor flags
LDFLAGS		+= "-X github.com/alexandregv/RP42/pkg/oauth.API_CLIENT_ID=${RP42_CLIENT_ID} -X github.com/alexandregv/RP42/pkg/oauth.API_CLIENT_SECRET=${RP42_CLIENT_SECRET}"
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

# Multi platforms 
ECHO 		= echo
ifeq ($(shell uname),Linux)
	ECHO	+= -e
endif

all: $(NAME)

$(NAME):
	@$(ECHO) "$(NAME)\t[$(C_PENDING)⏳ $(C_RESET)]"
	@go build -o build/$(NAME) -ldflags $(LDFLAGS) cmd/$(NAME)/main.go
	@cp -R assets/macOS/$(NAME).app/ build/
	@cp build/$(NAME) build/$(NAME).app/Contents/MacOS/RP42
	@$(ERASE)
	@$(ECHO) "$(NAME)\t[$(C_SUCCESS)✅ $(C_RESET)]"
	@$(ECHO) "$(C_SUCCESS)Compilation successful! 👌 (./build/)$(C_RESET)"

clean:
	@#$(RM) -r build/ $(HIDE_ERR)

fclean: clean
	@$(RM) -rf build/

re: fclean all

.PHONY: clean fclean all re
