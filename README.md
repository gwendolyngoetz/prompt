# README

## Example

![image](https://user-images.githubusercontent.com/195162/169644884-2200e447-7510-44c1-8106-6faa3f35dfe8.png)

## Setup

### Dependencies

* go
* make

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
