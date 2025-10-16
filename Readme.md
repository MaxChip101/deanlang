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

This is a string of characters that can determine the variable you're referencing, or what goto you're referencing

#### References

These are spots in memory that store: byte data, and goto points

---

### Syntax

There are only a few key words in the deanlang syntax.

#### Comments

The comments are `#`, you have to sandwich comments with the `#` to work, for example `# this is my comment #`.

#### Increment / Decrement

Incrementing is when you increase the main byte by one by doing `+`. Decrementing is when you decrease the main byte by one by doing `-`.

#### I/O

Input is when you get a single byte from stdin by doing `?` this saves it to the loaded byte. Output is when you write the main byte to stdout by doing `!`.

#### Variables

Variables are points of memory that store a single byte, variables are any character that is not a whitespace, key word, or comment, for example: `my_variable`. To write and load to and from variables you would need to use `:` and `.`, `:` saves the loaded byte into the variable, while `.` loads the variable's byte into the loaded byte. Variables need a way to reset what variable you're referencing, in order to do that you need to use `;`, this would forget what variable you're referencing. This means that you can have variables like this: `++ my -- _ + variable # this is still my_variable even though it is split up #`. If you don't want to reset your referenced variable, then you can use `/` to subtract the last character from the referenced variable.

#### Unloading

When the value in the loaded byte is unkown and can cause problems, you would use `,` to unload the loaded byte, this would just set it to zero.

#### Conditions

When needing to evaluate 2 values and do something based on if the condition is true, you would use `|`, this has to be sandwiched between a segment of code like: `| ++ |`. How the conditions work is it takes the referenced variable value and compares it to the loaded byte, if they are equal then the code in the segment will be executed, if it is false, then the code in the segment will be skipped.

#### Jumps

When needing to travel to parts of your code it is useful to have jumps, to use jumps you need to have `*` and `&` or `<` or `>`. The `*` saves a position in your code with the referenced jump variable, for example: `my_jump* # this makes a jump point with a tag of my_jump #`. To jump to these points you need to use `&`, this will jump to the jump point that has the same referenced jump variable. These can be changed throughout the program if needed. The `<` will jump backwards by a set amount of steps, this is set by the loaded byte, The same is with `>` but forward, it will use the loaded byte's value to tell how far forward to jump for example: `+>! # The write was skipped #.

#### Do Nothing / No Opp

If for some reason you need a couple of characters that do nothing that fix your code then you can use `~` which does nothing and cannot be used as a variable

---

### More

You can find examples in [examples](examples/)