const algosdk = require('algosdk');
require('dotenv').config({ path: './.env' });
const byteArrayToBase64 = require('./utils')
const AppService = require('./api-service')

async function mintToken() {
    const baseServer = 'https://testnet-algorand.api.purestake.io/ps2'
    const port = '';
    const token = { 'X-API-Key': process.env.PURESTAKE_API_TOKEN }

    try {

        const msig_addr = "HURKQWPXZDRBRW6ZI7RE2AUKBBOZW6Z3LVGJEYVEVYZORBNQY6DJPJ5YOY"
        const algodClient = new algosdk.Algodv2(token, baseServer, port);
        let params = await algodClient.getTransactionParams().do();

        const creator = msig_addr;
        const defaultFrozen = false;
        const unitName = "NFT";
        const assetName = "NFT";
        const assetURL = "https://www.algomint.io/";
        let note = undefined;
        const manager = msig_addr;
        const reserve = msig_addr;
        const freeze = msig_addr;
        const clawback = msig_addr;
        let assetMetadataHash = undefined;
        const total = 1;
        const decimals = 0;
        const txn = algosdk.makeAssetCreateTxnWithSuggestedParams(
            creator,
            note,
            total,
            decimals,
            defaultFrozen,
            manager,
            reserve,
            freeze,
            clawback,
            unitName,
            assetName,
            assetURL,
            assetMetadataHash,
            params,
        );


        //
        // Line 51 & 52 aswell as line 57 are the lines of interest in this file
        //
        let binaryMultisigTx = txn.toByte();
        let base64MultisigTx = byteArrayToBase64(binaryMultisigTx);
        console.log("-----RawTxn base64 -----")
        console.log(base64MultisigTx)
        let txID = "test-id-txn"
        let numSign = 2;
        console.log("-----txtID-----")
        console.log(txID)
        console.log("-----numSign-----")
        console.log(numSign)
        // not compatible with new api 
        //const response = await AppService(txID, base64MultisigTx, numSign);
        //console.log(response);
    }
    catch (err) {
        console.log("err", err);
    }
    process.exit();
};


mintToken();