[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-black.svg)](https://snapcraft.io/lfmg)   

# lfmg
lfmg is LICENSE file generator based on github api.   

## install
```
$ snap install lfmg
```

## test
It works fine in `Ubuntu 18.04`.

## command
```
   list, l           show "LICENSE" list
   generate, gen, g  generate "LICENSE".
   info, i           get license info
   body, b           show "LICENSE" content
   cache, c          cache control
   help, h           Shows a list of commands or help for one command
```

### e.g.
create mit "LICENSE" on working directory
```
$ lfmg g -y 2021 -o vompressor mit
```

$ create Apache-2.0 "LICENSE" on "~/project/LICENSE"
```
$ lfmg g -p ~/project/ apache-2.0
```

get license list
```
$ lfmg l

key : id : name
agpl-3.0 : AGPL-3.0 : GNU Affero General Public License v3.0
apache-2.0 : Apache-2.0 : Apache License 2.0
bsd-2-clause : BSD-2-Clause : BSD 2-Clause "Simplified" License
bsd-3-clause : BSD-3-Clause : BSD 3-Clause "New" or "Revised" License
bsl-1.0 : BSL-1.0 : Boost Software License 1.0
cc0-1.0 : CC0-1.0 : Creative Commons Zero v1.0 Universal
epl-2.0 : EPL-2.0 : Eclipse Public License 2.0
gpl-2.0 : GPL-2.0 : GNU General Public License v2.0
gpl-3.0 : GPL-3.0 : GNU General Public License v3.0
lgpl-2.1 : LGPL-2.1 : GNU Lesser General Public License v2.1
mit : MIT : MIT License
mpl-2.0 : MPL-2.0 : Mozilla Public License 2.0
unlicense : Unlicense : The Unlicense
```

gpl-3.0 info
```
$ lfmg i gpl-3.0

key: gpl-3.0
id:  GPL-3.0
url: http://choosealicense.com/licenses/gpl-3.0/

description:
Permissions of this strong copyleft license are conditioned on making available complete source code of licensed works and modifications, which include larger works using a licensed work, under the same license. Copyright and license notices must be preserved. Contributors provide an express grant of patent rights.

implementation:
Create a text file (typically named COPYING, as per GNU conventions) in the root of your source code and copy the text of the license into the file.

permissions: commercial-use modifications distribution patent-use private-use
conditions: include-copyright document-changes disclose-source same-license
limitations: liability warranty
```

## library
https://github.com/vompressor/license_generator

## TODO::
file permission   
err message   
bash hint   
move server github api to my api server   