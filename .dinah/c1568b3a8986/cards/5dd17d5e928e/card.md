---
title: "DinahVisor: a desktop GUI for Dinah"
number: 1
column: 5ea2db0272fc
state: ready
severity: major
priority: next
---
DinahVisor is a desktop GUI for Dinah, built as a Wails shell around a set of pages and a command log, styled by the Tela Design Language. It imports dinah as a Go library and runs the HTTP head's handler in process, so the pages talk to the same code the CLI does rather than to a reimplementation of it. This card is the founding brief: it exists to be specced into the first buildable pieces rather than built directly.

The shape was ratified after weighing three alternatives. A `dinah ui` verb inside the dinah repository was rejected, because a GUI is a product with its own release cadence, its own dependency tree and its own review needs, and burying it in the CLI's repository makes both harder to change. Serving the GUI from a webview against a running HTTP head was rejected for the same reason the whole-repo shape won: the interesting work is the pages, not the transport. What remains is this repository owning the entire GUI and depending on dinah as a library.

The cross-repo contract is dinah's HTTP handler, importable and mountable, with its route roster a published interface between the two repositories. That work lives on the dinah side and is not this card's to do. Until it lands, the go.mod pin here names whatever commit this repository builds against, and a change needed on the dinah side is filed on the Andoneer board rather than here.

What a spec would have to settle first: which pages exist and what each one is for, what the command log shows and how it relates to what the pages did, and whether the first buildable thing is a page or the shell that hosts one. The founding brief on the Andoneer board (dinah-283) carries the discussion this card came from.
