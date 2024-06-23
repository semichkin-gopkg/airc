package templates

const AIRTomlTemplate = `
root = "{{ .Root | noescape }}"
tmp_dir = "{{ .TmpDir | noescape }}"

[build]
  bin = "{{.Bin | noescape}}"
  cmd = "{{.Cmd | noescape}}"
  delay = {{.Delay}}
  exclude_dir = [{{range .ExcludeDir}}"{{.}}",{{end}}]
  exclude_file = [{{range .ExcludeFile}}"{{.}}",{{end}}]
  exclude_regex = [{{range .ExcludeRegex}}"{{.}}",{{end}}]
  exclude_unchanged = {{.ExcludeUnchanged}}
  follow_symlink = {{.FollowSymlink}}
  full_bin = "{{.FullBin}}"
  include_dir = [{{range .IncludeDir}}"{{.}}",{{end}}]
  include_ext = [{{range .IncludeExt}}"{{.}}",{{end}}]
  kill_delay = "{{.KillDelay}}"
  log = "{{.Log | noescape}}"
  send_interrupt = {{.SendInterrupt}}
  stop_on_error = {{.StopOnError}}

[color]
  app = "{{.ColorApp | noescape}}"
  build = "{{.ColorBuild | noescape}}"
  main = "{{.ColorMain | noescape}}"
  runner = "{{.ColorRunner | noescape}}"
  watcher = "{{.ColorWatcher | noescape}}"

[log]
  time = {{.LogTime}}

[misc]
  clean_on_exit = {{.MiscCleanOnExit}}
`
