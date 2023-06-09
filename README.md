# genpass

## Simple password generator

### Usage
```
  -a    Include alphabet characters.
  -c string
        Custom character set.
  -l int
        Length of each password. (default 20)
  -n    Include numbers.
  -r int
        Number of results (will be ignored if "-u" is specified). (default 1)
  -s    Include special characters.
  -u string
        Users (comma-separated).
```

### Samples

```bash
./genpass -a
# YragTrcoRiVeSpgPUPDZ

./genpass -s
# [*([&.^+"/]_./]+`,(_

./genpass -c "1AZ%" -l 28
# ZAZAA11A1A%Z%AA%AAAZ%1%1%%AA

./genpass -u admin,user1,user2,user3,demo,test -a -n
# admin   4zUUe8CHKHAMeO51eC84
# user1   MgqHIknbjEwDR6EVug7T
# user2   jrDEycHSRKCBS8ggYHHg
# user3   nCWVF0oBqYyh5XR2dDDg
# demo    d9p3Rpk6jXDwfWb6NgkZ
# test    9OqCP1qI6mPcGbK0DGF8
```

### Download

[Linux](https://github.com/molodsom/genpass/releases/latest/download/genpass_linux_amd64), [Windows](https://github.com/molodsom/genpass/releases/latest/download/genpass_windows_amd64.exe), Mac OS ([Intel](https://github.com/molodsom/genpass/releases/latest/download/genpass_darwin_amd64), [M1+](https://github.com/molodsom/genpass/releases/latest/download/genpass_darwin_arm64))
