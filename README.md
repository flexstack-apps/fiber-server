# fiber-server

A Go Fiber web server with graceful exit, structured logging, environment variable configuration, and local development tooling.

## Local development

### Quick start

See the [Prerequisites](#prerequisites) section for installing [mise](https://mise.jdx.dev/about.html).

```sh
# Install project dependencies
mise install

# Setup the project
mise run setup

# Start the development server
mise run
```

### Prerequisites

We use [mise](https://mise.jdx.dev/about.html) to run tasks and manage our tool version and environment variables.
This ensures that all developers and deployments are using the same versions of tools and dependencies.

1. [Install mise](https://mise.jdx.dev/getting-started.html)

```sh
curl https://mise.run | sh
```

2. Add mise shims to your shell profile and source it. This will activate mise in your shell
   and add shims to your path so they can be used in non-interactive places like IDEs.

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

3. Run `mise trust` to trust the project's `.mise.toml` file. This will allow mise to manage the project's tool versions and environment variables.

### Scripts

We use `mise` as our [task runner](https://mise.jdx.dev/tasks/running-tasks.html).

1. To see a list of development scripts, run `mise tasks`.
1. To run a script, run `mise run <script-name>`.

### VS Code and Cursor

If you're using VS Code or Cursor as your editor, we recommend the following extensions:

- [**TOML Language Support**](https://marketplace.visualstudio.com/items?itemName=be5invis.toml) Add TOML language support for the `mise` config.
- [**Dprint Code Formatter**](https://marketplace.visualstudio.com/items?itemName=dprint.dprint) Use `dprint` as the default formatter for the project.
- [**Go**](https://marketplace.visualstudio.com/items?itemName=golang.go) Golang support for VS Code
