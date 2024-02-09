// defaultcontent.go
package main

const HelpText = `Repose Commands:
Usage: repose [OPTIONS] <COMMAND>

Commands:
	init    - Initialize a new Repose project
	new     - Create new content. Usage: repose new [CONTENTTYPE] [FILENAME]
	build   - Build the site.
	preview - Setup a local server to preview the site
	help    - Show this help message 
	
Options:
	-r, --root <ROOT> Directory to use as root of project (default: ./)
`

const DefaultConfig = `sitename: Repose site
author: Creator
editor: nano
contentDirectory: content
outputDirectory: web
url: mysite.com
previewUrl: http://localhost:8080`

const NewMD = `---
title: {title}
description: {contentType} about {title}
tags: []
image: 
index: true
author: {author}
publish_date: 
template: {contentType}.tmpl
---
	
# {title}

`

const MarkdownTest = `
---
title: Markdown Test Page
description: A test page to check markdown rendering
tags: []
image: 
index: true
author: Creator
publish_date: 
template: default.tmpl
---

# h1 Heading
## h2 Heading
### h3 Heading
#### h4 Heading
##### h5 Heading
###### h6 Heading
___

## Typographic replacements
Enable typographer option to see result.
(c) (C) (r) (R) (tm) (TM) (p) (P) +-
test.. test... test..... test?..... test!....
!!!!!! ???? ,,  -- ---
"Smartypants, double quotes" and 'single quotes'

## Emphasis
**This is bold text**
__This is bold text__
*This is italic text*
_This is italic text_
~~Strikethrough~~

## Blockquotes
> Blockquotes can also be nested...
>> ...by using additional greater-than signs right next to each other...
> > > ...or with spaces between arrows.

## Lists
Unordered
+ Create a list by starting a line with '+', '-', or '*'
+ Sub-lists are made by indenting 2 spaces:
  - Marker character change forces new list start:
    * Ac tristique libero volutpat at
    + Facilisis in pretium nisl aliquet
    - Nulla volutpat aliquam velit
+ Very easy!

Ordered
1. Lorem ipsum dolor sit amet
2. Consectetur adipiscing elit
3. Integer molestie lorem at massa


1. You can use sequential numbers...
1. ...or keep all the numbers as '1.'

## Tables

| Option | Description |
| ------ | ----------- |
| data   | path to data files to supply the data that will be passed into templates. |
| engine | engine to be used for processing templates. Handlebars is the default. |
| ext    | extension to be used for dest files. |

Right aligned columns

| Option | Description |
| ------:| -----------:|
| data   | path to data files to supply the data that will be passed into templates. |
| engine | engine to be used for processing templates. Handlebars is the default. |
| ext    | extension to be used for dest files. |


## Links

[link text](http://dev.nodeca.com)
[link with title](http://nodeca.github.io/pica/demo/ "title text!")
Autoconverted link https://github.com/nodeca/pica (enable linkify to see)

This is a footnote.[^1]

[^1]: the footnote text.

## Images

![Minion](https://octodex.github.com/images/minion.png)
![Stormtroopocat](https://octodex.github.com/images/stormtroopocat.jpg "The Stormtroopocat")

Like links, Images also have a footnote style syntax

![Alt text][id]

With a reference later in the document defining the URL location:

[id]: https://octodex.github.com/images/dojocat.jpg  "The Dojocat"

### Definition lists

Term 1

:   Definition 1
with lazy continuation.

Term 2 with *inline markup*

:   Definition 2

        { some code, part of Definition 2 }

    Third paragraph of definition 2.

_Compact style:_

Term 1
  ~ Definition 1

Term 2
  ~ Definition 2a
  ~ Definition 2b
`

