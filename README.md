# pwnpatrol


This is a simple service that replicates the *Pwned Passwords* endpoint of the haveibeenpwned.com API.It can be used if you want that functionality without having to rely on a third party API.

## Quickstart

1. Download any of the SHA-1 passwords list from https://haveibeenpwned.com/Passwords. I strongly recommend using the torrent instead of the direct download link.
2. Decompress the file: `7z x pwned-passwords-ordered-by-hash.7z`
3. Initialize the DB file: `PWNPATROL_DUMPFILE=/tmp/pwned-passwords-ordered-by-hash.txt PWNPATROL_DUMPFILE=/tmp/pwnpatrol.db pwnpatrol initdb`
4. Start the server: `PWNPATROL_DBFILE=/tmp/pwnpatrol.db pwnpatrol serve`
