package internal

var EnglishTranslations = map[string]string{
	// Root command
	"CLI_SHORT_DESC": "CLI to create and manage Minecraft servers",

	// Login command
	"LOGIN_SHORT_DESC":         "Login via GitHub",
	"LOGIN_IP_ERROR":           "Could not get IP. Try manually.",
	"LOGIN_BROWSER_OPENING":    "Opening browser for GitHub login...",
	"LOGIN_TOKEN_ERROR":        "Error receiving token. Try again.",
	"LOGIN_SUCCESS":            "✅ Login successful!",
	"LOGIN_PORT_ERROR":         "Error opening port 4500:",
	"LOGIN_TOKEN_NOT_RECEIVED": "Token not received",
	"LOGIN_BROWSER_SUCCESS":    "✅ Login successful. You can close this window.",
	"LOGIN_WAITING":            "Waiting for login...",

	// Delete command
	"DELETE_SHORT_DESC":     "Delete your account and all associated data from Mineflared.",
	"DELETE_CONFIRM_PROMPT": "Are you sure you want to permanently delete your account and all data? (y/N): ",
	"DELETE_ABORTED":        "Aborted.",
	"DELETE_LOGIN_REQUIRED": "You must be logged in to delete your account.",
	"DELETE_REQUEST_FAILED": "Failed to connect to server:",
	"DELETE_SUCCESS":        "Account deleted successfully.",
	"DELETE_ERROR":          "Error deleting account:",
	
	// List command
	"LIST_SHORT_DESC":        "List all local servers",
	"LIST_SERVERS_DIR_ERROR": "Could not open servers directory: %s",
	"LIST_HEADER":            "Servers (current owner):",

	// Backup command
	"BACKUP_SHORT_DESC":       "Create a backup of your server",
	"BACKUP_SERVER_NOT_FOUND": "Server not found.",
	"BACKUP_FORMAT_PROMPT":    "Choose backup format (zip/rar): ",
	"BACKUP_INVALID_FORMAT":   "Invalid format. Only zip or rar are allowed.",
	"BACKUP_CHOOSE_PATH":      "Choose where to save the backup file.",
	"BACKUP_CURRENT_DIR":      "Current directory:",
	"BACKUP_PATH_PROMPT":      "Type the folder path to save (leave blank for current): ",
	"BACKUP_CREATING_ZIP":     "Creating ZIP backup...",
	"BACKUP_CREATING_RAR":     "Creating RAR backup (requires 'rar' installed)...",
	"BACKUP_ZIP_ERROR":        "Error creating ZIP: %s\n",
	"BACKUP_RAR_ERROR":        "Error creating RAR: %s\n",
	"BACKUP_DONE":             "Backup created at: %s",

	// Restore command
	"RESTORE_SHORT_DESC":     "Restore a server from backup",
	"RESTORE_CHOOSE_PATH":    "Choose a backup file to restore:",
	"RESTORE_CURRENT_DIR":    "Current directory:",
	"RESTORE_FILE_PROMPT":    "Type the file name to restore (with .zip or .rar): ",
	"RESTORE_FILE_EMPTY":     "No file name entered.",
	"RESTORE_FILE_NOT_FOUND": "Backup file not found.",
	"RESTORE_SERVER_EXISTS":  "A server with this name already exists. Delete it first if you want to restore.",
	"RESTORE_UNZIPPING":      "Extracting ZIP...",
	"RESTORE_UNZIP_ERROR":    "Error extracting ZIP: %s\n",
	"RESTORE_UNRAR":          "Extracting RAR (requires 'unrar' installed)...",
	"RESTORE_UNRAR_ERROR":    "Error extracting RAR: %s\n",
	"RESTORE_INVALID_FORMAT": "Unsupported file format.",
	"RESTORE_DONE":           "Restore completed at: %s",

	// Config command
	"CONFIG_RUNNING_BLOCK":    "Cannot configure while the server is running.",
	"START_CONFIGURING_BLOCK": "Cannot start server while configuration panel is open.",
	"CONFIG_SHORT_DESC":       "Open the web-based server configuration panel",
	"CONFIG_OPEN_BROWSER":     "Opening server config panel at: %s",

	// Create command
	"CREATE_SHORT_DESC":                 "Create a new Minecraft server",
	"CREATE_LOGIN_REQUIRED":             "You need to be logged in. Use 'mineflared-cli login'",
	"CREATE_SERVER_NAME_PROMPT":         "Enter server name: ",
	"CREATE_SERVER_NAME_EMPTY":          "Server name cannot be empty!",
	"CREATE_DIR_ERROR":                  "Error creating server directory:",
	"CREATE_SERVER_TYPE_PROMPT":         "Choose server type:",
	"CREATE_JAVA_OPTION":                "[1] Java",
	"CREATE_BEDROCK_OPTION":             "[2] Bedrock",
	"CREATE_OPTION_PROMPT":              "Option: ",
	"CREATE_JAVA_VERSION_PROMPT":        "Choose Java version:",
	"CREATE_JAVA_VERSION_OPTION":        "[%d] %s (Mods: %v, Plugins: %v)",
	"CREATE_INVALID_OPTION":             "Invalid option.",
	"CREATE_BEDROCK_WARNING":            "WARNING: Bedrock server only supports Windows for now.",
	"CREATE_BEDROCK_CANCEL":             "You can cancel now (type 'n' to cancel) or press Enter to continue.",
	"CREATE_ACTION_CANCELLED":           "Action cancelled.",
	"CREATE_DOWNLOADING":                "Downloading server file...",
	"CREATE_DOWNLOAD_ERROR":             "Error downloading file:",
	"CREATE_DOWNLOAD_COMPLETE":          "File downloaded to:",
	"CREATE_MODS_SUPPORT":               "This server supports MODS. Enter direct links to MODs (.jar, .zip) or press Enter to search for local files:",
	"CREATE_PLUGINS_SUPPORT":            "This server supports Plugins. Enter direct links to Plugins (.jar, .zip) or press Enter to search for local files:",
	"CREATE_NO_MODS_PLUGINS":            "This server does NOT support mods/plugins.",
	"CREATE_DOWNLOADING_FILE":           "Downloading:",
	"CREATE_DOWNLOAD_FILE_ERROR":        "Error downloading %s: %v",
	"CREATE_LOCAL_FILES_PROMPT":         "Searching for local files. Place files (.jar/.zip) in the current folder and press Enter when ready.",
	"CREATE_READ_DIR_ERROR":             "Error reading current directory:",
	"CREATE_MOVE_FILES_ERROR":           "Error moving files:",
	"CREATE_JAVA_SERVER_STARTING":       "Starting Java server...",
	"CREATE_JAVA_SERVER_EXECUTE":        "Execute:\n%s",
	"CREATE_BEDROCK_EXTRACTING":         "Extracting Bedrock Server...",
	"CREATE_BEDROCK_EXTRACT_ERROR":      "Error extracting:",
	"CREATE_BEDROCK_PACKS_PROMPT":       "Enter links to behavior/resource packs (.zip/.mcpack/.mcaddon), separated by space, or press Enter to search for local files:",
	"CREATE_BEDROCK_LOCAL_FILES_PROMPT": "Searching for local files. Place files (.zip/.mcpack/.mcaddon) in the current folder and press Enter when ready.",
	"CREATE_BEDROCK_MOVE_ERROR":         "Error moving file:",
	"CREATE_BEDROCK_READY":              "Bedrock server ready to start! Enter the bedrock-server-1.21.95.1 folder and run bedrock_server.exe",
	"CREATE_API_ERROR":                  "Error creating server",
	"CREATE_SERVER_CREATED":             "Server created! Waiting for DNS propagation...",
	"CREATE_SERVER_ACCESS":              "Access: https://%s.mineserver.theushen.me",
	"INVALID_FILE":                      "Invalid file: %s",
	"FILE_MOVED":                        "Moved: %s -> %s",
	"CREATE_JAVA_SERVER_START_ERROR":    "Error starting Java server: %s",
	"CREATE_JAVA_SERVER_KILL_ERROR":     "Error killing Java process: %s",
	"CREATE_JAVA_SERVER_KILLED":         "Initial Java server process killed for setup safety.",
	"CREATE_JAVA_SERVER_START_INFO":     "Server created! To start your server in the future, use: mineflared start %s",
	"CREATE_BEDROCK_SERVER_START_INFO":  "Bedrock server created! To start your server in the future, use: mineflared start %s",

	// Start command
	"START_SHORT_DESC":          "Start a Minecraft server",
	"START_SERVER_DIR_ERROR":    "Could not open server directory: %s",
	"START_SERVER_TYPE_UNKNOWN": "Could not determine if server is Java or Bedrock.",
	"START_EULA_PROMPT":         "Do you agree to the EULA (https://aka.ms/MinecraftEULA)? (Y/N): ",
	"START_EULA_WRITE_ERROR":    "Failed to write EULA file: %s",
	"START_EULA_ACCEPTED":       "EULA accepted, starting server...",
	"START_EULA_DECLINED":       "You must accept the EULA to start the server.",
	"START_JAVA_STARTING":       "Starting Java server...",
	"START_JAVA_ERROR":          "Java server exited with error: %s",
	"START_BEDROCK_NOT_FOUND":   "Bedrock server directory not found.",
	"START_BEDROCK_STARTING":    "Starting Bedrock server...",
	"START_BEDROCK_ERROR":       "Bedrock server exited with error: %s",
	"START_LOGIN_REQUIRED":      "You must be logged in to use this command.",
	"START_JAVA21_REQUIRED":     "[!!!] JAVA 21 or newer is required to run this server. Download at https://adoptium.net/temurin/releases/?version=21, install, close and reopen your terminal, and try again.",

	// Status command
	"STATUS_SHORT_DESC":     "Check server status",
	"STATUS_LOGIN_REQUIRED": "❌ You need to be logged in first!",
	"STATUS_REQUEST_ERROR":  "❌ Error creating request:",
	"STATUS_QUERY_ERROR":    "❌ Error querying status:",
	"STATUS_DECODE_ERROR":   "❌ Error decoding response:",
	"STATUS_API_ERROR":      "❌ Error: %v",
	"STATUS_UNKNOWN_ERROR":  "❌ Unknown error. HTTP Code: %d",
	"STATUS_SERVER_STATUS":  "✅ Server status:",
	"STATUS_STATUS_LINE":    "   • Status: %s",
	"STATUS_MESSAGE_LINE":   "   • Message: %s",

	// DNS update
	"DNS_IP_ERROR":       "Could not get current IP.",
	"DNS_UPDATE_ERROR":   "Error updating DNS:",
	"DNS_UPDATE_SUCCESS": "DNS updated successfully.",
	"DNS_NO_UPDATE":      "IP hasn't changed, DNS doesn't need updating.",

	// Language command
	"LANGUAGE_SHORT_DESC":     "Change the interface language",
	"LANGUAGE_CURRENT":        "Current language: %s",
	"LANGUAGE_PROMPT":         "Choose language:\n[1] English\n[2] Portuguese",
	"LANGUAGE_OPTION_PROMPT":  "Option: ",
	"LANGUAGE_INVALID_OPTION": "Invalid option. Please choose 1 for English or 2 for Portuguese.",
	"LANGUAGE_CHANGED":        "Language changed to: %s",

	// First run language selection
	"FIRST_RUN_LANGUAGE":   "Welcome to mineflared-cli!\nPlease select your preferred language:",
	"FIRST_RUN_ENGLISH":    "[1] English",
	"FIRST_RUN_PORTUGUESE": "[2] Portuguese (Português)",
}

