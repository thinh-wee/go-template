package app

import "os"

var (
	Binary          string = "PROGRAM"
	Language        string = "golang"
	LanguageVersion string = "go1.18.4"
	Repository      string = "local"
	GitBranch       string = "unknow"
	GitRevision     string = "unknow"
	Version         string = "develop"
	Os              string = "windows"
	Arch            string = "amd64"
	Build           string = "dev@vscode"
	BuildDate       string = "date -u +\"%Y-%m-%dT%H:%M:%SZ\""

	ENVIRON     string = "DEFAULT (DEVELOP)"
	CONFIG_PATH string = "./config/setting.yaml"
)

func init() {
	// alow commands
	if err := allow_arguments(
		"--help", "-h",
		"version", "--version", "-v",
		"start", "--start-env", "--environ", "--config",
	); err != nil {
		println(err.Error())
		os.Exit(0)
	}
	// validate help command
	if is_help_command() {
		println("\r\n[HELP]")
		println("-----------------------------------------------------------------")
		println(Binary, "[OPTIONS] COMMAND")
		println("-----------------------------------------------------------------")
		println("\r\n--environ=[\"DEVELOP\",\"QC\",\"PRODUCTION\",...] start :")
		println("    Start the program with the imported environment. ")
		println("\r\n--config={PATH} start :")
		println("    Start the program with config from input path.")
		println("")
		println("-----------------------------------------------------------------")
		println(Binary, "OPTION")
		println("-----------------------------------------------------------------")
		println("\r\n-h, --help :")
		println("    Prints the synopsis and a list of the most commonly used commands.")
		println("\r\n-v, --version or version :")
		println("    Prints the git suite version that the git program came from.")
		println("\r\n--start-env=[\"DEVELOP\",\"QC\",\"PRODUCTION\",...] :")
		println("    Start the program with the imported environment.")
		println("    This option is equivalent to \"", Binary, "--environ={ENV} start\".")

		os.Exit(0)
	}
	// validate get version command
	if is_get_version_command() {
		println("\r\n[VERSION INFO]")
		println("-----------------------------------------------------------------")
		println("BINARY:", Binary)
		println("LANGUAGE:", Language)
		println("LANGUAGE_VERSION:", LanguageVersion)
		println("OS:", Os)
		println("ARCH:", Arch)
		println("REPO:", Repository)
		println("VERSION:", Version, "( Branch:", GitBranch, ", Revision:", GitRevision, ")")
		println("BUID_BY:", Build)
		println("BUID_DATE:", BuildDate)
		os.Exit(0)
	}

	// validate start command
	if is_start_command() {
		println("\r\n[START] with OPTIONS")
		println("-----------------------------------------------------------------")
		println("ENVIRON=", ENVIRON)
		println("CONFIG_PATH=", CONFIG_PATH)
		println("-----------------------------------------------------------------")
	}
}

func is_help_command() bool {
	help_value, ok, err := get_argument("--help", "-h")
	if err != nil {
		println(err.Error())
		os.Exit(0)
	}
	if ok {
		if len(help_value) > 0 {
			println("The --help argument need not have a value.")
			os.Exit(0)
		}
		return true
	}
	return false
}

func is_get_version_command() bool {
	version_value, ok, err := get_argument("version", "--version", "-v")
	if err != nil {
		println(err.Error())
		os.Exit(0)
	}
	if ok {
		if len(version_value) > 0 {
			println("The --help argument need not have a value.")
			os.Exit(0)
		}
		return true
	}
	return false
}

func is_start_command() (set bool) {
	// set ENVIRON with arg --environ
	environ_value, ok, err := get_argument("--environ")
	if err != nil {
		println(err.Error())
		os.Exit(0)
	}
	if ok {
		ENVIRON = environ_value
		set = true
	}
	// set CONFIG_PATH with arg --config
	config_path_value, ok, err := get_argument("--config")
	if err != nil {
		println(err.Error())
		os.Exit(0)
	}
	if ok {
		CONFIG_PATH = config_path_value
		set = true
	}
	// apply command "start"
	start_env_value, ok, err := get_argument("--start-env", "start")
	if err != nil {
		println(err.Error())
		os.Exit(0)
	}
	if !ok {
		println("The arguments is \"--start-env={ENV}\" or \"start\" command is required to launch.")
		os.Exit(1)
	}
	// replace ENVIRON with arg --start-env
	if len(start_env_value) > 0 {
		ENVIRON = start_env_value
		set = true
	}
	return set
}
