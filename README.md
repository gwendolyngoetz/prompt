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

## Setup

### Dependencies

* go
* make

In addition a [Nerd Font](https://www.nerdfonts.com) needs to be installed on the system for icons to display. The Cousine Nerd Font is used in the sample image.

### Install
```bash
make
make installl
```


### Add to .bashrc
```bash
function set_prompt {
    PS1=$($HOME/.local/bin/prompt)
}

if [[ -f "$HOME/.local/bin/prompt" ]]; then
    PROMPT_COMMAND=set_prompt
fi
```
