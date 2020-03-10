package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func BindataRestoreAssetsWithTemplates(dir, name string, templateData interface{}) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return BindataRestoreAssetWithTemplate(dir, name, templateData)
	}
	// Dir
	for _, child := range children {
		err = BindataRestoreAssetsWithTemplates(dir, filepath.Join(name, child), templateData)
		if err != nil {
			return err
		}
	}
	return nil
}

func BindataRestoreAssetWithTemplate(dir, name string, templateData interface{}) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}

	name = strings.Replace(name, "templates/", "", -1)
	name = strings.Replace(name, "templates\\", "", -1)
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), BindataFillTemplate(data, templateData), info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

func BindataFillTemplate(data []byte, templateData interface{}) []byte {
	if templateData == nil {
		return data
	}

	t := template.New("")

	t, err := t.Parse(string(data))
	if err != nil {
		log.Fatal(err)
	}

	var outputString bytes.Buffer
	err = t.Execute(&outputString, templateData)
	if err != nil {
		log.Fatal(err)
	}

	return outputString.Bytes()
}
