Listens for a HTTP POST, and triggers a repository dispatch event via the Github API

Github PAT:
If  $HOME/.ghtoken exists, access token is read from that file, otherwise prompt on startup for access token from stdin. 

Genesis: This is a a quick and dirty way to glue a generic ghost webhook post to generating a Github repository dispatch event.
