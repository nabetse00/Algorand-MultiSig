import { RawTxnBackendResponseType } from "./../models/apiModels";
export class AppService {
  public async getRawTxn(txID: string): Promise<RawTxnBackendResponseType> {
    const response = await fetch(
      "http://localhost:8081/ms-multisig-db/v1/getrawtxn?id=" + txID
    );
    return await response.json();
  }
  //change to signed
  public async addSignedTxn(signer: any, txn: any, txid: any) {
    const response = await fetch(
      `http://localhost:8081/ms-multisig-db/v1/addsignedtxn`,
      {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          signer_public_address: signer,
          signed_transaction: txn,
          txn_id: txid,
        }),
      }
    );
    return await response.json();
  }

  public async addRawTxn(
    txid: string,
    txn: string,
    numSign: number,
    version: number,
    addrs: string[]
  ) {
    const response = await fetch(
      `http://localhost:8081/ms-multisig-db/v1/addrawtxn`,
      {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          txn_id: txid,
          transaction: txn,
          number_of_signs_required: Number(numSign),
          signers_addresses: addrs,
          version: version,
        }),
      }
    );
    return await response.json();
  }

  public async getTxnIds(addr: string) {
    const response = await fetch(
      `http://localhost:8081/ms-multisig-db/v1/gettxnids?addr=` + addr,
      {
        method: "GET",
        headers: { "Content-Type": "application/json" },
      }
    );
    return await response.json();
  }
}
