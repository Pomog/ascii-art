package configuration

type Config struct {
	ResultsFileName            string
	FileNameWithSymbolsDefault string
	LettersToBeColored         string
	ValidColors                []string
	ValidAligns                []string
	ValidAsciiArtSourse        []string
	Banner                     string
	InputString                string
}

func initConfig() *Config {
	return &Config{
		ResultsFileName:            "result.txt",
		FileNameWithSymbolsDefault: "standard.txt",
		ValidColors:                []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet", "white"},
		ValidAligns:                []string{"left", "center", "right", "justify"},
		ValidAsciiArtSourse:        []string{"standard", "shadow", "thinkertoy"},
		Banner:                     "standard",
		InputString:                "",
	}
}

var ConfigInstance = initConfig()
