# TODO

- Von VCF zu sqlite3

- Von sqlite3 zu vcf

- Tests für verschiedene Cases schreiben
    - 2 Emails + 1 Nummer + Kommentar + Adresse
    - 1 Email + 2 Nummern + Kommentar + Adresse

- systemd-unit generieren um script/binary zu deployen
    - Installations-Script?
    - Welche Dinge können sich unterscheiden pro system?
        - Wo liegt sqlite3-Datei?
            - bei gnome-contacts in `~/.local/share/evolution/addressbook/system/contacts.db`
        - Welcher Ordner wird gesynct?
            - in meinem Setup `~/docs/contacts`

