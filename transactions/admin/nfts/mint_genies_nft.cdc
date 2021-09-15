import NonFungibleToken from "../../../contracts/NonFungibleToken.cdc"
import Genies from "../../../contracts/Genies.cdc"

transaction(recipientAddress: Address, editionID: UInt32) {
    
    // local variable for storing the minter reference
    let minter: &{Genies.NFTMinter}
    let recipient: &{Genies.GeniesNFTCollectionPublic}

    prepare(signer: AuthAccount) {
        // borrow a reference to the NFTMinter resource in storage
        self.minter = signer.getCapability(Genies.MinterPrivatePath)
            .borrow<&{Genies.NFTMinter}>()
            ?? panic("Could not borrow a reference to the NFT minter")

        // get the recipients public account object
        let recipientAccount = getAccount(recipientAddress)

        // borrow a public reference to the receivers collection
        self.recipient = recipientAccount.getCapability(Genies.CollectionPublicPath)
            .borrow<&{Genies.GeniesNFTCollectionPublic}>()
            ?? panic("Could not borrow a reference to the collection receiver")
    }

    execute {
        // mint the NFT and deposit it to the recipient's collection
        let geniesNFT <- self.minter.mintNFT(editionID: editionID)
        self.recipient.deposit(token: <- (geniesNFT as @NonFungibleToken.NFT))
    }
}
