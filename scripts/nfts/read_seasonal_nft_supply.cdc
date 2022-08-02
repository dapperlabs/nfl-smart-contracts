import AllDaySeasonal from "../../contracts/AllDaySeasonal.cdc"

// This scripts returns the number of AllDay currently in existence.

pub fun main(): UInt64 {    
    return AllDaySeasonal.totalSupply
}

