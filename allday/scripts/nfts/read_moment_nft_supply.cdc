import AllDay from "../../contracts/AllDay.cdc"

// This scripts returns the number of AllDay currently in existence.

pub fun main(): UInt64 {    
    return AllDay.totalSupply
}

