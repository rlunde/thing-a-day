# thing-a-day
A web API to provide periodic random selections of data for websites, tweets, etc. 

I've been meaning to do this for ever and ever, but never do. Hopefully I'll get it done.

The idea is that lots of people (not just me) want to make a semi-static site, or one
that has static elements that change periodically. The best way to describe it may be
an example. 

For years, my humorscope.com site chugged away happily in the background,
and produced a page with a "prediction" for each of the 12 zodiac signs, that was randomly
plucked from a database I made with around two thousand silly predictions. What I wanted it
to do was:

* every prediction should be unique for that day (e.g. Taurus and Gemini should not have the same thing.)
* no prediction should be re-used too often, so try not to use one that had been used for any sign for the past two weeks.
* if I added new items to the database, that had never been used, make it slightly more likely to pick one of those.

I also had a database of funny or inspirational quotes, and randomly picked one of those
per day.

This project is about generalizing that and setting it up so anyone can use it by adding
stuff to a database and using an API.

