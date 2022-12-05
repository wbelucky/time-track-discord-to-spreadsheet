# time-track-discord-to-spreadsheet

This is a discord bot that records the time spent in the discord in a spreadsheet.

The time when you enter the voice chat is written in the range A2:A, and the time when you exit is written in the range C2:C.
The range can be aggregated in a pivot table by day, etc., and visualized in a graph.

## Requirements

1. You need to create gcp project, add spreadsheet API to the project and create service account.

- put `credentials/secret.json`

2. You need to create discord bot.
3. run `cp .env.sample .env` and fill .env

