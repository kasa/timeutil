# timeutil
Uses a different arithmetic for time.Time and has some utilities functions.

It truncates the resulting date instead of normalizing like the standard library.

- 2010-01-28 + 1 month = 2010-02-28
- 2010-01-29 + 1 month = 2010-02-28
- 2010-01-30 + 1 month = 2010-02-28
- 2012-01-30 + 1 month = 2012-02-29 (2012 is a leap year).
