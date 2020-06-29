# GoRESP

this is my own version of the RESP protocol. it is 
compatible with the normal RESP protocol, but bas a 
few convenience features i threw into it also.

---

### normal RESP

- `\r\n` **delimiter**
- `+` **simple string** 
    - example: `+foo\r\n`
- `$` **bulk string**
    - example: `$3\r\nfoo\r\n`
- `:` **integer**
    - example: `:420\r\n`
- `-` **error**
    - example: `-foo\r\n`
- `$-1` **null**
    - example: `$-1\r\n`
    
---

### custom RESP

- `!` **nil** _shorthand for null_
    - example: `!\r\n`
- `.` **float**
    - example: `.69.666\r\n`