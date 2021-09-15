# NFL Smart Contracts


## NFL Contract Addresses
TBC


## Entities

### Series
Series encompass periods of time and will be named using strings like: `Summer 2021` or `Series 3`. 
More that one series can be open at any given time, and in order for an Edition to be created, it must have a SeriesID.

**Fields**
- FlowID
- Name

**Transactions**
- MintSeries: Mints a new series onto Flow
- CloseSeries: Stops any new Editions from using the specified series

### Sets
Sets are categories of plays: `Greatest Touchdowns` or similar. Sets have a name and description.
There can be many sets but only one may be used on an Edition. An Edition must have a SetID to be created.

**Fields**
- FlowID
- Name
- Description

**Transactions**
- MintSet: Mints a new set onto Flow
- CloseSet: Stops any new Editions from using this set (?)
### Plays
Plays contain the actual play metadata, including stats from Sport Radar. This will contain Player, Team, and Game metadata

**Fields**
- FlowID
- Metadata

**Transactions**
- MintPlay: Mints a new Play on Flow


### Editions
Editions are the combination of a SeriesID, PlayID, and FlowID and are what moments are minted out of.
They also have a Max and Current Edition size so we can specify how many moments can ever be minted from 
the edition. 

The MaxEditionSize should be able to be added at any point and if empty allows for perpetual minting(for example, for editions we want to mint an unlimited number of). 
Once it exists it is locked and can't be changed. 
Moments are minted out of editions, given and EditionID and a number to mint.

**Fields**
- FlowID
- SeriesID
- SetID
- PlayID
- MaxEditionSize
- CurrentEditionSize
- Rarity

**Transactions**
- MintEdition: Mints a new Edition on Flow. It should check that no edition exists with the specific SetID/PlayID combination
- SetMaxEditionSizeFromCurrentSize: Should set the max edition size to whatever the current edition size is to avoid minting any more moments


### Moment NFT
Moments are minted out of editions and can be minted. These are the NFTs that will be sold in packs

**Fields**
- FlowID
- EditionID

**Transactions**
- MintMoments: Mints moments out of the EditionID. Can only mint up to the MaxEditionSize
