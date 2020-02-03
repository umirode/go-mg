package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"strings"
	"text/template"
)

type config struct {
	output  string
	name    string
	network string
	address string
}

type templates struct {
	gitignore *asset
	modules   *asset
	goMod     *asset
	goSum     *asset
	mainGo    *asset
	makefile  *asset
	readmeMd  *asset
}

//go:generate go-bindata templates

func main() {
	config := getConfig()
	templates := getTemplates()

	outputPath := createOutputDirectory(config.output, config.name)

	saveTemplate(outputPath, templates.gitignore, nil)
	saveTemplate(outputPath, templates.modules, nil)
	saveTemplate(outputPath, templates.goSum, nil)
	saveTemplate(outputPath, templates.makefile, nil)

	saveTemplate(outputPath, templates.readmeMd, struct {
		Name string
	}{
		Name: config.name,
	})
	saveTemplate(outputPath, templates.goMod, struct {
		Name string
	}{
		Name: config.name,
	})
	saveTemplate(outputPath, templates.mainGo, struct {
		Network string
		Address string
	}{
		Network: config.network,
		Address: config.address,
	})
}

func saveTemplate(output string, template *asset, data interface{}) {
	templateFileName := strings.Replace(template.info.Name(), "templates/", "", -1)
	templateString := string(template.bytes)

	if data == nil {
		saveFile(output, templateFileName, templateString)
	} else {
		saveFile(output, templateFileName, fillTemplate(templateString, data))
	}
}

func createOutputDirectory(outputPath string, name string) string {
	if string(outputPath[len(outputPath)-1:]) != "/" {
		outputPath += "/"
	}
	outputPath += name

	err := os.MkdirAll(outputPath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	return outputPath
}

func saveFile(outputPath string, name string, text string) {
	file, err := os.Create(outputPath + "/" + name)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write([]byte(text))
	if err != nil {
		log.Fatal(err)
	}
}

func fillTemplate(text string, data interface{}) string {
	t := template.New("")

	t, err := t.Parse(text)
	if err != nil {
		log.Fatal(err)
	}

	var outputString bytes.Buffer
	err = t.Execute(&outputString, data)
	if err != nil {
		log.Fatal(err)
	}

	return outputString.String()
}

func getTemplates() *templates {
	templates := &templates{}

	templateReadme, err := templatesReadmeMd()
	if err != nil {
		log.Fatal(err)
	}
	templates.readmeMd = templateReadme

	templateGitignore, err := templatesGitignore()
	if err != nil {
		log.Fatal(err)
	}
	templates.gitignore = templateGitignore

	templateModules, err := templatesModules()
	if err != nil {
		log.Fatal(err)
	}
	templates.modules = templateModules

	templateGoMod, err := templatesGoMod()
	if err != nil {
		log.Fatal(err)
	}
	templates.goMod = templateGoMod

	templateGoSum, err := templatesGoSum()
	if err != nil {
		log.Fatal(err)
	}
	templates.goSum = templateGoSum

	templateMainGo, err := templatesMainGo()
	if err != nil {
		log.Fatal(err)
	}
	templates.mainGo = templateMainGo

	templateMakefile, err := templatesMakefile()
	if err != nil {
		log.Fatal(err)
	}
	templates.makefile = templateMakefile

	return templates
}

func getConfig() *config {
	config := &config{}

	flag.StringVar(&config.output, "output", "", "Output directory path.")
	flag.StringVar(&config.name, "name", "", "Microservice name.")
	flag.StringVar(&config.network, "network", "", "Network type (tcp, tcp4, tcp6, unix or unixpacket).")
	flag.StringVar(&config.address, "address", "", "Address for tcp (\":56001\", \"127.0.0.1\", \"127.0.0.1:56001\") or path for unix socket (example/path/microservice.socket).")

	flag.Parse()

	if config.output == "" {
		log.Fatal("Output is required. Use -help flag for info.")
	}
	if config.name == "" {
		log.Fatal("Name is required. Use -help flag for info.")
	}
	if config.network == "" {
		log.Fatal("Network is required. Use -help flag for info.")
	}
	if config.address == "" {
		log.Fatal("Address is required. Use -help flag for info.")
	}

	if config.network != "tcp" &&
		config.network != "tcp4" &&
		config.network != "tcp6" &&
		config.network != "unix" &&
		config.network != "unixpacket" {
		log.Fatal("Unexpected network. Use -help flag for info.")
	}

	return config
}
