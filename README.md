# Clean Chat Message

A lightweight, efficient Go utility designed to clean up and format messy chat messages. It normalizes irregular spacing, collapses excessive consecutive punctuation marks, and ensures proper sentence capitalization.

## Features

* **Whitespace Normalization:** Automatically strips away redundant multi-spaces between words.
* **Punctuation Collapsing:** Reduces consecutive identical punctuation marks (e.g., `????` or `!!!`) down to a single clean character.
* **Smart Sentence Capitalization:** Capitalizes the first letter of sentences following punctuation marks (`!`, `?`, `.`, `,`).
* **Rune-Safe Processing:** Processes strings as `rune` slices to safely handle standard and multi-byte characters.

---

## How It Works

The utility processes raw text input through a three-stage cleaning pipeline:

1. **Tokenization & Spacing:** Uses `strings.Fields` to split the message into individual words, instantly removing any extra spaces or tabs.
2. **Punctuation Deduplication:** Iterates through each word. If it encounters a punctuation mark (`!`, `?`, `.`, `,`), it skips over any identical, consecutive characters that immediately follow it.
3. **Reassembly & Capitalization:** Joins the cleaned words back together with a single space and scans the full string to capitalize the beginning of new sentences.

### Input / Output Example

* **Input:** `"heyyy!!!     how are  you????"`
* **Output:** `"Heyyy! How are you?"`

---

## Project Structure

```text
cleanChatMessage/
├── main.go                     # Application entry point and testing setup
├── cleanChatMessage.go         # Core chat-cleaning orchestration pipeline
├── collapseRepeatedPunctuation.go # Algorithmic logic for punctuation stripping
└── README.md                   # Project documentation