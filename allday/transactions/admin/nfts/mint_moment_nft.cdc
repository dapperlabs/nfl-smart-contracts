import NonFungibleToken from "../../../contracts/NonFungibleToken.cdc"
import AllDay from "../../../contracts/AllDay.cdc"

transaction(recipientAddress: Address, editionID: UInt64, serialNumber: UInt64?) {
    
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

    execute {
        // mint the NFT and deposit it to the recipient's collection
        let momentNFT <- self.minter.mintNFT(editionID: editionID, serialNumber: serialNumber)
        self.recipient.deposit(token: <- (momentNFT as @NonFungibleToken.NFT))
    }
}

