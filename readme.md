# Exploring API's in GO and REPL patterns


## Types

for json endpoints I used: https://mholt.github.io/json-to-go/ to convert them into structs
## Extras

I wanted to add up / down arrows
I also didn't want to install a package that emulated the terminal to capture commands (probably should have)
but the terminal operates in raw mode requiring \n\r manually to reset the cursor position.

This allows the editing of the current input line.
It also required handling Ctrl + C seperatly and restoring the terminal state. 
There is two exit modes -> exit command first disables rawmode then exits normally
using ctrl+c is handled by a defered clean up statemtn in the internal cli package