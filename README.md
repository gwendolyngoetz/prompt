# README

## Example

![image](https://user-images.githubusercontent.com/195162/169644884-2200e447-7510-44c1-8106-6faa3f35dfe8.png)

## Notes

### Tested on

* Ubuntu 20.10
  * Alacritty
* MacOS Monterey
  * Alacritty
  * iTerm2 
* Windows 10
  * Windows Terminal 

## Build and Install

### Dependencies

* go
* make

In addition a [Nerd Font](https://www.nerdfonts.com) needs to be installed on the system for icons to display. The Cousine Nerd Font is used in the sample image.

### Install
```bash
make
make install
```

## Configure Shell

### Bash
Add to `~/.bashrc`

```bash
function set_prompt {
    PS1=$($HOME/.local/bin/prompt)
}

if [[ -f "$HOME/.local/bin/prompt" ]]; then
    PROMPT_COMMAND=set_prompt
fi
```

### Powershell
Add to `$env:UserProfile\Documents\PowerShell\Microsoft.PowerShell_profile.ps1`

```powershell
function prompt {
    if (Test-Path -Path "$env:LocalAppData\prompt\promptwin.exe" -PathType Leaf) {
        # Only displaying the first line by defailt so forcing it
        $a = invoke-expression "$env:LocalAppData\prompt\promptwin.exe"
        $a[0] + "`r`n" + $a[1]
    }
}


# Renders poorly upon first opening the shell, but fine afterwards
cls

# Required for unicode icons to show correctly
[Console]::OutputEncoding = [Text.UTF8Encoding]::UTF8

```

