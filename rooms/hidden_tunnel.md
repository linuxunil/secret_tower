# Hidden Tunnel

A narrow stone passage winds downward. The air grows colder with each step.
Faint sounds of dripping water echo from below.

---

### Exits
- Up → Wine Cellar
- Down → Underground Lake
- Secret → Library

### Objects
- **loose_stones** — Stones that shift underfoot.
- **dripping_water** — Drops echoing in the dark.

### Notes
Flags:
- indoor

On Enter:
- seen_hidden_tunnel=true

On Leave:
- log("leaving hidden_tunnel")