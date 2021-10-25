## Installation

1. Q: What is the working directory to run `stratos-chain-test`?

   A: No matter you use linux or Mac, the working folder is the folder that has the `stchaind` and `stchaincli` binary files.

   **For Linux users**, you can make a working directory(e.g., `testnet`) and directly download the pre-built binary files from

      ```bash
      mkdir testnet
      cd testnet
      wget https://github.com/stratosnet/stratos-chain/releases/download/v0.5.0/stchaincli
      wget https://github.com/stratosnet/stratos-chain/releases/download/v0.5.0/stchaind
      ```
      into `testnet` and follow the other operations inside this folder. 
      
      `testnet` is your working directory.

   **For self-build users**, make a folder(e.g., `testnet`). Then, enter `testnet` and clone the source

      ```bash
      cd testnet
      git clone https://github.com/stratosnet/stratos-chain.git
      cd stratos-chain
      git checkout v0.5.0
      make build
      ```
      After `make build`, you will find the `stchaind` and `stchaincli` binary files in `testnet/stratos-chain/build`.

     `build` is your working directory, and you can do other operations inside this folder.

     After installation, your working directory looks like
     ```bash
     .
     ├── config
     ├── data
     ├── stchaincli
     └── stchaind
     ```
     The `config` folder
     ```bash
     .
     ├── addrbook.json
     ├── app.toml
     ├── config.toml
     ├── genesis.json
     ├── node_key.json
     └── priv_validator_key.json
     ```
   In `config` folder:

   `addrbook.json` stores peer addresses.

   `app.toml` contains the default settings required for `app`.

   `config.toml` contains various options pertaining to the `stratos-chain` configurations.

   `genesis.json` defines the initial state upon genesis of `stratos-chain`.

   `node_key.json` contains the node private key and should thus be kept secret.

   `priv_validator_key.json` contains the validator address, public key and private key, and should thus be kept secret.

## Configurations

1. Q: How to change `moniker` in `config.toml`?

   A: In your working directory, use an editor to open `config/config.toml`. Search `moniker`(usually located in Line #16), give your node a preferred name inside the quotes, like "EricStratos" here.
      ```bash
      # A custom human readable name for this node
      moniker = "EricStratos"
      ``` 

## Run the node

1. Q: `ERROR: Dialing failed (attempts: 12): auth failure: secret conn failed: read tcp 192.168.23.128:49286->23.88.62.88:26656: i/o timeout module=pex addr=8571c63a215968708df92d4e753932a83fa7d1a9@23.88.62.88:26656`. How to resolve it?

   A: Many participants reported this error. It happens when you start your node or while the node is running. Actually, this error message is not strange in a p2p network and may be caused by the P2P address handling. You can find lots of information addressing this error online, such as [this](https://github.com/tendermint/tendermint/issues/3716). Some people said, "I just had to leave it running a while longer, it started syncing after all the errors". We are working on it, trying to find a workaround. Currently, you can check your network connection with `curl localhost:26657/net_info | grep moniker`, which will show the node that you connect with. if you have some nodes connected, your node is fine.

2. Q: Couldn't read `GenesisDoc`(or `config.toml`) file: no such file or directory?
   
   A: All `stratos-chain` commands need to be run in the **working directory**. First, check if you started node in the working directory; second, check your `config` folder if you already moved `genesis.json` and `config.toml` here. If this error is still there, re-download these two files, change `moniker` in `config.toml` file and move both of them into `config` folder.


3. Q: Unknown address: account st1xxx... does not exist?
   
   A: This error happens when:
   * your node has not caught up with the current blockchain height, that is, the synchronization process has not finished yet.
   * your node has no or very small amount of tokens. Note that 1Stos=1000,000,000ustos.
   * the flag `--chain-id` in your command is incorrect. The current `chain-id` is shown [on this page](https://big-dipper-test.thestratos.org/) right next to the search bar at the top of the page.

4. Q: How can I get test tokens and where?
   
   A: You can get test tokens from our `faucet` service using the following command
      ```bash
      curl -X POST https://faucet-test.thestratos.org/faucet/<your-wallet-address>
      ```
     For example: in your terminal,
     ```bash
     curl -X POST https://faucet-test.thestratos.org/faucet/st1gwtcnptte6fpxck3f9xs45ufrru9sz2500edn8
     ```
   Note:

       * 1 stos = 1000000000 ustos
       * By default, faucet will send 100stos(100000000000ustos) to the given wallet address
       * maximum 3 faucet requests to arbitrary wallet address from a single IP within an hour
       * maximum 1 faucet request to a fixed wallet address within an hour


5  Q: Doc, Doc, Doc?

   A: Yes, please find them at

   * [How to connect Stratos Chain Testnet](https://github.com/stratosnet/stratos-chain-testnet)

   * [Stratos Chain README](https://github.com/stratosnet/stratos-chain)

   * [Stratos Chain `stchaincli` Commands](https://github.com/stratosnet/stratos-chain/wiki/Stratos-Chain-%60stchaincli%60-Commands)

   * [Stratos Chain `stchaind` Commands](https://github.com/stratosnet/stratos-chain/wiki/Stratos-Chain-%60stchaind%60-Commands)

   * [Stratos Chain REST APIs](https://github.com/stratosnet/stratos-chain/wiki/Stratos-Chain-REST-APIs)

   * [Stratos SDS README](https://github.com/stratosnet/sds)

   For Chinese version(中文版):

   * [Stratos Chain 介绍](https://github.com/stratosnet/stratos-chain/wiki/Stratos-Chain-%E4%BB%8B%E7%BB%8D)

   * [Stratos chain testnet 测试网说明](https://github.com/stratosnet/stratos-chain-testnet/wiki/Stratos-chain-testnet-%E6%B5%8B%E8%AF%95%E7%BD%91%E8%AF%B4%E6%98%8E)

   * [Stratos SDS 使用说明](https://github.com/stratosnet/sds/wiki/Stratos-SDS-%E4%BD%BF%E7%94%A8%E8%AF%B4%E6%98%8E)