var PortugueseTranslations = map[string]string{
	// Root command
	"CLI_SHORT_DESC": "CLI para criar e gerenciar servidores de Minecraft",

	// Login command
	"LOGIN_SHORT_DESC":         "Faz login via GitHub",
	"LOGIN_IP_ERROR":           "Não foi possível obter IP. Tente manualmente.",
	"LOGIN_BROWSER_OPENING":    "Abrindo navegador para login com GitHub...",
	"LOGIN_TOKEN_ERROR":        "Erro ao receber token. Tente novamente.",
	"LOGIN_SUCCESS":            "✅ Login realizado com sucesso!",
	"LOGIN_PORT_ERROR":         "Erro ao abrir porta 4500:",
	"LOGIN_TOKEN_NOT_RECEIVED": "Token não recebido",
	"LOGIN_BROWSER_SUCCESS":    "✅ Login realizado com sucesso. Você pode fechar esta janela.",
	"LOGIN_WAITING":            "Aguardando login...",

	//Delete command
	"DELETE_SHORT_DESC":     "Apaga sua conta e todos os dados associados do Mineflared.",
	"DELETE_CONFIRM_PROMPT": "Tem certeza que deseja apagar permanentemente sua conta e todos os dados? (y/N): ",
	"DELETE_ABORTED":        "Operação cancelada.",
	"DELETE_LOGIN_REQUIRED": "Você precisa estar logado para apagar sua conta.",
	"DELETE_REQUEST_FAILED": "Falha ao conectar ao servidor:",
	"DELETE_SUCCESS":        "Conta apagada com sucesso.",
	"DELETE_ERROR":          "Erro ao apagar a conta:",

	// List command
	"LIST_SHORT_DESC":        "Lista todos os servidores locais",
	"LIST_SERVERS_DIR_ERROR": "Não foi possível abrir o diretório de servidores: %s",
	"LIST_HEADER":            "Servidores (proprietário atual):",

	// Restore command
	"RESTORE_SHORT_DESC":     "Restaurar um servidor do backup",
	"RESTORE_CHOOSE_PATH":    "Escolha um arquivo de backup para restaurar:",
	"RESTORE_CURRENT_DIR":    "Diretório atual:",
	"RESTORE_FILE_PROMPT":    "Digite o nome do arquivo para restaurar (com .zip ou .rar): ",
	"RESTORE_FILE_EMPTY":     "Nenhum nome de arquivo informado.",
	"RESTORE_FILE_NOT_FOUND": "Arquivo de backup não encontrado.",
	"RESTORE_SERVER_EXISTS":  "Já existe um servidor com esse nome. Apague-o antes de restaurar.",
	"RESTORE_UNZIPPING":      "Extraindo ZIP...",
	"RESTORE_UNZIP_ERROR":    "Erro ao extrair ZIP: %s\n",
	"RESTORE_UNRAR":          "Extraindo RAR (requer 'unrar' instalado)...",
	"RESTORE_UNRAR_ERROR":    "Erro ao extrair RAR: %s\n",
	"RESTORE_INVALID_FORMAT": "Formato de arquivo não suportado.",
	"RESTORE_DONE":           "Restauração concluída em: %s",

	// Create command
	"CREATE_SHORT_DESC":                 "Cria um novo servidor Minecraft",
	"CREATE_LOGIN_REQUIRED":             "Você precisa estar logado. Use 'mineflared-cli login'",
	"CREATE_SERVER_NAME_PROMPT":         "Digite o nome do servidor: ",
	"CREATE_SERVER_NAME_EMPTY":          "Nome do servidor não pode ser vazio!",
	"CREATE_DIR_ERROR":                  "Erro ao criar diretório do servidor:",
	"CREATE_SERVER_TYPE_PROMPT":         "Escolha o tipo de servidor:",
	"CREATE_JAVA_OPTION":                "[1] Java",
	"CREATE_BEDROCK_OPTION":             "[2] Bedrock",
	"CREATE_OPTION_PROMPT":              "Opção: ",
	"CREATE_JAVA_VERSION_PROMPT":        "Escolha a versão do Java:",
	"CREATE_JAVA_VERSION_OPTION":        "[%d] %s (Mods: %v, Plugins: %v)",
	"CREATE_INVALID_OPTION":             "Opção inválida.",
	"CREATE_BEDROCK_WARNING":            "ATENÇÃO: O servidor Bedrock só suporta Windows por enquanto.",
	"CREATE_BEDROCK_CANCEL":             "Você pode cancelar agora (digite 'n' para cancelar) ou pressione Enter para continuar.",
	"CREATE_ACTION_CANCELLED":           "Ação cancelada.",
	"CREATE_DOWNLOADING":                "Baixando arquivo do servidor...",
	"CREATE_DOWNLOAD_ERROR":             "Erro ao baixar o arquivo:",
	"CREATE_DOWNLOAD_COMPLETE":          "Arquivo baixado em:",
	"CREATE_MODS_SUPPORT":               "Este servidor suporta MODS. Insira os links diretos para os MODs (.jar, .zip) ou pressione Enter para buscar arquivos localmente:",
	"CREATE_PLUGINS_SUPPORT":            "Este servidor suporta Plugins. Insira os links diretos para os Plugins (.jar, .zip) ou pressione Enter para buscar arquivos localmente:",
	"CREATE_NO_MODS_PLUGINS":            "Este servidor NÃO suporta mods/plugins.",
	"CREATE_DOWNLOADING_FILE":           "Baixando:",
	"CREATE_DOWNLOAD_FILE_ERROR":        "Erro ao baixar %s: %v",
	"CREATE_LOCAL_FILES_PROMPT":         "Buscando arquivos localmente. Coloque os arquivos (.jar/.zip) na pasta atual e pressione Enter quando pronto.",
	"CREATE_READ_DIR_ERROR":             "Erro ao ler diretório atual:",
	"CREATE_MOVE_FILES_ERROR":           "Erro ao mover arquivos:",
	"CREATE_JAVA_SERVER_STARTING":       "Iniciando servidor Java...",
	"CREATE_JAVA_SERVER_EXECUTE":        "Execute:\n%s",
	"CREATE_BEDROCK_EXTRACTING":         "Descompactando Bedrock Server...",
	"CREATE_BEDROCK_EXTRACT_ERROR":      "Erro ao descompactar:",
	"CREATE_BEDROCK_PACKS_PROMPT":       "Insira os links para behavior/resource packs (.zip/.mcpack/.mcaddon), separados por espaço, ou pressione Enter para buscar arquivos localmente:",
	"CREATE_BEDROCK_LOCAL_FILES_PROMPT": "Buscando arquivos localmente. Coloque os arquivos (.zip/.mcpack/.mcaddon) na pasta atual e pressione Enter quando pronto.",
	"CREATE_BEDROCK_MOVE_ERROR":         "Erro ao mover arquivo:",
	"CREATE_BEDROCK_READY":              "Servidor Bedrock pronto para iniciar! Entre na pasta bedrock-server-1.21.95.1 e execute bedrock_server.exe",
	"CREATE_API_ERROR":                  "Erro ao criar servidor",
	"CREATE_SERVER_CREATED":             "Servidor criado! Aguarde a propagação do DNS...",
	"CREATE_SERVER_ACCESS":              "Acesse: https://%s.mineserver.theushen.me",
	"INVALID_FILE":                      "Arquivo inválido: %s",
	"FILE_MOVED":                        "Movido: %s -> %s",
	"CREATE_JAVA_SERVER_START_ERROR":    "Erro ao iniciar o servidor Java: %s",
	"CREATE_JAVA_SERVER_KILL_ERROR":     "Erro ao finalizar o processo Java: %s",
	"CREATE_JAVA_SERVER_KILLED":         "Processo Java inicial foi finalizado por segurança.",
	"CREATE_JAVA_SERVER_START_INFO":     "Servidor criado! Para iniciar seu servidor no futuro, use: mineflared start %s",
	"CREATE_BEDROCK_SERVER_START_INFO":  "Servidor Bedrock criado! Para iniciar seu servidor no futuro, use: mineflared start %s",

	// Backup command
	"BACKUP_SHORT_DESC":       "Criar um backup do seu servidor",
	"BACKUP_SERVER_NOT_FOUND": "Servidor não encontrado.",
	"BACKUP_FORMAT_PROMPT":    "Escolha o formato do backup (zip/rar): ",
	"BACKUP_INVALID_FORMAT":   "Formato inválido. Use apenas zip ou rar.",
	"BACKUP_CHOOSE_PATH":      "Escolha onde salvar o arquivo de backup.",
	"BACKUP_CURRENT_DIR":      "Diretório atual:",
	"BACKUP_PATH_PROMPT":      "Digite o caminho da pasta para salvar (deixe em branco para o atual): ",
	"BACKUP_CREATING_ZIP":     "Criando backup ZIP...",
	"BACKUP_CREATING_RAR":     "Criando backup RAR (requer 'rar' instalado)...",
	"BACKUP_ZIP_ERROR":        "Erro ao criar ZIP: %s\n",
	"BACKUP_RAR_ERROR":        "Erro ao criar RAR: %s\n",
	"BACKUP_DONE":             "Backup criado em: %s",

	// Config command
	"CONFIG_RUNNING_BLOCK":    "Não é possível configurar enquanto o servidor está rodando.",
	"START_CONFIGURING_BLOCK": "Não é possível iniciar o servidor enquanto o painel de configuração está aberto.",
	"CONFIG_SHORT_DESC":       "Abrir o painel web de configuração do servidor",
	"CONFIG_OPEN_BROWSER":     "Abrindo painel de configuração em: %s",

	// Start command
	"START_SHORT_DESC":          "Inicia um servidor Minecraft",
	"START_SERVER_DIR_ERROR":    "Não foi possível abrir o diretório do servidor: %s",
	"START_SERVER_TYPE_UNKNOWN": "Não foi possível determinar se o servidor é Java ou Bedrock.",
	"START_EULA_PROMPT":         "Você concorda com o EULA (https://aka.ms/MinecraftEULA)? (S/N): ",
	"START_EULA_WRITE_ERROR":    "Falha ao escrever o arquivo EULA: %s",
	"START_EULA_ACCEPTED":       "EULA aceita, iniciando servidor...",
	"START_EULA_DECLINED":       "Você precisa aceitar o EULA para iniciar o servidor.",
	"START_JAVA_STARTING":       "Iniciando servidor Java...",
	"START_JAVA_ERROR":          "Servidor Java finalizado com erro: %s",
	"START_BEDROCK_NOT_FOUND":   "Diretório do servidor Bedrock não encontrado.",
	"START_BEDROCK_STARTING":    "Iniciando servidor Bedrock...",
	"START_BEDROCK_ERROR":       "Servidor Bedrock finalizado com erro: %s",
	"START_LOGIN_REQUIRED":      "Você precisa estar logado para usar este comando.",
	"START_JAVA21_REQUIRED":     "[!!!] É necessário ter o JAVA 21 ou superior para rodar este servidor. Baixe em https://adoptium.net/temurin/releases/?version=21, instale, feche e reabra o terminal e tente novamente.",

	// Status command
	"STATUS_SHORT_DESC":     "Verifica o status do servidor",
	"STATUS_LOGIN_REQUIRED": "❌ Você precisa estar logado primeiro!",
	"STATUS_REQUEST_ERROR":  "❌ Erro ao criar requisição:",
	"STATUS_QUERY_ERROR":    "❌ Erro ao consultar status:",
	"STATUS_DECODE_ERROR":   "❌ Erro ao decodificar resposta:",
	"STATUS_API_ERROR":      "❌ Erro: %v",
	"STATUS_UNKNOWN_ERROR":  "❌ Erro desconhecido. Código HTTP: %d",
	"STATUS_SERVER_STATUS":  "✅ Status do servidor:",
	"STATUS_STATUS_LINE":    "   • Status: %s",
	"STATUS_MESSAGE_LINE":   "   • Mensagem: %s",

	// DNS update
	"DNS_IP_ERROR":       "Não foi possível obter o IP atual.",
	"DNS_UPDATE_ERROR":   "Erro ao atualizar DNS:",
	"DNS_UPDATE_SUCCESS": "DNS atualizado com sucesso.",
	"DNS_NO_UPDATE":      "IP não mudou, DNS não precisa de atualização.",

	// Language command
	"LANGUAGE_SHORT_DESC":     "Altera o idioma da interface",
	"LANGUAGE_CURRENT":        "Idioma atual: %s",
	"LANGUAGE_PROMPT":         "Escolha o idioma:\n[1] Inglês\n[2] Português",
	"LANGUAGE_OPTION_PROMPT":  "Opção: ",
	"LANGUAGE_INVALID_OPTION": "Opção inválida. Por favor, escolha 1 para Inglês ou 2 para Português.",
	"LANGUAGE_CHANGED":        "Idioma alterado para: %s",

	// First run language selection
	"FIRST_RUN_LANGUAGE":   "Bem-vindo ao mineflared-cli!\nPor favor, selecione seu idioma preferido:",
	"FIRST_RUN_ENGLISH":    "[1] Inglês (English)",
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
