Listens for an HTTP POST, and triggers a repository dispatch event via the Github API

access token is read out of $HOME/.ghtoken

This is a a quick and dirty way to glue a generic ghost webhook post to generating a Github repository dispatch event.
