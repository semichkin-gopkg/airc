package cmd

import (
	"bytes"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/semichkin-gopkg/airc/internal/templates"
	"html/template"
	"os"
	"time"
)

type BuildVariables struct {
	Root   string `env:"AIRC_ROOT" envDefault:"/app" envExpand:"true"`
	TmpDir string `env:"AIRC_TMP_DIR" envDefault:"/apptmp"`

	BuildBin              string        `env:"AIRC_BUILD_BIN" envDefault:"$AIRC_ROOT/main" envExpand:"true"`
	BuildCmdPathToSource  string        `env:"AIRC_BUILD_CMD_PATH_TO_SOURCE" envDefault:"$AIRC_ROOT/cmd/app/main.go" envExpand:"true"`
	BuildCmd              string        `env:"AIRC_BUILD_CMD" envDefault:"cd $AIRC_ROOT && go build -o $AIRC_BUILD_BIN $AIRC_BUILD_CMD_PATH_TO_SOURCE" envExpand:"true"`
	BuildDelay            uint          `env:"AIRC_BUILD_DELAY" envDefault:"1000" envExpand:"true"`
	BuildExcludeDir       []string      `env:"AIRC_BUILD_EXCLUDE_DIR" envDefault:"assets,tmp,vendor" envExpand:"true"`
	BuildExcludeFile      []string      `env:"AIRC_BUILD_EXCLUDE_FILE" envDefault:"" envExpand:"true"`
	BuildExcludeRegex     []string      `env:"AIRC_BUILD_EXCLUDE_REGEX" envDefault:"" envExpand:"true"`
	BuildExcludeUnchanged bool          `env:"AIRC_BUILD_EXCLUDE_UNCHANGED" envDefault:"false" envExpand:"true"`
	BuildFollowSymlink    bool          `env:"AIRC_BUILD_FOLLOW_SYMLINK" envDefault:"false" envExpand:"true"`
	BuildFullBin          string        `env:"AIRC_BUILD_FULL_BIN" envDefault:"" envExpand:"true"`
	BuildIncludeDir       []string      `env:"AIRC_BUILD_INCLUDE_DIR" envDefault:"[]" envExpand:"true"`
	BuildIncludeExt       []string      `env:"AIRC_BUILD_INCLUDE_EXT" envDefault:"go,tpl,tmpl,html" envExpand:"true"`
	BuildKillDelay        time.Duration `env:"AIRC_BUILD_KILL_DELAY" envDefault:"1s" envExpand:"true"`
	BuildLog              string        `env:"AIRC_BUILD_LOG" envDefault:"$AIRC_TMP_DIR/build-errors.log" envExpand:"true"`
	BuildSendInterrupt    bool          `env:"AIRC_BUILD_SEND_INTERRUPT" envDefault:"false" envExpand:"true"`
	BuildStopOnError      bool          `env:"AIRC_BUILD_STOP_ON_ERROR" envDefault:"false" envExpand:"true"`

	ColorApp     string `env:"AIRC_COLOR_APP" envDefault:"" envExpand:"true"`
	ColorBuild   string `env:"AIRC_COLOR_BUILD" envDefault:"yellow" envExpand:"true"`
	ColorMain    string `env:"AIRC_COLOR_MAIN" envDefault:"magenta" envExpand:"true"`
	ColorRunner  string `env:"AIRC_COLOR_RUNNER" envDefault:"green" envExpand:"true"`
	ColorWatcher string `env:"AIRC_COLOR_WATCHER" envDefault:"cyan" envExpand:"true"`

	LogTime bool `env:"AIRC_LOG_TIME" envDefault:"false" envExpand:"true"`

	MiscCleanOnExit bool `env:"AIRC_MISC_CLEAN_ON_EXIT" envDefault:"false" envExpand:"true"`
}

func build(output string) error {
	var vars BuildVariables
	if err := env.Parse(&vars, env.Options{
		OnSet: func(tag string, value interface{}, isDefault bool) {
			_ = os.Setenv(tag, fmt.Sprintf("%v", value))
		},
	}); err != nil {
		return err
	}

	tmpl, err := template.New("AirTomlTemplate").Funcs(template.FuncMap{
		"noescape": noescape,
	}).Parse(templates.AIRTomlTemplate)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, vars); err != nil {
		return err
	}

	return os.WriteFile(output, buf.Bytes(), 0666)
}

func noescape(str string) template.HTML {
	return template.HTML(str)
}