const logo50 = `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" id="body_1" width="50" height="50">

<g transform="matrix(0.5 0 0 0.5 0 0)">
	<g transform="matrix(1 0 0 1 50 50)">
	</g>
	<g transform="matrix(NaN NaN NaN NaN 0 0)">
	</g>
	<g transform="matrix(0.33 0 0 0.33 50 50)">
		<g>
			<g transform="matrix(0.13 0 -0 -0.13 -67.64 -17.51)">
                <path transform="matrix(1 0 0 1 -744.86 -1378.36)"  d="M1180 2320C 1172 2315 1120 2302 1065 2292C 610 2203 266 1802 243 1334C 236 1194 251 1086 296 951C 348 793 426 666 543 548C 577 513 611 472 617 457C 632 424 674 420 690 449C 695 460 700 471 700 474C 700 487 666 510 646 510C 601 510 450 689 385 818C 143 1303 308 1873 769 2146C 904 2225 1150 2298 1193 2271C 1222 2252 1252 2267 1248 2298C 1246 2319 1239 2326 1220 2327C 1206 2328 1188 2325 1180 2320zM679 486C 685 479 688 465 684 456C 676 435 634 435 626 456C 619 475 636 500 655 500C 662 500 673 494 679 486z" stroke="none" fill="currentcolor" fill-rule="nonzero" />
			</g>
			<g transform="matrix(0.13 0 -0 -0.13 -61.96 -68.42)">
                <path transform="matrix(1 0 0 1 -787.5 -1760.21)"  d="M1214 2245C 1204 2237 1181 2230 1164 2230C 1071 2229 870 2155 747 2077C 630 2002 487 1842 422 1715C 385 1640 343 1505 330 1418C 323 1369 313 1330 309 1330C 296 1330 290 1299 300 1279C 313 1256 353 1254 370 1275C 381 1288 380 1296 370 1318C 328 1406 428 1697 555 1855C 693 2027 890 2147 1103 2187C 1179 2201 1194 2201 1231 2186C 1247 2179 1280 2202 1280 2220C 1280 2230 1251 2260 1241 2260C 1237 2260 1225 2253 1214 2245zM359 1316C 375 1296 362 1270 335 1270C 308 1270 299 1287 312 1313C 322 1333 343 1335 359 1316z" stroke="none" fill="currentcolor" fill-rule="nonzero" />
			</g>
			<g transform="matrix(0.13 0 -0 -0.13 87.78 -106.99)">
                <path transform="matrix(1 0 0 1 -1910.56 -2049.5)"  d="M1716 2208C 1705 2193 1704 2182 1711 2168C 1721 2149 1722 2148 1760 2146C 1816 2142 2024 1973 2036 1922C 2049 1868 2089 1851 2110 1891C 2122 1913 2109 1940 2086 1940C 2080 1940 2046 1970 2010 2006C 1974 2043 1910 2097 1868 2127C 1825 2157 1790 2185 1790 2190C 1790 2203 1760 2230 1745 2230C 1737 2230 1724 2220 1716 2208zM1774 2198C 1788 2176 1762 2146 1737 2156C 1720 2162 1714 2200 1727 2214C 1738 2225 1762 2217 1774 2198z" stroke="none" fill="currentcolor" fill-rule="nonzero" />
			</g>
			<g transform="matrix(0.13 0 -0 -0.13 -2.25 -94.9)">
                <path transform="matrix(1 0 0 1 -1235.3 -1958.81)"  d="M1148 2149C 986 2128 810 2042 671 1915C 629 1877 586 1840 575 1834C 550 1820 549 1781 573 1766C 596 1751 624 1769 628 1803C 633 1846 773 1969 877 2023C 1029 2101 1090 2115 1275 2115C 1415 2115 1444 2112 1510 2091C 1660 2044 1822 1946 1852 1883C 1861 1863 1873 1855 1890 1855C 1911 1855 1915 1860 1915 1885C 1915 1906 1910 1916 1896 1918C 1886 1920 1846 1946 1807 1977C 1622 2120 1391 2180 1148 2149zM615 1800C 615 1782 609 1774 592 1772C 579 1770 567 1775 564 1784C 554 1810 568 1832 592 1828C 609 1826 615 1818 615 1800z" stroke="none" fill="currentcolor" fill-rule="nonzero" />
			</g>
			<g transform="matrix(0.13 0 -0 -0.13 -47.57 1.27)">
                <path transform="matrix(1 0 0 1 -895.38 -1237.54)"  d="M894 1995C 886 1982 858 1959 833 1944C 808 1930 757 1888 720 1851C 608 1741 532 1604 494 1444C 471 1347 479 1152 510 1055C 606 752 846 541 1155 489C 1194 482 1235 471 1247 463C 1276 444 1310 461 1310 495C 1310 529 1276 546 1246 527C 1226 514 1213 513 1149 525C 851 578 621 786 538 1077C 506 1189 506 1372 539 1485C 571 1599 634 1711 713 1796C 787 1876 895 1954 918 1945C 938 1937 960 1958 960 1985C 960 2006 943 2020 919 2020C 914 2020 903 2009 894 1995zM1284 524C 1304 516 1305 486 1285 470C 1256 446 1225 491 1252 518C 1266 532 1264 532 1284 524z" stroke="none" fill="currentcolor" fill-rule="nonzero" />
			</g>
			<g transform="matrix(0.13 0 -0 -0.13 -30.42 -30.28)">
                <path transform="matrix(1 0 0 1 -1024.04 -1474.14)"  d="M1185 1994C 1174 1992 1140 1985 1110 1979C 1033 1964 886 1891 818 1833C 744 1771 670 1677 631 1598C 558 1449 539 1255 581 1100C 592 1056 599 1013 596 1004C 587 980 608 950 635 950C 679 950 692 1012 650 1025C 596 1042 575 1338 618 1475C 685 1689 859 1868 1060 1931C 1197 1974 1363 1977 1428 1938C 1454 1922 1463 1920 1475 1930C 1507 1956 1479 2012 1445 1992C 1435 1987 1410 1987 1386 1991C 1344 1999 1221 2001 1185 1994zM659 1006C 675 986 662 960 635 960C 622 960 609 967 606 976C 599 995 616 1020 635 1020C 642 1020 653 1014 659 1006z" stroke="none" fill="currentcolor" fill-rule="nonzero" />
			</g>
			<g transform="matrix(0.13 0 -0 -0.13 -7.15 -5.76)">
                <path transform="matrix(1 0 0 1 -1198.57 -1290.29)"  d="M1214 1930C 1209 1925 1169 1911 1126 1900C 1009 1870 928 1826 841 1744C 706 1617 639 1467 638 1289C 637 1171 651 1108 700 1005C 810 776 1026 640 1279 640C 1416 640 1553 686 1659 766C 1683 785 1711 800 1720 800C 1764 800 1775 872 1732 878C 1704 882 1680 864 1680 840C 1680 815 1594 754 1512 720C 1318 638 1096 663 928 783C 512 1081 622 1722 1113 1865C 1151 1876 1184 1879 1214 1875C 1251 1869 1260 1871 1268 1887C 1288 1923 1243 1959 1214 1930zM1745 840C 1745 807 1705 799 1695 829C 1686 857 1697 872 1723 868C 1739 866 1745 858 1745 840z" stroke="none" fill="currentcolor" fill-rule="nonzero" />
			</g>
			<g transform="matrix(0.13 0 -0 -0.13 -130.81 -13.78)">
                <path transform="matrix(1 0 0 1 -271.08 -1350.41)"  d="M317 1894C 313 1890 310 1876 310 1863C 310 1850 295 1809 276 1772C 151 1521 127 1224 210 950C 223 909 232 861 231 843C 229 816 233 810 254 804C 306 791 331 846 286 875C 262 891 234 972 210 1094C 193 1180 195 1380 214 1481C 243 1634 321 1823 361 1833C 393 1841 384 1894 350 1898C 336 1900 321 1898 317 1894zM295 860C 306 842 289 810 269 810C 247 810 235 832 244 855C 250 872 286 875 295 860z" stroke="none" fill="currentcolor" fill-rule="nonzero" />
			</g>
			<g transform="matrix(0.13 0 -0 -0.13 56.09 30.89)">
                <path transform="matrix(1 0 0 1 -1672.83 -1015.42)"  d="M1792 1268C 1777 1253 1776 1218 1790 1204C 1802 1192 1779 1107 1745 1039C 1692 935 1578 826 1515 818C 1491 816 1485 810 1485 791C 1485 778 1493 762 1503 756C 1518 747 1525 749 1540 768C 1550 780 1581 807 1607 828C 1709 905 1786 1016 1816 1128C 1826 1165 1840 1199 1847 1203C 1876 1221 1855 1280 1820 1280C 1811 1280 1799 1275 1792 1268zM1838 1261C 1859 1248 1848 1215 1822 1212C 1798 1208 1784 1230 1794 1255C 1800 1272 1818 1274 1838 1261z" stroke="none" fill="currentcolor" fill-rule="nonzero" />
			</g>
			<g transform="matrix(0.13 0 -0 -0.13 95.41 62.23)">
                <path transform="matrix(1 0 0 1 -1967.76 -780.36)"  d="M2265 1264C 2255 1254 2252 1241 2256 1231C 2293 1146 2231 914 2126 746C 2072 659 1970 546 1889 484C 1810 423 1664 347 1640 353C 1595 365 1574 310 1613 286C 1628 277 1635 279 1648 296C 1658 307 1691 328 1721 342C 2010 474 2237 773 2295 1096C 2303 1145 2318 1199 2327 1215L2327 1215L2343 1245L2321 1223C 2310 1211 2293 1203 2286 1206C 2263 1215 2258 1242 2275 1259C 2288 1272 2294 1272 2310 1262C 2321 1255 2330 1254 2330 1259C 2330 1268 2307 1280 2290 1280C 2285 1280 2274 1273 2265 1264z" stroke="none" fill="currentcolor" fill-rule="nonzero" />
			</g>
			<g transform="matrix(0.13 0 -0 -0.13 -44.49 34.1)">
                <path transform="matrix(1 0 0 1 -918.53 -991.3)"  d="M713 1253C 698 1248 696 1204 709 1196C 714 1193 728 1157 740 1118C 752 1078 777 1020 795 990C 836 922 940 821 999 791C 1023 778 1052 757 1064 743C 1089 714 1125 718 1135 751C 1142 771 1126 800 1108 800C 1103 800 1106 793 1114 784C 1133 765 1134 758 1118 742C 1100 724 1082 727 1070 749C 1063 764 1064 772 1076 784C 1090 799 1089 800 1067 800C 1033 800 929 871 880 928C 815 1003 765 1113 765 1180C 765 1253 753 1270 713 1253z" stroke="none" fill="currentcolor" fill-rule="nonzero" />
			</g>
			<g transform="matrix(0.13 0 -0 -0.13 8.73 77.43)">
                <path transform="matrix(1 0 0 1 -1317.64 -666.36)"  d="M2134 999C 2130 992 2126 968 2125 945C 2122 856 1983 662 1847 555C 1680 423 1494 358 1280 357C 1084 356 915 408 750 520C 638 596 500 750 500 799C 500 825 471 841 448 829C 420 814 426 793 474 742C 499 716 551 660 592 617C 841 350 1217 256 1575 371C 1809 446 2010 618 2126 842C 2151 891 2175 930 2180 930C 2198 930 2211 969 2200 990C 2188 1012 2146 1017 2134 999zM2194 984C 2201 965 2184 940 2165 940C 2146 940 2129 965 2136 984C 2139 993 2152 1000 2165 1000C 2178 1000 2191 993 2194 984z" stroke="none" fill="currentcolor" fill-rule="nonzero" />
			</g>
			<g transform="matrix(0.13 0 -0 -0.13 11.77 112.53)">
                <path transform="matrix(1 0 0 1 -1340.43 -403.11)"  d="M2112 628C 2105 621 2100 606 2100 594C 2100 566 1967 439 1869 374C 1780 315 1648 257 1539 228C 1467 209 1429 206 1275 205C 1108 205 1088 207 995 233C 940 249 859 279 815 300C 713 349 570 449 570 472C 570 497 533 515 514 499C 495 483 500 450 522 439C 555 422 567 415 643 362C 842 225 1026 168 1275 167C 1426 166 1492 176 1616 216C 1779 269 1922 356 2060 485C 2101 523 2144 560 2155 566C 2179 580 2186 606 2170 625C 2155 643 2128 644 2112 628zM2165 600C 2165 582 2159 574 2142 572C 2129 570 2117 575 2114 584C 2104 610 2118 632 2142 628C 2159 626 2165 618 2165 600z" stroke="none" fill="currentcolor" fill-rule="nonzero" />
			</g>
		</g>
	</g>
</g>
</svg>`
