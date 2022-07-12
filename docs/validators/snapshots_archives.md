<!--
order: 6
-->

# Snapshots & Archive Nodes

Quickly sync your node with bitos using a snapshot or serve queries for prev versions using archive nodes {synopsis}

## List of Snapshots and Archives

Below is a list of publicly available snapshots that you can use to sync with the bitos mainnet and
archived [9001-1 mainnet](https://github.com/tharsis/mainnet/tree/main/bitos_9001-1):

<!-- markdown-link-check-disable -->
:::: tabs
::: tab Snapshots

| Name        | URL                                                                     |
| -------------|------------------------------------------------------------------------ |
| `Staketab`   | [github.com/staketab/nginx-cosmos-snap](https://github.com/staketab/nginx-cosmos-snap/blob/main/docs/bitos.md) |
| `Polkachu`   | [polkachu.com](https://www.polkachu.com/tendermint_snapshots/bitos)                   |
| `Nodes Guru` | [snapshots.nodes.guru/bitos_9001-2/](snapshots.nodes.guru/bitos_9001-2/)                   |
:::
::: tab Archives
<!-- markdown-link-check-disable -->

| Name           | URL                                                                             |
| ---------------|---------------------------------------------------------------------------------|
| `Nodes Guru`   | [snapshots.nodes.guru/bitos_9001-1](https://snapshots.nodes.guru/bitos_9001-1/)                                    |
| `Polkachu`     | [polkachu.com/tendermint_snapshots/bitos](https://www.polkachu.com/tendermint_snapshots/bitos)                           |
| `Forbole`      | [bigdipper.live/bitos_9001-1](https://s3.bigdipper.live.eu-central-1.linodeobjects.com/bitos_9001-1.tar.lz4) |
:::
::::

To access snapshots and archives, follow the process below (this code snippet is to access a snapshot of the current network, `bitos_9001-2`, from Nodes Guru):

```bash
cd $HOME/.bitosd/data
wget https://snapshots.nodes.guru/bitos_9001-2/bitos_9001-2-410819.tar
tar xf bitos_9001-2-410819.tar
```
