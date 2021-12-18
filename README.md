# connpassSearcher

connpassSearcher gets ICT event informations from https://connpass.com/

## Installation

When you have installed the Go, Please execute following `go get` command:

```sh
go get -u github.com/qt-luigi/connpassSearcher
```

## Usage

```sh
$ ./connpassSearcher -h
Usage of ./connpassSearcher:
  -k string
    	keyword for event (default empty)
  -m int
    	month to hold event (default this month)
  -y int
    	year to hold event (default this year)
```
## Example

```sh
$ date
2021年 12月18日 土曜日 16時41分19秒 JST

$ ./connpassSearcher -k Go (other parameters are default: -k 2021 -m 12)
12/1(水) 19:00〜21:30 [オンライン] Go Language Specification 輪読会 #32 https://gospecreading.connpass.com/event/231503/
...
12/26(日) 14:00〜17:00 [オンライン] 香川Go言語  わいわい会 (GAMEを作ろう)#3 https://gdgshikoku.connpass.com/event/233967/

$ ./connpassSearcher -k Go -y 2020 (other parameter is default: -m 12)
12/9(水) 19:00〜21:30 [オンライン] Go Language Specification輪読会 #11 https://gospecreading.connpass.com/event/197221/
...
12/26(土) 10:00〜12:00 [オンライン] 【オンライン】Women Who Go Tokyo ハンズオン - Go言語で作るインタプリタ#4 https://womenwhogo-tokyo.connpass.com/event/199316/

$ ./connpassSearcher -k Python (other parameters are default: -k 2021 -m 12)
results_available over 100
```

## Thanks

Using package: https://github.com/tenntenn/connpass

## License

MIT

## Author

Ryuji Iwata

## Note

This tool is mainly using by myself. :-)
