import NonFungibleToken from "../../../contracts/NonFungibleToken.cdc"
import Showdown from "../../../contracts/Showdown.cdc"

transaction(recipientAddress: Address, editionID: UInt32) {
    
    // local variable for storing the minter reference
    let minter: &{Showdown.NFTMinter}
    let recipient: &{Showdown.MomentNFTCollectionPublic}

    prepare(signer: AuthAccount) {
        // borrow a reference to the NFTMinter resource in storage
        self.minter = signer.getCapability(Showdown.MinterPrivatePath)
            .borrow<&{Showdown.NFTMinter}>()
            ?? panic("Could not borrow a reference to the NFT minter")

        // get the recipients public account object
        let recipientAccount = getAccount(recipientAddress)

        // borrow a public reference to the receivers collection
        self.recipient = recipientAccount.getCapability(Showdown.CollectionPublicPath)
            .borrow<&{Showdown.MomentNFTCollectionPublic}>()
            ?? panic("Could not borrow a reference to the collection receiver")
    }

    execute {
        // mint the NFT and deposit it to the recipient's collection
        let momentNFT <- self.minter.mintNFT(editionID: editionID)
        self.recipient.deposit(token: <- (momentNFT as @NonFungibleToken.NFT))
    }
}
