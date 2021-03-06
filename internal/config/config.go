// Copyright 2017 Debpkg authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package config

import (
	"fmt"
	"runtime"

	"gopkg.in/yaml.v2"
)

// PkgSpecFile represents a single debian package
type PkgSpecFile struct {
	Name            string `yaml:"name"`
	Version         string `yaml:"version"`
	Architecture    string `yaml:"architecture"`
	Maintainer      string `yaml:"maintainer"`
	MaintainerEmail string `yaml:"maintainer_email"`
	Homepage        string `yaml:"homepage"`
	Section         string `yaml:"section"`
	Depends         string `yaml:"depends"`
	Recommends      string `yaml:"recommends"`
	Suggests        string `yaml:"suggests"`
	Conflicts       string `yaml:"conflicts"`
	Provides        string `yaml:"provides"`
	Replaces        string `yaml:"replaces"`
	Priority        string `yaml:"priority"`
	BuiltUsing      string `yaml:"built_using"`
	Description     struct {
		Short string `yaml:"short"`
		Long  string `yaml:"long"`
	}
	Files []struct {
		Src     string `yaml:"file"`
		Dest    string `yaml:"dest"`
		Content string `yaml:"content"`
	} `yaml:",flow"`
	Directories      []string `yaml:",flow"`
	EmptyDirectories []string `yaml:"emptydirs,flow"`
	ControlExtra     struct {
		Preinst  string `yaml:"preinst"`
		Postinst string `yaml:"postinst"`
		Prerm    string `yaml:"prerm"`
		Postrm   string `yaml:"postrm"`
	} `yaml:"control_extra"`
}

// PkgSpecFileUnmarshal loads the configuration data into a PkgSpecFile structure
func PkgSpecFileUnmarshal(data []byte) (*PkgSpecFile, error) {
	cfg := &PkgSpecFile{
		Name:            "unknown",
		Version:         "0.1.0+dev",
		Architecture:    "any",
		Maintainer:      "anonymous",
		MaintainerEmail: "anon@foo.bar",
		Homepage:        "https://www.google.com",
		Section:         "misc",
		Priority:        "optional",
		BuiltUsing:      runtime.Version(),
	}
	cfg.Description.Long = "-"
	cfg.Description.Short = "-"

	err := yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("problem unmarshaling config file: %v", err)
	}

	return cfg, nil
}
