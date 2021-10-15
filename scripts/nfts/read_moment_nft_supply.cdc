import Showdown from "../../contracts/Showdown.cdc"

// This scripts returns the number of Showdown currently in existence.

pub fun main(): UInt64 {    
    return Showdown.totalSupply
}
