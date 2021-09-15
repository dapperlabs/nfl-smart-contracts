import Genies from "../../contracts/Genies.cdc"

// This scripts returns the number of Genies currently in existence.

pub fun main(): UInt64 {    
    return Genies.totalSupply
}
