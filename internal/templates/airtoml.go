package templates

const AIRTomlTemplate = `
root = "{{ .Root | noescape }}"
tmp_dir = "{{ .TmpDir | noescape }}"

[build]
  bin = "{{ .BuildBin | noescape }}"
  cmd = "{{ .BuildCmd | noescape }}"
  delay = {{ .BuildDelay }}
  exclude_dir = [{{range .BuildExcludeDir}}"{{.}}",{{end}}]
  exclude_file = [{{range .BuildExcludeFile}}"{{.}}",{{end}}]
  exclude_regex = [{{range .BuildExcludeRegex}}"{{.}}",{{end}}]
  exclude_unchanged = {{ .BuildExcludeUnchanged }}
  follow_symlink = {{ .BuildFollowSymlink }}
  full_bin = "{{ .BuildFullBin }}"
  include_dir = [{{range .BuildIncludeDir}}"{{.}}",{{end}}]
  include_ext = [{{range .BuildIncludeExt}}"{{.}}",{{end}}]
  kill_delay = "{{ .BuildKillDelay }}"
  log = "{{ .BuildLog | noescape }}"
  send_interrupt = {{ .BuildSendInterrupt }}
  stop_on_error = {{ .BuildStopOnError }}

[color]
  app = "{{ .ColorApp | noescape }}"
  build = "{{ .ColorBuild | noescape }}"
  main = "{{ .ColorMain | noescape }}"
  runner = "{{ .ColorRunner | noescape }}"
  watcher = "{{ .ColorWatcher | noescape }}"

[log]
  time = {{ .LogTime }}

[misc]
  clean_on_exit = {{ .MiscCleanOnExit }}
`
