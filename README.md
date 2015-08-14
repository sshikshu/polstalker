#Twitter data Reader
Twitter data analyzer.


###Usage
  --count (default 100): number of tweets to analyze

  --hash (default false): targets's hashtags arranged by date

  --hash-freq (default false): targets's hashtags arranged by frequency

  --loc (default false): target's locations arranged by date

  --loc-freq (default false): target's locations arranged by frequency

  --media (default false): targets's media arranged by date

  --media-freq (default false): targets's media arranged by frequency

  --mention (default false): targets's user mentions arranged by date

  --mention-freq (default false): targets's user mentions arranged by frequency

  --name (default prvn_30): target's twitter handle

  --nocolor (default false): plain colorless output

  --summary (default true): summary of the target's account

  --url (default false): targets's urls arranged by date

  --url-freq (default false): targets's urls arranged by frequency


###Example
  ```bash
  polstalker-twitter -name prvn_30 --hash-freq --mention-freq
  ```
