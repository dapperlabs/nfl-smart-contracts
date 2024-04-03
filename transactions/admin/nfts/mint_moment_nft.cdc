import NonFungibleToken from "NonFungibleToken"
import AllDay from "AllDay"

transaction(recipientAddress: Address, editionID: UInt64, serialNumber: UInt64?) {
    
    // local variable for storing the minter reference
    let minter: &AllDay.Admin
    let recipient: &AllDay.Collection

    prepare(signer: auth(BorrowValue) &Account) {
        // borrow a reference to the NFTMinter resource in storage
        self.minter = signer.storage.borrow<&AllDay.Admin>(from: AllDay.AdminStoragePath)
            ?? panic("Could not borrow a reference to the NFT minter")

        // get the recipients public account object
        let recipientAccount = getAccount(recipientAddress)

        // borrow a public reference to the receivers collection
        self.recipient = recipientAccount.capabilities.borrow<&AllDay.Collection>(AllDay.CollectionPublicPath)
            ?? panic("Could not borrow a reference to the collection receiver")
    }

    execute {
        // mint the NFT and deposit it to the recipient's collection
        let momentNFT <- self.minter.mintNFT(editionID: editionID, serialNumber: serialNumber)
        self.recipient.deposit(token: <- (momentNFT as @{NonFungibleToken.NFT}))
    }
}

