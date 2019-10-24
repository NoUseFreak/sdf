# SDF

I wrote `sdf` because I wanted my project structure to be organized. But setting
this up for every repository, and navigating to the correct directory felt too
much work. 

More lazy development tools are being added as they are needed.

## Installation

```bash
brew install nousefreak/brew/sdf
```

Because `sdf` needsto be able to do some bash magic, it will need to wrapper.
Run the following once to complete the installation.

```bash
sdf setup
```

Once this is done, reload your session and you are good to go.

## Commands

__cd__

> Change to a directory using fuzzy search.

```bash
sdf cd g/nuf/sdf
# will cd to ~/Projects/src/github.com/NoUseFreak/sdf
```

__clone__

> Clone a project into the project structure.

```bash
sdf clone g/NoUseFreak/sdf
# Will clone and cd to ~/Projects/src/github.com/NoUseFreak/sdf
```

__env__

> Manage environment variables your don't want to save in your projects.

- __cat__ Print a set
- __edit__ Edit an existing set
- __list__ List all available sets
- __new__ Create a new set
- __rm__ Remove a set
- __use__ Exports a set to the current session

## Configuration

After setup, you will have a `~/.sdf/main.yml` file, this contains some settings
that can tailor the `sdf` commands to your way of working.

- __profile__ Profile name if you want more profiles, default is main
- __projectdir__ Root of all projects. 
