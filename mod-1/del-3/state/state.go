package state

import (
    "fmt"
    "strings"
    "errors"
)

west := "HS rev korn kylling"
boat := ""
east := ""

func ViewState() string {
    return fmt.Sprintf("[%s ---V \\____%s_____/ _______%s______Ø---]", west, boat, east)
}

func PutInBoat(item string){
    // When putting item in the boat, we dont know where the item
    // currently is. It could be in West or East.
    // The main example for this is for homosapian (HS). We will have to
    // put HS in the boat from West and East so that it can fetch the item
    // from the other side
    if strings.Contains(west, item) {
        tmpWest = strings.ReplaceAll(west, item, "")
        tmpWest = strings.TrimSpace(tmpWest)

        if tmpWest == "rev kylling" || tmpWest == "kylling rev" {
            return errors.New("You can't leave 'rev' and 'kylling' together alone")
        }

        if len(strings.Fields(boat)) > 2 {
            return errors.New("You cant bring more than TWO items/things on the boat")
        }

        west = tmpWest
    }


    if strings.Contains(east, item) {
        tmpEast = strings.ReplaceAll(east, item, "")
        tmpEast = strings.TrimSpace(tmpEast)

        if tmpEast == "rev kylling" || tmpEast == "kylling rev" {
            return errors.New("You can't leave 'rev' and 'kylling' together alone")
        }

        if len(strings.Fields(boat)) > 2 {
            return errors.New("You cant bring more than TWO items/things on the boat")
        }

        east = tmpEast
    }

    return true
}

func CrossRiverTo(place string){
    if place == "west" {
        // All contents of boat gets put into west
        west = west + boat 
    } else if place == "east" {
        // All contents of boat gets put into east
        east = east + boat
    } else {
        return errors.New(fmt.Sprintf("'%s' is an invalid place", place))
    }
}
