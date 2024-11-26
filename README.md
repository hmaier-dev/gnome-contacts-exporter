## gnome-contacts-exporter

This program is my cli to the `contacts.db` used by [gnome-contacts](https://gitlab.gnome.org/GNOME/gnome-contacts),
which usually lies in `~/.local/share/evolution/addressbook/system/contacts.db`.

My intention was to automate the exporting into a `vCard`-File (`contacts.vcf`), which can be done by hand in gnome-contacts.




### Usage

For exporting use the arguments:

- `--source contacts.db --destination contacts.vcf`

and for importing use:

- `--source contacts.vcf --destination contacts.db`


### TODO

- [ ] converting from `vcf` to `sqlite3` (Import)

- [ ] converting from `sqlite3` to `vcf` (Export)

- [ ] implement daemon-function, which listens to `sqlite3`/`vcf` and imports/exports automatically

- [ ] adding `--dry-run` for testing

- [ ] add `Earthfile` for testing-environment

- [ ] Write Tests for different cases
    - [ ] 2 emails + 1 phone-numbers + comments + adress
        - [ ] Export
        - [ ] Import
    - [ ] 1 email + 2 phone-numbers + comments + adress
        - [ ] Export
        - [ ] Import
- [ ] Find out how to use CI to build binarys
- [ ] build `systemd-unit` to run the exporter


