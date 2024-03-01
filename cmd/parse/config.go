package parse

import (
	"os"
	"path"
	"path/filepath"

	"github.com/opentdp/go-helper/filer"
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/strutil"
	"gopkg.in/yaml.v3"

	"tdp-cloud/cmd/args"
)

// 配置管理

var YamlFile string

func readYaml(c any) error {

	if YamlFile == "" {
		return nil
	}
	if !filer.Exists(YamlFile) {
		return nil
	}

	bytes, err := os.ReadFile(YamlFile)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(bytes, c)

}

func saveYaml(c any) error {

	bytes, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile(YamlFile, bytes, 0644)

}

// 初始化存储目录

func initDataset() {

	if args.Dataset.Secret == "" {
		args.Dataset.Secret = strutil.Rand(32)
	}

	if args.Dataset.Dir != "" && args.Dataset.Dir != "." {
		os.MkdirAll(args.Dataset.Dir, 0755)
	}

}

// 初始化日志能力

func initLogger() {

	config := &logman.Config{
		Level:    args.Logger.Level,
		Target:   args.Logger.Target,
		Storage:  args.Logger.Dir,
		Filename: "server",
	}

	if !filepath.IsAbs(config.Storage) {
		config.Storage = path.Join(args.Dataset.Dir, config.Storage)
	}

	logman.SetDefault(config)

}
