# CLAUDE.md

NFL All Day smart contracts in Cadence for the Flow blockchain. Defines AllDay NFTs and PackNFT structures.

## Commands

- Test: `make test` (runs `lib/go/test`)
- CI: `make ci`
- Direct test: `cd lib/go/test && go test ./... -tags=no_cgo`

## Code Style

- Cadence contracts follow Flow NFT standards (NonFungibleToken, MetadataViews)
- Go tests use flow-emulator for local blockchain simulation
- Test assertions via testify/stretchr

## Testing

Tests run against Flow emulator in `lib/go/test/`:
```bash
cd lib/go/test
CGO_ENABLED=0 go test -tags=no_cgo ./...
```

## Project Structure

```
contracts/
  AllDay.cdc              # Main NFT contract
  PackNFT.cdc             # Pack NFT contract
  imports/                # Flow contract imports
scripts/                  # Cadence scripts
  badges/                 # Badge queries
  editions/               # Edition queries
  nfts/                   # NFT queries
  plays/                  # Play queries
  series/                 # Series queries
  sets/                   # Set queries
transactions/             # Cadence transactions
  admin/                  # Admin transactions
  user/                   # User transactions
lib/go/
  test/                   # Go integration tests
```

## Key Concepts

- Series > Sets > Plays > Editions > Moments
- Badges can be attached to NFTs
- PackNFT for distribution mechanics
