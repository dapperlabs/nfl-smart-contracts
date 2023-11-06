import NonFungibleToken from "../../../contracts/NonFungibleToken.cdc"
import AllDay from "../../../contracts/AllDay.cdc"

transaction(recipientAddress: Address, editionIDs: [UInt64], counts: [UInt64], serialNumbers: [UInt64?]) {
    
    // local variable for storing the minter reference
    let minter: &{AllDay.NFTMinter}
    let recipient: &{AllDay.MomentNFTCollectionPublic}

    prepare(signer: AuthAccount) {
        // borrow a reference to the NFTMinter resource in storage
        self.minter = signer.getCapability(AllDay.MinterPrivatePath)
            .borrow<&{AllDay.NFTMinter}>()
            ?? panic("Could not borrow a reference to the NFT minter")

        // get the recipients public account object
        let recipientAccount = getAccount(recipientAddress)

        // borrow a public reference to the receivers collection
        self.recipient = recipientAccount.getCapability(AllDay.CollectionPublicPath)
            .borrow<&{AllDay.MomentNFTCollectionPublic}>()
            ?? panic("Could not borrow a reference to the collection receiver")

    }

    pre {
        editionIDs.length == counts.length: "must pass arrays of same length"
        editionIDs.length == serialNumbers.length: "must pass arrays of same length"
    }

    execute {
        var i = 0
        while i < editionIDs.length {
            var remaining = counts[i]
            while remaining > 0 {
                // mint the NFT and deposit it to the recipient's collection
                self.recipient.deposit(token: <- self.minter.mintNFT(editionID: editionIDs[i], serialNumber: serialNumbers[i]))
                remaining = remaining - 1
            }
            i = i + 1
        }
    }
}

