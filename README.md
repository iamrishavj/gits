# gits
my CLI tool to manage and switch between different git accounts with SSH support. In future maybe will add using different sessions with multiple git accounts at the same time.


## Steps + Questions to be asked

 - figure out how GIT stores and uses username and email + also SSH storage too.
 - Can it be changed while running a command?
 - can I make a superset of git and pass userprofile like: gits PROFILE_NAME ... whatever git command

## CLI tooling options

### Check git accounts

gits ls\
Lists all the active accounts

### Create

gits create PROFILE_NAME USERNAME EMAIL\
Note: maybe generate a SSH and ask user to add the SSH key to their account

### Delete

gits delete PROFILE_NAME\
Note: ask for confirmation\
Also clear out SSH keys

### Update

gits update PROFILE_NAME NEW_USERNAME NEW_EMAIL\
Note: Add same details for either USERNAME or EMAIL if not updated

### SSH

gits generate PROFILE_NAME ssh\
Note: if SSH exists ask still want to continue, if yes generate and copy to clipboard, no then copy old one

### etc

All git commands should work as is..