# MVP - wilding-trello tool

I wanted to archive a bunch of cards I had made in Trello, so I built this very simple tool. 


## Current functionality  

Give an boardID (https://trello.com/b/**######**/...), outputs all  Card names and descriptions from a Trello board with the Cards Name as an `<h1>`


- to run you need to get an appkey and a token
- [trello app key ]( https://trello.com/app-key)
- there you can generate a token as well
- use those to add the correct environmental variables
- now copy the id from the board you want to get the cards from

`export TRELLO_APPKEY=... && export TRELLO_TOKEN=... && mw-trello-client {boardID} > fileofyourchoice.md` 
