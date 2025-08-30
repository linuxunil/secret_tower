# Library

Dusty shelves sag under the weight of ancient books. The faint smell of parchment and mildew hangs in the air.
A tall ladder leans precariously against the highest shelf, and cobwebs gather thick in the corners.
Every so often, a faint draft stirs the pages of an open tome, though no window is visible.
It feels as if some secret passage lies hidden among the rows.

---

### Exits
- South → Study [locked, key=study_key]
- Secret → Hidden Tunnel

### Objects
- **ladder** — A tall, movable ladder. [portable]
- **old_tome** — An open, ancient tome with fluttering pages.
- **secret_mechanism** — A hidden mechanism among the shelves.

### Notes
Flags:
- indoor

On Enter:
- seen_library=true

On Leave:
- log("leaving library")