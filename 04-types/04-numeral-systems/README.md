# Numeral systems
As humans, we have many different systems for expressing the quantities of something. Using the decimal numeral system, we might say we have 7 oranges; or 14 carrots; or 42 dollars. Other numeral systems which are used in programming include the binary numeral system and the hexadecimal numeral system.

## Decimal Numeral System
- Add each number in the position of the decimal

| DECIMAL | ten millions   | millions       | hundred thousands | ten thousands  | thousands      | hundreds       | tens           | ones           |
|---------|----------------|----------------|-------------------|----------------|----------------|----------------|----------------|----------------|
|         | 10<sup>7</sup> | 10<sup>6</sup> | 10<sup>5</sup>    | 10<sup>4</sup> | 10<sup>3</sup> | 10<sup>2</sup> | 10<sup>1</sup> | 10<sup>0</sup> |
|    12   |                |                |                   |                |                |                | 1              | 2              |
|  42420  |                |                |                   | 4              | 2              | 4              | 2              | 0              |


## Binary Numeral System
- Add "1" in each column starting from the highest that will match upto the required number
- for 42 = 32 + 8 + 2

| Binary | one twenty eights | sixty fours | thirty twos | sixteens | eights | fours | twos | ones           |
|---------|----------------|----------------|-------------------|----------------|----------------|----------------|----------------|----------------|
|         |  2<sup>7</sup> | 2<sup>6</sup>  |   2<sup>5</sup>   | 2<sup>4</sup>  | 2<sup>3</sup>  | 2<sup>2</sup>  | 2<sup>1</sup>  | 2<sup>0</sup>  |
|    2    |                |                |                   |                |                |                | 1              | 0              |
|  42     |                |                |          1        |      0         |     1          |     0          | 1              | 0              |

## Hexadecimal Numeral System
- Adding digits to the number from 0-9 and use A-F for numbers 10-15

| HexaDecimal |  |   |  | 65,536's | 4096's | 256's | 16's | ones           |
|---------|----------------|----------------|-------------------|----------------|----------------|----------------|----------------|----------------|
|         |  16<sup>7</sup> | 16<sup>6</sup>  |   16<sup>5</sup>   | 16<sup>4</sup>  | 16<sup>3</sup>  | 16<sup>2</sup>  | 16<sup>1</sup>  | 16<sup>0</sup>  |
|    8    |                |                |                   |                |                |                | 0              | 8             |
|    10    |                |                |          1        |      0         |     1          |     0          | 0             | A              |
|    16    |                |                |          1        |      0         |     1          |     0          | 1             | 0              |
|    10    |                |                |          1        |      0         |     0          |     3          | 8             | F              |