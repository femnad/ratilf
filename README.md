# ratilf

Helpers for focusing windows with [Ratpoison](https://www.nongnu.org/ratpoison/).

## Usage

### Rofi Modes

These list windows when no non-flag arguments are provided and focus a window then exit with a status of `1` if any non-flag arguments are provided. Intended to be used alongside [rofi](https://github.com/DaveDavenport/rofi), as rofi modes:

### ratfocus

List all windows and focus the selected one

```
rofi -show ratfocus -modi ratfocus:ratfocus
```

### ratsame

Show and focus windows of the same class

```
rofi -show ratsame -modi ratsame:ratsame
```

### ratterm

Show and focus windows of given terminal class

```
rofi -show ratterm -modi ratterm:'ratterm -class <terminal-class>
```

## Standalone

Intended to be executed directly by Ratpoison:

### ratraise

Run or raise window of given executable and class (optional)

```
ratraise -exec <executable> -class <class-name>
```

## Example

Example configuration with some dedicated submaps

```
bind g exec rofi -show ratfocus -modi ratfocus:ratfocus

newkmap cexe
definekey cexe c exec rofi -show ratterm -modi ratterm:'ratterm -class Tilix'
definekey cexe g exec rofi -show ratsame -modi ratsame:ratsame
definekey root h readkey cexe

newkmap run-or-raise-map
definekey run-or-raise-map c exec ratraise -exec tilix
definekey root space readkey run-or-raise-map
```
