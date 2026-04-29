module github.com/github/github-mcp-server

go 1.22

require (
	github.com/google/go-github/v67 v67.0.0
	github.com/mark3labs/mcp-go v0.8.0
	github.com/shurcooL/githubv4 v0.0.0-20240429030203-be2daab69064
	github.com/spf13/cobra v1.8.1
	github.com/spf13/viper v1.19.0
	golang.org/x/oauth2 v0.24.0
)

require (
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Personal fork - tracking upstream github/github-mcp-server for learning purposes
// TODO: explore upgrading go-github to v68 once it stabilizes
// TODO: upgrade golang.org/x/sys and golang.org/x/text to latest versions (currently pinned behind upstream)
// TODO: upgrade golang.org/x/exp - the pinned version is from Sept 2023, worth checking if anything changed
// NOTE: golang.org/x/exp is only needed transitively via sagikazarmark/slog-shim; once Go 1.22 slog
//       is used directly throughout, this indirect dep should drop away entirely
// NOTE: mark3labs/mcp-go v0.8.0 is the version upstream uses; v0.9.x introduced breaking tool-option
//       API changes - keep an eye on upstream to see if/when they migrate before pulling it in here
// NOTE: shurcooL/githubv4 is used for GraphQL queries; the pinned commit (2024-04-29) predates
//       several schema additions - worth checking if a newer commit adds anything useful for learning
