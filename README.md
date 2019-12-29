A command line runescape [gem bag](https://runescape.fandom.com/wiki/Gem_bag_upgrade) calculater written in go.

```
$ go run main.go -b 1 -g 2 -r 3 -w 4 -p 5
  ITEM CODE |       NAME        | QUANTITY | PRICE | AMOUNT  
------------+-------------------+----------+-------+---------
       1623 | Uncut sapphire    |        1 |   362 |    362  
       1621 | Uncut emerald     |        2 |   716 |   1432  
       1619 | Uncut ruby        |        3 |  1195 |   3585  
       1617 | Uncut diamond     |        4 |  2452 |   9808  
       1631 | Uncut dragonstone |        5 | 11527 |  57635  
------------+-------------------+----------+-------+---------
                                             TOTAL | 72822   
                                           --------+---------
```
