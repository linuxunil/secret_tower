Great—let’s lock the spec with **Option A (bracketed metadata)**. Clean to read, trivial to parse.

# Room Markdown v1.0 (Bracket Meta)

**One file per room.** The **room ID** is auto-generated from the H1 (sanitize: lowercase, non-alnum → `_`, collapse `_+`, trim `_`).

## Structure (fixed order)

```markdown
# <Room Name>

<One or more paragraphs of description.>

---

### Exits
- <Direction> → <Target Room Name> [locked, key=<Key Name>, one-way, hidden]

### Objects
- **<Object Name>** — <short description> [portable, hidden, key=<Key Name>]

### Notes
<optional free text>
```

### Notes on syntax

* Arrow: accept either `→` or `->`.
* **Metadata**: everything inside `[...]` is a comma-separated list of **flags** (e.g., `locked`, `hidden`, `one-way`) and **key=value** pairs (e.g., `key=Rusty Key`).
* **Sanitize** any human name to an ID the same way you do for the room H1 (e.g., `Great Hall` → `great_hall`, `Rusty Key` → `rusty_key`).

---

## Example (contributors write this)

```markdown
# Entrance Hall

You stand in a grand stone hall. A heavy wooden door leads outside.
To the north, a staircase climbs into darkness. To the east, a wide archway opens into a dining room.

---

### Exits
- North → Great Hall
- East → Dining Room
- Outside → World Exit [locked, key=Rusty Key]

### Objects
- **Rusty Key** — a corroded iron key [portable]
- **Wall Sconce** — fixed to the wall, its flame gutters [hidden]

### Notes
Acts as the hub where the player begins.
```

### What your compiler should emit

```json
{
  "id": "entrance_hall",
  "name": "Entrance Hall",
  "description": "You stand in a grand stone hall. A heavy wooden door leads outside.\nTo the north, a staircase climbs into darkness. To the east, a wide archway opens into a dining room.",
  "exits": {
    "north":   { "to": "great_hall" },
    "east":    { "to": "dining_room" },
    "outside": { "to": "world_exit", "locked": true, "key": "rusty_key" }
  },
  "objects": [
    { "id": "rusty_key", "name": "Rusty Key", "description": "a corroded iron key", "portable": true },
    { "id": "wall_sconce", "name": "Wall Sconce", "description": "fixed to the wall, its flame gutters", "hidden": true }
  ],
  "notes": "Acts as the hub where the player begins."
}
```

---

## Parsing rules (tight, line-based)

**Anchors**

* Name: first H1 — `^#\s+(.+)$`
* Separator: `^---\s*$`
* Section headers (exact titles): `### Exits`, `### Objects`, `### Notes`

**List item formats**

* **Exit line** (accepts `→` or `->`):

  ```
  - <dir> (spaces ok) →|-> <target> (up to '[') [<meta>]
  ```

  Regex:

  ```
  ^-\s*(?P<dir>[^-\[\n]+?)\s*(?:→|->)\s*(?P<target>[^\[\n]+?)(?:\s*\[(?P<meta>[^\]]+)\])?\s*$
  ```
* **Object line**:

  ```
  - **<name>** — <desc> [<meta>]
  ```

  Regex:

  ```
  ^-\s*\*\*(?P<name>.+?)\*\*\s*—\s*(?P<desc>[^\[\n]+?)(?:\s*\[(?P<meta>[^\]]+)\])?\s*$
  ```

**Metadata parsing**

* Split `meta` on commas. For each token:

  * If it contains `=`, split once: `k=v`. Sanitize `v` to id if it names an entity (`key`, `to`, etc.).
  * Else it’s a boolean flag ⇒ set `flag: true`.
* Normalize direction to a lowercase id: `outside`, `north_east`, etc. (sanitize like names; you can also map common aliases: `n`→`north`).

---

## Quick “contract” (what to validate at compile time)

* H1 must exist; description must be non-empty.
* Exit directions unique within a room.
* Object IDs unique within a room.
* Cross-file check: `exit.to` refers to an existing room ID.
* Unknown metadata keys are allowed (warn, don’t fail) for forwards-compat.

If you want, I can package this into a tiny CLI that walks `rooms/*.md` and writes `rooms.json`, plus a couple of sample files for unit tests.
