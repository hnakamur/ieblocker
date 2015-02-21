ieversionlocker
===============

A command line tool to lock the version of the Internet Explorer.
Please run this command at the Administrator Command Prompt.

[README in Japanese](README_ja.md)

## Usage

1. Download https://github.com/hnakamur/ieversionlocker/raw/master/dist/windows_32bit/ieversionlocker.exe with the Internet Explorer and save it in the "Downloads" folder.
2. Open the Administrator Command Prompt.
    * Windows Vista or Windows 7:
        1. Press WINDOWS key.
        2. Click the "All programs" menu.
        3. Click the "Accessories" folder.
        4. Right click the "Command Prompt" menu.
        5. Click the "Run as admnistrator" menu.
    * Windows 8 or 8.1:
        1. Press WINDOWS+X key in the desktop window.
        2. Press A key to select the "Command Prompt (Administrator)" menu.
3. Run the following commands

```
cd %USERPROFILE%\Downloads
ieversionlocker.exe -l
```

If you want to unlock the version of the Internet Explorer, run the following commands.

```
cd %USERPROFILE%\Downloads
ieversionlocker.exe -u
```

## License

MIT
