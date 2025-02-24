import algosdk from "algosdk";
import configData from "../config/config.json";

async function initiateAlgodClient() {
  const baseServer = "https://testnet-algorand.api.purestake.io/ps2";
  const port = "";
  const token = { "X-API-Key": configData.PS_TOKEN };

  const algodClient = new algosdk.Algodv2(token, baseServer, port);

  return algodClient;
}

export async function checkMultiSigAddrFunds(addr: string): Promise<boolean> {
  const algodClient = await initiateAlgodClient();
  let accountInfo = await algodClient.accountInformation(addr).do();
  if (accountInfo.amount > 0) {
    console.log("Account balance: %d microAlgos", accountInfo.amount);
    return true;
  }
  console.error("Multisig Address Account has no fund");
  return false;
}

export default initiateAlgodClient;
