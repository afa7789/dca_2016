
# DCA_2016

A friend told me XMR was not as performant as BTC, I checked how much I wanted to check I could have gain from it compared to BTC. 

So I used my package of price fetching from CSV build in satsukashii ( check my gitub ) https://github.com/afa7789/satsukashii

Basically I get the prices of both in the time series and after that I check how much we would buy with 10 dollars every friday from 2016 to now.

## Usage

```bash
    go run main.go
```

## Output

```bash
Investment period: 2016-01-01 to 2025-05-01
Total invested weekly: $4860.00
Number of weeks: 486
Value per week: $10.00
Bitcoin:
	Total coins accumulated: 1.497135,
	Value on 2025-05-01: $141087.07
Monero:
	Total coins accumulated: 414.051742,
	Value on 2025-05-01: $116901.42
Amount got per week for:
	Bitcoin: $290.302620,
	Monero: $240.537907
Total value of:
	Bitcoin: $141087.07,
	Monero: $116901.42
```

![alt text](assets/image.png)

