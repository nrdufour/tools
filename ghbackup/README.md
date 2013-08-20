# Github Backup

## what does it do?

### well, mirroring!

A small Go program to quickly backup all the repositories based on a github HTTP request.

The usual usage is: 

`curl https://api.github.com/users/<username>/repos | ./ghbackup`

where <username> is your Github username.

You can also use it to retrieve for a given organization:

`curl https://api.github.com/orgs/<orgname>/repos | ./ghbackup`

**ghbackup** will for each repository:

+ clone it with --mirror option
+ add 2 local properties (useful for cgit):
  + gitweb.description: will contain the description set in github, or the project name if the description is empty
  + gitweb.owner: will be set with the actual github owner

### Second pass

If you run it again with the same repositories already mirrored, **ghbackup** will try to update the mirror from Github instead.

## why?

The main purpose was to retrieve and backup all the repositories I have and setup a cgit server both at home and on the cloud, in case Github fails for some reason (like a DDoS for exemple ;-) ).

Or to make a short story: *don't put your eggs in the same basket*.

## it looks hacky/awful!

Yeah yeah, I know. It's just a simple tool I wrote quickly and does the job for now. Nothing forbids you to make it better.

---
Thanks,

Nicolas Dufour