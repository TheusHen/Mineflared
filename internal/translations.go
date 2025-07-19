package internal

var EnglishTranslations = map[string]string{
	// Root command
	"CLI_SHORT_DESC": "CLI to create and manage Minecraft servers",
	
	// Login command
	"LOGIN_SHORT_DESC": "Login via GitHub",
	"LOGIN_IP_ERROR": "Could not get IP. Try manually.",
	"LOGIN_BROWSER_OPENING": "Opening browser for GitHub login...",
	"LOGIN_TOKEN_ERROR": "Error receiving token. Try again.",
	"LOGIN_SUCCESS": "✅ Login successful!",
	"LOGIN_PORT_ERROR": "Error opening port 4500:",
	"LOGIN_TOKEN_NOT_RECEIVED": "Token not received",
	"LOGIN_BROWSER_SUCCESS": "✅ Login successful. You can close this window.",
	"LOGIN_WAITING": "Waiting for login...",
	
	// Create command
	"CREATE_SHORT_DESC": "Create a new Minecraft server",
	"CREATE_LOGIN_REQUIRED": "You need to be logged in. Use 'mineflared-cli login'",
	"CREATE_SERVER_NAME_PROMPT": "Enter server name: ",
	"CREATE_SERVER_NAME_EMPTY": "Server name cannot be empty!",
	"CREATE_DIR_ERROR": "Error creating server directory:",
	"CREATE_SERVER_TYPE_PROMPT": "Choose server type:",
	"CREATE_JAVA_OPTION": "[1] Java",
	"CREATE_BEDROCK_OPTION": "[2] Bedrock",
	"CREATE_OPTION_PROMPT": "Option: ",
	"CREATE_JAVA_VERSION_PROMPT": "Choose Java version:",
	"CREATE_JAVA_VERSION_OPTION": "[%d] %s (Mods: %v, Plugins: %v)",
	"CREATE_INVALID_OPTION": "Invalid option.",
	"CREATE_BEDROCK_WARNING": "WARNING: Bedrock server only supports Windows for now.",
	"CREATE_BEDROCK_CANCEL": "You can cancel now (type 'n' to cancel) or press Enter to continue.",
	"CREATE_ACTION_CANCELLED": "Action cancelled.",
	"CREATE_DOWNLOADING": "Downloading server file...",
	"CREATE_DOWNLOAD_ERROR": "Error downloading file:",
	"CREATE_DOWNLOAD_COMPLETE": "File downloaded to:",
	"CREATE_MODS_SUPPORT": "This server supports MODS. Enter direct links to MODs (.jar, .zip) or press Enter to search for local files:",
	"CREATE_PLUGINS_SUPPORT": "This server supports Plugins. Enter direct links to Plugins (.jar, .zip) or press Enter to search for local files:",
	"CREATE_NO_MODS_PLUGINS": "This server does NOT support mods/plugins.",
	"CREATE_DOWNLOADING_FILE": "Downloading:",
	"CREATE_DOWNLOAD_FILE_ERROR": "Error downloading %s: %v",
	"CREATE_LOCAL_FILES_PROMPT": "Searching for local files. Place files (.jar/.zip) in the current folder and press Enter when ready.",
	"CREATE_READ_DIR_ERROR": "Error reading current directory:",
	"CREATE_MOVE_FILES_ERROR": "Error moving files:",
	"CREATE_JAVA_SERVER_STARTING": "Starting Java server...",
	"CREATE_JAVA_SERVER_EXECUTE": "Execute:\n%s",
	"CREATE_BEDROCK_EXTRACTING": "Extracting Bedrock Server...",
	"CREATE_BEDROCK_EXTRACT_ERROR": "Error extracting:",
	"CREATE_BEDROCK_PACKS_PROMPT": "Enter links to behavior/resource packs (.zip/.mcpack/.mcaddon), separated by space, or press Enter to search for local files:",
	"CREATE_BEDROCK_LOCAL_FILES_PROMPT": "Searching for local files. Place files (.zip/.mcpack/.mcaddon) in the current folder and press Enter when ready.",
	"CREATE_BEDROCK_MOVE_ERROR": "Error moving file:",
	"CREATE_BEDROCK_READY": "Bedrock server ready to start! Enter the bedrock-server-1.21.95.1 folder and run bedrock_server.exe",
	"CREATE_API_ERROR": "Error creating server",
	"CREATE_SERVER_CREATED": "Server created! Waiting for DNS propagation...",
	"CREATE_SERVER_ACCESS": "Access: https://%s.mineserver.theushen.me",
	"INVALID_FILE": "Invalid file: %s",
	"FILE_MOVED": "Moved: %s -> %s",
	
	// Status command
	"STATUS_SHORT_DESC": "Check server status",
	"STATUS_LOGIN_REQUIRED": "❌ You need to be logged in first!",
	"STATUS_REQUEST_ERROR": "❌ Error creating request:",
	"STATUS_QUERY_ERROR": "❌ Error querying status:",
	"STATUS_DECODE_ERROR": "❌ Error decoding response:",
	"STATUS_API_ERROR": "❌ Error: %v",
	"STATUS_UNKNOWN_ERROR": "❌ Unknown error. HTTP Code: %d",
	"STATUS_SERVER_STATUS": "✅ Server status:",
	"STATUS_STATUS_LINE": "   • Status: %s",
	"STATUS_MESSAGE_LINE": "   • Message: %s",
	
	// DNS update
	"DNS_IP_ERROR": "Could not get current IP.",
	"DNS_UPDATE_ERROR": "Error updating DNS:",
	"DNS_UPDATE_SUCCESS": "DNS updated successfully.",
	"DNS_NO_UPDATE": "IP hasn't changed, DNS doesn't need updating.",
	
	// Language command
	"LANGUAGE_SHORT_DESC": "Change the interface language",
	"LANGUAGE_CURRENT": "Current language: %s",
	"LANGUAGE_PROMPT": "Choose language:\n[1] English\n[2] Portuguese",
	"LANGUAGE_OPTION_PROMPT": "Option: ",
	"LANGUAGE_INVALID_OPTION": "Invalid option. Please choose 1 for English or 2 for Portuguese.",
	"LANGUAGE_CHANGED": "Language changed to: %s",
	
	// First run language selection
	"FIRST_RUN_LANGUAGE": "Welcome to mineflared-cli!\nPlease select your preferred language:",
	"FIRST_RUN_ENGLISH": "[1] English",
	"FIRST_RUN_PORTUGUESE": "[2] Portuguese (Português)",
}

