# deanlang

### About

deanlang is an interpereted esoteric scripting language

---

### Usage

The command to use deanlang is: `deanl`

Some flags are: `--debug` and `--help`

deanlang scripts have a file extension of: `.dl`

deanlang scripts are ran by doing: `deanl script.dl`

deanlang can also debug what the script is doing under the hood: `deanl script.dl --debug`

deanlang has a help command to help with understanding the interpereter: `deanl --help`

The file extension for a deanlang script can be anything but it is recommended to use: `.dl` to tell that it is a deanlang script

---

### Functionality

In deanlang, there are a couple of things that you should keep in mind.

#### Main Byte

This is a single byte of memory that can be manipulated

#### The Reference

This is a string of characters that can determine the variable being referenced, or the goto point being referenced.

#### References

These are spots in memory that store: a singular byte, and goto points.

---

### Syntax

There are only a few key words in the deanlang syntax.

#### Comments

The comments are `#`, comments must be sandwiched with another `#` to work, for example: `# this is my comment #`.

#### Increment / Decrement

Incrementing increases the main byte by one with a `+`.

Decrementing decreases the main byte by one with a `-`.

#### I/O

`?` would read from stdin and write it to the main byte.

`!` would write the main byte's value into stdout.

#### Variables

Variables are the byte value of the reference.

`:` saves the main byte's value to the referenced variable.

`.` loads the referenced variable's byte value into the main byte.

#### Reference Operators

`;` clears the reference

`/` deletes the last character of the reference

#### Unloading

Unloading zeros out the main byte with `,`.

#### If Statements

If statements check if the referenced variable byte value is the same as the main byte's value. If statements cannot be stacked

`{` starts an if statement, this is where the condition gets checked.

`}` ends an if statement.

If the condition is true, the code in the if statement will run, if not it will skip to the `}`
#### Gotos

Gotos allow for going to different points in a script.

`*` starts a goto point with the label of what the reference is: `my_point*` creates a goto point with the label of: `my_point`.

`&` goes to a goto point with the label of the reference: `my_point&` goes to the point labeled: `my_point`.

#### Jumps

`<` jumps backwards by the amount in the main byte.

`>` jumps forwards by the amount in the main byte.

#### Do Nothing / No Opp

The do nothing operator is `~` which cannot have anything assigned to it.

---

### More

You can find examples in [examples](examples/)