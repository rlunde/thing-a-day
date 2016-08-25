# thing-a-day

DISCLAIMER: I'm not expecting anyone to look at this project, or use it, or contribute to it, for
quite a while. I'm using it as a learning exercise more than anything else at the moment, with the
goal of eventually knowing go, backbone, bootstrap, and so on. I will remove this disclaimer when
I think the project has reached MVP (minimum viable project) stage. That won't be for a while.

This is a web API to provide periodic random selections of data for websites, tweets, etc. 

The basic idea is that lots of people (not just me) want to make a "semi-static" site, or one
that has static elements that change periodically, *without any action by the author*. The best 
way to describe it may be an example. 

For many years, my humorscope.com site chugged away happily in the background,
and produced a new main page every day with a "prediction" for each of the 12 zodiac signs, which were randomly
plucked from a database I made with around two thousand silly predictions that I had written over the
years. What I wanted it to do was:

* every item should be unique for that day (e.g. Taurus and Gemini should not have the same thing.)
* no item should be re-used too often, so try not to use one that had been used for any sign for the past two weeks.
* if I added new items to the database -- things that had never been used -- make it slightly more likely to pick one of those.

I also had a database of funny or inspirational quotes, and randomly picked one of those
per day.

This project is about generalizing that and setting it up so anyone can use it by adding
stuff to a database and using an API.

I'm redoing the humorscope site and making it use thing-a-day, and making the code open source
for that as well, so there's a simple example of a real site that uses this.

So:

 1. schedule a periodic task
   * generate a static web page
   * send a tweet
   * send an email
 2. use a simple DSL (domain-specific language)
   * choose N random items from a database that are all unique
   * refer back to an item that was previously chosen
   * don't re-use items that have been recently chosen
 3. prefer never-used (or new) items in the database
   * but don't use them all up at once


# License

This project is open source software.

See the LICENSE file for a link to the text of the license.
