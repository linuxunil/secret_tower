# Great Hall

Long tables stretch across this cavernous hall. Banners of forgotten heraldry hang from the rafters.
A grand spiral staircase leads upward.

---

### Exits
- South → Entrance Hall
- West → Study [locked, key=study_key]
- East → Dining Room
- Up → Spiral Stair

### Objects
- **dining_tables** — Long wooden tables.
- **faded_banners** — Ancient heraldic banners.

### Notes
Flags:
- indoor

On Enter:
- seen_great_hall=true

On Leave:
- log("leaving great_hall")