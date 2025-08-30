# Dining Room

A long mahogany table dominates the room. Silverware lies in neat arrangement, though untouched for decades.
The kitchen is through a door to the east.

---

### Exits
- West → Entrance Hall
- East → Kitchen

### Objects
- **silverware_set** — A neat but dusty set of silverware.
- **dining_table** — A long mahogany table.
- **chairs** — Old dining chairs.

### Notes
Flags:
- indoor

On Enter:
- seen_dining_room=true

On Leave:
- log("leaving dining_room")