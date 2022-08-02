import NonFungibleToken from "../../../contracts/NonFungibleToken.cdc"
import AllDaySeasonal from "../../../contracts/AllDaySeasonal.cdc"

transaction(recipientAddress: Address, editionID: UInt64) {
    
    // local variable for storing the minter reference
    let minter: &{AllDaySeasonal.NFTMinter}
    let recipient: &{AllDaySeasonal.SeasonalNFTCollectionPublic}

    prepare(signer: AuthAccount) {
        // borrow a reference to the NFTMinter resource in storage
        self.minter = signer.getCapability(AllDaySeasonal.MinterPrivatePath)
            .borrow<&{AllDay.NFTMinter}>()
            ?? panic("Could not borrow a reference to the NFT minter")

        // get the recipients public account object
        let recipientAccount = getAccount(recipientAddress)

        // borrow a public reference to the receivers collection
        self.recipient = recipientAccount.getCapability(AllDaySeasonal.CollectionPublicPath)
            .borrow<&{AllDaySeasonal.SeasonalNFTCollectionPublic}>()
            ?? panic("Could not borrow a reference to the collection receiver")
    }

    execute {
        // mint the NFT and deposit it to the recipient's collection
        let seasonalNFT <- self.minter.mintNFT(editionID: editionID)
        self.recipient.deposit(token: <- (seasonalNFT as @NonFungibleToken.NFT))
    }
}

