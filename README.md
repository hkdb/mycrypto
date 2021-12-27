# MyCrypto - A Terminal-Based CryptoCurrency Portfolio Tracker
maintained by: hkdb@3df.io

![ScreenShot](readme/screenshot.png)

## SUMMARY
This is a terminal-based CryptoCurrency Portfolio Tracker that was made to track crypto assets across multiple wallets and platforms locally (as opposed to a cloud service). It grabs the latest data from CoinMarketCap.com via API and then calculates those numbers against the users existing portfolio which is manually filled-in by the user in CSV format. This tracker supports multiple portfolios.

## FEATURES

- Check price, 1h change, 24h change, 7d change, and last updated of all your coins
- Check price, 1h change, 24h change, 7d change, and last updated of a specific coin
- Check price, 1h change, 24h change, 7d change, and last updated and your coin holdings, avg buy/sell price, total cost, current value, change in $, and change in % of all your holdings separately and collectively. It will also calculate your holdings percentage between BTC, Big Caps (> $10 billion), and Small Caps (< $10 billion)
- All of the above bullet point with transaction history
- Check price, 1h change, 24h change, 7d change, and last updated and your coin holdings, avg buy/sell price, total cost, current value, change in $, and change in % of all your holdings for a specific coin
- All of the above bullet point with transaction history
- Track multiple portfolios

## CHANGELOG

- 12/26/2021 - Added portfolio.conf and custom path to it in settings.conf to define default portfolio and coins
- 12/24/2021 - Fixed terms
- 12/24/2021 - Multiple features - v0.1.0
   - Changed variable type from struct to map to store api responses for scalability
   - settings.conf to set what coins to query CMC for data
   - Layout improvements
   - Multiple portfolios support
   - Check price data only
   - Check price data for specific coin only
   - Removed show price data prior to showing portfolio data
   - Random small improvements
- 01/28/2021 - Decoupled code base and added optional show portfolio feature
- 01/16/2021 - first commit

## PLATFORMS
- Linux: Tested on Pop! OS 20.04 LTS
- Windows: Not tested but should work?
- Mac: Not tested but should work?

## DEPENDENCIES
- CoinMarketCap.com API account and API key
- Manually filling in all of your transactions into csv format

## INSTALLATION
1. Clone this repo:
   
   ```
   git clone https://github.com/hkdb/mycrypto.git
   ``` 
2. Enter into the repo directory, create a settings.conf file based on settings.conf-sample, and fill-in all the variables accordingly. Each variable has associating comments to help you figure out exactly what you need to fill in.
   
   ``` 
   CMC_API = Your own CMC API key
   PORTFOLIO_PATH =  The path to your portfolios so that you can put it anywhere you want.
   PORTFOLIO_CONF_PATH = The path to your portfolio.conf so you can set your default portfolio name and the coins to query CMC. I personally like using the same path as PORTFOLIO_PATH.
   ```
3. Make sure you create the paths that you specified and the CSV files of the coins you want to track based on sample-btc.csv

4. Create a portfolio.conf based on portfolio.conf-sample, and fill-in all the variables accordingly. Each variable has associating comments to help you figure out exactly what you need to fill in.

   ```
   COINS = The coins in your portfolio separated by commas
   PORTFOLIO = The name of your default portfolio. ie. hodl, trade, or whatever you want to name it.
   ```
5. Execute the following in terminal:
   
   ```
   ./install.sh
   ```
   This install script does the following:

   ```
   Creates a symlink to binary in ~/.local/bin
   ```

6. Create <PORTFOLIO>-<coin ticker>.csv (ie hodl-btc.csv) based on sample-btc.csv and fill-in all of your transactions in the portfolios folder
   
   ```
   Date,Coin,Cost,Fee,Total,Amount,Wallet
   12/04/2020,BTC,100.00 ,0.59,10.59,0.0030633,Electrum
   ```

## ADDING NEW COINS

1. Add the coin to portfolio.conf
2. Add a CSV file for the coin (ie. hodl-btc.csv)

## USAGE

Execute the following command from the terminal:

```
mycrypto
```
For more granular output, check the help menu:

```
mycrypto -h
```

```
Usage of mycrypto:
  -p string
    	USAGE: -p <portfolio>
    	
    	       (default "none")
  -q string
    	USAGE: -q <coin symbol>
    	
    	       (default "none")
  -s string
    	USAGE: -s <option>:
    			t: transactions
    			p: price
    	
    	       (default "none")
```


## BUILD
To build from source, execute the following from with in the repo directory:

```
go get
go build
```

## DISCLAIMER

This repo is sponsored by 3DF OSI and is maintained by volunteers. 3DF Limited, 3DF OSI, and its volunteers including the author in no way make any guarantees. Please use at your own risk!

To Learn more, please visit:

https://osi.3df.io

https://3df.io 
