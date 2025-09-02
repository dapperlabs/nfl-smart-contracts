import AllDay from "AllDay"

// This script returns an Edition for an id number, if it exists.

access(all) fun main(editionID: UInt64): Result {
    return Result(editionData: AllDay.getEditionData(id: editionID))
}


access(all) struct Result {
    access(all) let id: UInt64
    access(all) let seriesID: UInt64
    access(all) let setID: UInt64
    access(all) let playID: UInt64
    access(all) var maxMintSize: UInt64?
    access(all) let tier: String
    access(all) var numMinted: UInt64
    access(all) let parallel: String

    view init (editionData: AllDay.EditionData) {
        self.id = editionData.id
        self.seriesID = editionData.seriesID
        self.setID = editionData.setID
        self.playID = editionData.playID
        self.maxMintSize = editionData.maxMintSize
        self.tier = editionData.tier
        self.numMinted = editionData.numMinted
        self.parallel = editionData.getParallel()
    }
}