# My World Map
<!--
This is the top-level heading. It can be the title of your world or adventure.
Not required by the parser, just for you.
-->

## [room-id] Room Title
<!--
Every room starts with a heading: ## [id] Title
- The [id] is a unique slug (lowercase, no spaces).
- The Title is the display name shown to the player.
-->

This is the room description.
Write one or more sentences of text here.
Keep it natural — this is what the player will read.

Exits:
<!--
List the exits as "- direction -> room-id"
- direction must be one of: north, south, east, west, up, down (short forms n, s, e, w, u, d also ok)
- room-id is the [id] of the destination room
-->
- north -> other-room-id
- east -> another-room-id

Items:
<!--
List items found in the room.
- Format: - item-id* : description
- Use * after the id if the item can be picked up (portable).
- Omit * for fixed or decorative items.
-->
- torch* : A wooden torch that still holds a little pitch.
- statue : A cracked stone statue.

Flags:
<!--
Optional tags that describe the room.
Use them however you want (examples: start, dark, indoor).
-->
- start
- indoor

---

## [another-room] Another Room
<!--
Repeat the same pattern for every room.
You can put as many rooms as you like in this file.
-->
A different description goes here.

Exits:
- west -> room-id

Items:
- note : A folded piece of paper with faint writing.

Flags:
- dark
