# fiber-server

A Go Fiber web server with graceful exit, structured logging, environment variable configuration, and local development tooling.

## Local development

### Quick start

See [Prerequisites](#prerequisites) for installing [mise](https://mise.jdx.dev/about.html) –
an all-in-one tool for managing project dependencies, environment variables, and running tasks.

```sh
# Install project dependencies
mise install

# Setup the project
mise run setup

# Start the development server
# By default: https://localhost:3000
mise run
```

### Development scripts

- To see a list of development scripts, run `mise tasks`.
- To run a script, run `mise run <script-name>`.

### VS Code and Cursor

If you're using VS Code or Cursor as your editor, we recommend the following extensions:

- [**TOML Language Support**](https://marketplace.visualstudio.com/items?itemName=be5invis.toml) Add TOML language support for the `mise` config.
- [**Dprint Code Formatter**](https://marketplace.visualstudio.com/items?itemName=dprint.dprint) Use `dprint` as the default formatter for the project.
- [**Go**](https://marketplace.visualstudio.com/items?itemName=golang.go) Golang support for VS Code

We also recommend the following settings in your `.vscode/settings.json`:

```json
{
	"editor.defaultFormatter": "dprint.dprint",
	"editor.formatOnSave": true,
	"[go]": {
		"editor.defaultFormatter": "dprint.dprint"
	},
	"[json]": {
		"editor.defaultFormatter": "dprint.dprint"
	},
	"[jsonc]": {
		"editor.defaultFormatter": "dprint.dprint"
	},
	"[yaml]": {
		"editor.defaultFormatter": "dprint.dprint"
	},
	"[toml]": {
		"editor.defaultFormatter": "dprint.dprint"
	},
	"[markdown]": {
		"editor.defaultFormatter": "dprint.dprint"
	},
	"[dockerfile]": {
		"editor.defaultFormatter": "dprint.dprint"
	}
}
```

### Prerequisites

We use [mise](https://mise.jdx.dev/about.html) to run tasks, manage tool versions,
and manage environment variables.

1. [Install mise](https://mise.jdx.dev/getting-started.html)

```sh
curl https://mise.run | sh
```

2. Add mise to your shell profile. This activates mise in your shell, ensuring the correct tool versions are used for your environment.

```sh
# Zsh
echo 'eval "$(~/.local/bin/mise activate zsh)"' >> ~/.zshrc
echo 'eval "$(~/.local/bin/mise activate zsh --shims)"' >> ~/.zprofile
source ~/.zshrc

# Bash 
echo 'eval "$(~/.local/bin/mise activate bash)"' >> ~/.bashrc
echo 'eval "$(~/.local/bin/mise activate bash --shims)"' >> ~/.bash_profile
source ~/.bashrc

# Fish
echo '~/.local/bin/mise activate fish | source' >> ~/.config/fish/config.fish
fish_add_path ~/.local/share/
```

3. Run `mise trust` to trust the project's `.mise.toml` file.
