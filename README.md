# Thing-A-Day
=================================

An open source web service to return a semi-randomly selected set of
items in a database, so that:

  * no items are duplicated in the same set (unless the set size exceeds the number of items available)
  * items are cached for an interval (e.g. a day) so you get the same items for that whole interval
  * items are non re-used that have been recently used, unless there aren't enough items to avoid that


