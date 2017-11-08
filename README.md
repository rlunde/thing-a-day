DISCLAIMER: I'm not expecting anyone to look at this project, or use it, or contribute to it, for
quite a while. I'm using it as a learning exercise more than anything else at the moment, with the
goal of eventually knowing go, docker, ember, and so on. I will remove this disclaimer when
I think the project has reached MVP (minimum viable project) stage. That won't be for a while.

NOTE: I'm not abandoning this project, but I've decided to spend some time really focusing on
another project (goodwin), which I will then use to try to help make this project.

================================

# Thing-A-Day
=================================

An open source web service to return a semi-randomly selected set of
items in a database, so that:

  * no items are duplicated in the same set (unless the set size exceeds the number of items available)
  * items are cached for an interval (e.g. a day) so you get the same items for that whole interval
  * items are non re-used that have been recently used, unless there aren't enough items to avoid that


