# NFL Smart Contracts

## NFL Contract Addresses
| Network   | Address                        |                                                                                     |
| ----------|:------------------------------:| -----------------------------------------------------------------------------------:|
| Testnet   |      4dfd62c88d1b6462          | [Flow View Source](https://flow-view-source.com/mainnet/account/0x4dfd62c88d1b6462) |
| Mainnet   |      e4cf4bdc1751c65d          | [Flow View Source](https://flow-view-source.com/testnet/account/0xe4cf4bdc1751c65d) |

## Entities

### Series
Series encompass periods of time and will be named using strings like: `Summer 2021` or `Series 3`. 
More that one series can be open at any given time, and in order for an Edition to be created, it must have a SeriesID.

**On Chain Fields**
- FlowID
- Name
- Active

**Transactions**
- CreateSeries: Mints a new series onto Flow
- CloseSeries: Stops any new Editions from using the specified series
### Sets
Sets are categories: `Greatest Touchdowns` or similar. Sets have a unique name.An Edition must have a SetID to be created.
Sets do not close and cannot be retired. Sets contain a dictionary of all the SetID/PlayID combinations that exist within
an Edition. This is checked everytime a new Edition is created to ensure they are unique.

**On Chain Fields**
- FlowID
- Name

**Transactions**
- CreateSet: Mints a new set onto Flow
### Plays
Plays contain the actual play metadata, including stats from NFL and Elias. 
This will contain Player, Team, and Game metadata some of which may be blank depending on the type of moment.

**On Chain Fields**
- FlowID
- Classification (Name TBC: example, PLAYER_GAME, TEAM_GAME, PLAYER_MELT, TEAM_MELT)
- Metadata (stored as a string map. This can technically be anything, but the agreeed upon fields are as follows)
  - PlayType
  - HomeTeamName
  - AwayTeamName
  - TeamName
  - GameDate
  - HomeTeamScore
  - AwayTeamScore
  - PlayerFirstName
  - PlayerLastName
  - PlayerPosition
  - PlayerNumber

**Transactions**
- CreatePlay: Mints a new Play on Flow
### Editions
Editions are the combination of a SeriesID, SetID, and PlayID and are what moments are minted out of.
They also have a Max and Current Edition size so we can specify how many moments can ever be minted from 
the edition. 

The MaxEditionSize is optional. If it is not set, moments can be minted unlimitedly. An Edition will close, if either of these things happen:
- The max number of moments are minted
- The CloseEdition transaction is used
- 
`MaxEditionSize` cannot be changed once it is set.

**Fields**
- FlowID
- SeriesID
- SetID
- PlayID
- MaxEditionSize
- Tier
- NumMinted

**Transactions**
- CreateEdition: Mints a new Edition on Flow.
- CloseEdition: Closes an Edition so no new moments can be minted from it. This is irreversible. The Edition is closed by setting the MaxEditionSize to the value of NumMinted.

### Moment NFT
Moments are minted out of editions. You can think of Editions as a "cookie cutter" for moments. The Serial Number is what makes each MomentNFT unique. These are the NFTs that will be sold in packs. 

**Fields**
- FlowID
- EditionID
- Serial Number

**Transactions**
- MintMomentNFT: Mints a moment out of an EditionID
