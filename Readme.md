# deanlang

### About

deanlang is an interpereted esoteric scripting language

---

### Usage

the command to use deanlang is `deanl` which requires an input file which is the deanlang script.

The file extension for a deanlang script can be anything but it is recommended to use `.dl` to tell that it is a deanlang script

---

### Functionality

In deanlang, there are a couple of things that you should keep in mind.

#### Main Byte

This is a single byte of memory that can be manipulated

#### The Reference

This is a string of characters that can determine the variable being referenced, or the goto point being referenced.

#### References

These are spots in memory that store: singular bytes, and goto points.

---

### Syntax

There are only a few key words in the deanlang syntax.

#### Comments

The comments are `#`, comments must be sandwiched with another `#` to work, for example `# this is my comment #`.

#### Increment / Decrement

Incrementing increases the main byte by one with a `+`.

Decrementing decreases the main byte by one with a `-`.

#### I/O

`?` would read from stdin and write it to the main byte.

`!` would write the main byte's value into stdout.

#### Variables



#### Unloading

Unloading zeros out the main byte with `,`.

#### If Statements


#### Gotos

Gotos allow for going to different points in a script.

`*` starts a goto point with the label of what the reference is: `my_point*` creates a goto point with the label of `my_point`.

`&` goes to a goto point with the label of the reference: `my_point&` goes to the point labeled `my_point`.

#### Jumps



#### Do Nothing / No Opp

The do nothing operator is `~` which cannot have anything assigned to it.

---

### More

You can find examples in [examples](examples/)