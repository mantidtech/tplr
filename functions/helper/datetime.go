package helper

import "time"

// Now is defined as a variable so that it can be overridden as required (e.g. unit testing)
var Now = time.Now
