---
title: The GUI's pages and their surfaces
number: 2
column: 5ea2db0272fc
state: ready
severity: major
priority: soon
---
Decide what pages DinahVisor has and what each one shows. This is the GUI half of the surfaces question that was being worked on the Andoneer board as dinah-338; the other half, which covers how a command renders and where the GUI boundary falls inside dinah itself, stayed there.

The pages are the product, so this card is where the product gets decided rather than a layout exercise. Every state the GUI can be in should be addressable, which follows from the URL-is-king principle in the operator's architectural principles document, and elevation should mean interactivity per the Tela Design Language. Neither of those settles what a page is for, which is what a spec has to answer.

Related work on the dinah side: dinah-152 carries the importable, mountable HTTP handler this GUI runs, and its route roster is the published interface these pages read. A page needing a route the roster does not carry is a card on the dinah board, not a workaround here.
