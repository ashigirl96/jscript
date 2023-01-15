# Jscript

## Install

```shell
go install github.com/ashigirl96/jscript@latest
```

## Command

1. Show scripts from package.json 'scripts'
2. Run one script that selected from package.json scripts

### Show 'scripts'

```shell
❯ jscript
dev:        	next dev
build:      	next build
start:      	next start
lint:       	next lint
```

### Run Script

```shell
> jscript run
build  -- next build
dev    -- next dev
lint   -- next lint
start  -- next start
```

## Shell completions

mimic https://deno.land/manual@v1.29.3/getting_started/setup_your_environment#shell-completions

Current shells that are supported:

- bash
- fish
- powershell
- zsh

### bash example

Output the completions and add them to the environment:

```shell
> jscript completion bash > /usr/local/etc/bash_completion.d/jscript.bash
> source /usr/local/etc/bash_completion.d/jscript.bash
```


### zsh example

You should have a directory where the completions can be saved:

```shell
> mkdir ~/.zsh
> jscript completion zsh > ~/.zsh/_jscript
```

And ensure the completions get loaded in your ~/.zshrc:

```shell
fpath=(~/.zsh $fpath)
autoload -Uz compinit
compinit -u
```

If after reloading your shell and completions are still not loading, you may need to remove `~/.zcompdump/` to remove previously generated completions and then `compinit` to generate them again.

## Support OS

- (Untested) Linux
- Mac OS X