var PortugueseTranslations = map[string]string{
	// Root command
	"CLI_SHORT_DESC": "CLI para criar e gerenciar servidores de Minecraft",
	
	// Login command
	"LOGIN_SHORT_DESC": "Faz login via GitHub",
	"LOGIN_IP_ERROR": "Não foi possível obter IP. Tente manualmente.",
	"LOGIN_BROWSER_OPENING": "Abrindo navegador para login com GitHub...",
	"LOGIN_TOKEN_ERROR": "Erro ao receber token. Tente novamente.",
	"LOGIN_SUCCESS": "✅ Login realizado com sucesso!",
	"LOGIN_PORT_ERROR": "Erro ao abrir porta 4500:",
	"LOGIN_TOKEN_NOT_RECEIVED": "Token não recebido",
	"LOGIN_BROWSER_SUCCESS": "✅ Login realizado com sucesso. Você pode fechar esta janela.",
	"LOGIN_WAITING": "Aguardando login...",
	
	// Create command
	"CREATE_SHORT_DESC": "Cria um novo servidor Minecraft",
	"CREATE_LOGIN_REQUIRED": "Você precisa estar logado. Use 'mineflared-cli login'",
	"CREATE_SERVER_NAME_PROMPT": "Digite o nome do servidor: ",
	"CREATE_SERVER_NAME_EMPTY": "Nome do servidor não pode ser vazio!",
	"CREATE_DIR_ERROR": "Erro ao criar diretório do servidor:",
	"CREATE_SERVER_TYPE_PROMPT": "Escolha o tipo de servidor:",
	"CREATE_JAVA_OPTION": "[1] Java",
	"CREATE_BEDROCK_OPTION": "[2] Bedrock",
	"CREATE_OPTION_PROMPT": "Opção: ",
	"CREATE_JAVA_VERSION_PROMPT": "Escolha a versão do Java:",
	"CREATE_JAVA_VERSION_OPTION": "[%d] %s (Mods: %v, Plugins: %v)",
	"CREATE_INVALID_OPTION": "Opção inválida.",
	"CREATE_BEDROCK_WARNING": "ATENÇÃO: O servidor Bedrock só suporta Windows por enquanto.",
	"CREATE_BEDROCK_CANCEL": "Você pode cancelar agora (digite 'n' para cancelar) ou pressione Enter para continuar.",
	"CREATE_ACTION_CANCELLED": "Ação cancelada.",
	"CREATE_DOWNLOADING": "Baixando arquivo do servidor...",
	"CREATE_DOWNLOAD_ERROR": "Erro ao baixar o arquivo:",
	"CREATE_DOWNLOAD_COMPLETE": "Arquivo baixado em:",
	"CREATE_MODS_SUPPORT": "Este servidor suporta MODS. Insira os links diretos para os MODs (.jar, .zip) ou pressione Enter para buscar arquivos localmente:",
	"CREATE_PLUGINS_SUPPORT": "Este servidor suporta Plugins. Insira os links diretos para os Plugins (.jar, .zip) ou pressione Enter para buscar arquivos localmente:",
	"CREATE_NO_MODS_PLUGINS": "Este servidor NÃO suporta mods/plugins.",
	"CREATE_DOWNLOADING_FILE": "Baixando:",
	"CREATE_DOWNLOAD_FILE_ERROR": "Erro ao baixar %s: %v",
	"CREATE_LOCAL_FILES_PROMPT": "Buscando arquivos localmente. Coloque os arquivos (.jar/.zip) na pasta atual e pressione Enter quando pronto.",
	"CREATE_READ_DIR_ERROR": "Erro ao ler diretório atual:",
	"CREATE_MOVE_FILES_ERROR": "Erro ao mover arquivos:",
	"CREATE_JAVA_SERVER_STARTING": "Iniciando servidor Java...",
	"CREATE_JAVA_SERVER_EXECUTE": "Execute:\n%s",
	"CREATE_BEDROCK_EXTRACTING": "Descompactando Bedrock Server...",
	"CREATE_BEDROCK_EXTRACT_ERROR": "Erro ao descompactar:",
	"CREATE_BEDROCK_PACKS_PROMPT": "Insira os links para behavior/resource packs (.zip/.mcpack/.mcaddon), separados por espaço, ou pressione Enter para buscar arquivos localmente:",
	"CREATE_BEDROCK_LOCAL_FILES_PROMPT": "Buscando arquivos localmente. Coloque os arquivos (.zip/.mcpack/.mcaddon) na pasta atual e pressione Enter quando pronto.",
	"CREATE_BEDROCK_MOVE_ERROR": "Erro ao mover arquivo:",
	"CREATE_BEDROCK_READY": "Servidor Bedrock pronto para iniciar! Entre na pasta bedrock-server-1.21.95.1 e execute bedrock_server.exe",
	"CREATE_API_ERROR": "Erro ao criar servidor",
	"CREATE_SERVER_CREATED": "Servidor criado! Aguarde a propagação do DNS...",
	"CREATE_SERVER_ACCESS": "Acesse: https://%s.mineserver.theushen.me",
	"INVALID_FILE": "Arquivo inválido: %s",
	"FILE_MOVED": "Movido: %s -> %s",
	
	// Status command
	"STATUS_SHORT_DESC": "Verifica o status do servidor",
	"STATUS_LOGIN_REQUIRED": "❌ Você precisa estar logado primeiro!",
	"STATUS_REQUEST_ERROR": "❌ Erro ao criar requisição:",
	"STATUS_QUERY_ERROR": "❌ Erro ao consultar status:",
	"STATUS_DECODE_ERROR": "❌ Erro ao decodificar resposta:",
	"STATUS_API_ERROR": "❌ Erro: %v",
	"STATUS_UNKNOWN_ERROR": "❌ Erro desconhecido. Código HTTP: %d",
	"STATUS_SERVER_STATUS": "✅ Status do servidor:",
	"STATUS_STATUS_LINE": "   • Status: %s",
	"STATUS_MESSAGE_LINE": "   • Mensagem: %s",
	
	// DNS update
	"DNS_IP_ERROR": "Não foi possível obter o IP atual.",
	"DNS_UPDATE_ERROR": "Erro ao atualizar DNS:",
	"DNS_UPDATE_SUCCESS": "DNS atualizado com sucesso.",
	"DNS_NO_UPDATE": "IP não mudou, DNS não precisa de atualização.",
	
	// Language command
	"LANGUAGE_SHORT_DESC": "Altera o idioma da interface",
	"LANGUAGE_CURRENT": "Idioma atual: %s",
	"LANGUAGE_PROMPT": "Escolha o idioma:\n[1] Inglês\n[2] Português",
	"LANGUAGE_OPTION_PROMPT": "Opção: ",
	"LANGUAGE_INVALID_OPTION": "Opção inválida. Por favor, escolha 1 para Inglês ou 2 para Português.",
	"LANGUAGE_CHANGED": "Idioma alterado para: %s",
	
	// First run language selection
	"FIRST_RUN_LANGUAGE": "Bem-vindo ao mineflared-cli!\nPor favor, selecione seu idioma preferido:",
	"FIRST_RUN_ENGLISH": "[1] Inglês (English)",
	"FIRST_RUN_PORTUGUESE": "[2] Português",
}

func GetTranslation(key string) string {
	cfg := GetConfig()
	
	if cfg.Language == "en" {
		if translation, ok := EnglishTranslations[key]; ok {
			return translation
		}
	} else {
		if translation, ok := PortugueseTranslations[key]; ok {
			return translation
		}
	}
	
	// Default to Portuguese if translation not found or language not set
	if translation, ok := PortugueseTranslations[key]; ok {
		return translation
	}
	
	// Return the key if no translation is found
	return key
}

func SetLanguage(lang string) {
	cfg := GetConfig()
	cfg.Language = lang
	SaveConfig()
}

func GetLanguageName() string {
	cfg := GetConfig()
	if cfg.Language == "en" {
		return "English"
	}
	return "Português"
}