import NonFungibleToken from "NonFungibleToken"
import AllDay from "AllDay"

// This transaction transfers a AllDay NFT from one account to another.

transaction(recipientAddress: Address, withdrawID: UInt64) {
    prepare(signer: auth(BorrowValue) &Account) {
        
        // get the recipients public account object
        let recipient = getAccount(recipientAddress)

        // borrow a reference to the signer's NFT collection
        let collectionRef = signer.storage.borrow<auth(NonFungibleToken.Withdraw) &AllDay.Collection>(from: AllDay.CollectionStoragePath)
            ?? panic("Could not borrow a reference to the owner's collection")

        // borrow a public reference to the receivers collection
        let depositRef = recipient..capabilities.get<&AllDay.Collection>(AllDay.CollectionPublicPath)!.borrow()!

        // withdraw the NFT from the owner's collection
        let nft <- collectionRef.withdraw(withdrawID: withdrawID)

        // Deposit the NFT in the recipient's collection
        depositRef.deposit(token: <-nft)
    }
}
