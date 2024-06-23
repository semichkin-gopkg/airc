package cmd

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/semichkin-gopkg/airc/internal/templates"
	"github.com/semichkin-gopkg/env"
)

type BuildVariables struct {
	Root   string `env:"AIRC_ROOT,expand" envDefault:"/app"`
	TmpDir string `env:"AIRC_TMP_DIR,expand" envDefault:"../apptmp"`

	Bin              string        `env:"AIRC_BIN,expand" envDefault:"$AIRC_TMP_DIR/main"`
	Src              string        `env:"AIRC_SRC,expand" envDefault:"$AIRC_ROOT/cmd/app/main.go"`
	Cmd              string        `env:"AIRC_CMD,expand" envDefault:"cd $AIRC_ROOT && go build -o $AIRC_BIN $AIRC_SRC"`
	Delay            uint          `env:"AIRC_DELAY,expand" envDefault:"1000"`
	ExcludeDir       []string      `env:"AIRC_EXCLUDE_DIR,expand" envDefault:"assets,tmp,vendor"`
	ExcludeFile      []string      `env:"AIRC_EXCLUDE_FILE,expand" envDefault:""`
	ExcludeRegex     []string      `env:"AIRC_EXCLUDE_REGEX,expand" envDefault:""`
	ExcludeUnchanged bool          `env:"AIRC_EXCLUDE_UNCHANGED,expand" envDefault:"false"`
	FollowSymlink    bool          `env:"AIRC_FOLLOW_SYMLINK,expand" envDefault:"false"`
	FullBin          string        `env:"AIRC_FULL_BIN,expand" envDefault:""`
	IncludeDir       []string      `env:"AIRC_INCLUDE_DIR,expand" envDefault:""`
	IncludeExt       []string      `env:"AIRC_INCLUDE_EXT,expand" envDefault:"go,tpl,tmpl,html"`
	KillDelay        time.Duration `env:"AIRC_KILL_DELAY,expand" envDefault:"1s"`
	Log              string        `env:"AIRC_LOG,expand" envDefault:"$AIRC_TMP_DIR/build-errors.log"`
	SendInterrupt    bool          `env:"AIRC_SEND_INTERRUPT,expand" envDefault:"true"`
	StopOnError      bool          `env:"AIRC_STOP_ON_ERROR,expand" envDefault:"false"`

	ColorApp     string `env:"AIRC_COLOR_APP" envDefault:"" envExpand:"true"`
	ColorBuild   string `env:"AIRC_COLOR_BUILD" envDefault:"yellow" envExpand:"true"`
	ColorMain    string `env:"AIRC_COLOR_MAIN" envDefault:"magenta" envExpand:"true"`
	ColorRunner  string `env:"AIRC_COLOR_RUNNER" envDefault:"green" envExpand:"true"`
	ColorWatcher string `env:"AIRC_COLOR_WATCHER" envDefault:"cyan" envExpand:"true"`

	LogTime bool `env:"AIRC_LOG_TIME" envDefault:"false" envExpand:"true"`

	MiscCleanOnExit bool `env:"AIRC_MISC_CLEAN_ON_EXIT" envDefault:"false" envExpand:"true"`
}

func build(_ context.Context, configPath string) error {
	vars, err := env.Fill[BuildVariables](
		env.WithOnSetFn(func(tag string, value interface{}, isDefault bool) {
			if err := os.Setenv(tag, fmt.Sprintf("%v", value)); err != nil {
				log.Println("Warning: failed to set env variable:", err)
			}
		}))
	if err != nil {
		return err
	}

	tmpl, err := template.New("AirTomlTemplate").Funcs(template.FuncMap{
		"noescape": noescape,
	}).Parse(templates.AIRTomlTemplate)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	if err = tmpl.Execute(buf, vars); err != nil {
		return err
	}

	if err = os.MkdirAll(filepath.Dir(configPath), os.ModePerm); err != nil {
		return err
	}

	return os.WriteFile(configPath, buf.Bytes(), os.ModePerm)
}

func noescape(str string) template.HTML {
	return template.HTML(str)
}
