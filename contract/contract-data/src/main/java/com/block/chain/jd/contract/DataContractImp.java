package com.block.chain.jd.contract;

import com.jd.blockchain.contract.ContractEventContext;
import com.jd.blockchain.contract.EventProcessingAware;
import com.jd.blockchain.crypto.AsymmetricKeypair;
import com.jd.blockchain.crypto.Crypto;
import com.jd.blockchain.crypto.CryptoAlgorithm;
import com.jd.blockchain.crypto.SignatureFunction;
import com.jd.blockchain.ledger.BlockchainKeypair;
import com.jd.blockchain.ledger.Event;
import com.jd.blockchain.ledger.TypedKVEntry;

import utils.Bytes;

public class DataContractImp implements EventProcessingAware, DataContract {

    private ContractEventContext eventContext;


    @Override
    public String getDataName() {
        return "MNIST";
    }

    @Override
    public String getDataInfo() {
        return "The MNIST database of handwritten digits, available from this page, has a training set of 60,000 examples, and a test set of 10,000 examples. It is a subset of a larger set available from NIST. The digits have been size-normalized and centered in a fixed-size image.";
    }

    @Override
    public void setKVWithVersion(String address, String key, String value, long version) {
        eventContext.getLedger().dataAccount(Bytes.fromBase58(address)).setText(key, value, version);
    }

    @Override
    public void setKV(String address, String key, String value) {
        // 查询最新版本，初始为-1
        // 查询已提交区块数据，不包括此操作所在未提交区块的所有数据
        // TypedKVEntry[] entries = eventContext.getLedger().getDataEntries(eventContext.getCurrentLedgerHash(), address, key);
        // 可查询包括此操作所在未提交区块的所有数据
        TypedKVEntry[] entries = eventContext.getUncommittedLedger().getDataEntries(address, key);
        long version = -1;
        if (null != entries && entries.length > 0) {
            version = entries[0].getVersion();
        }
        eventContext.getLedger().dataAccount(Bytes.fromBase58(address)).setText(key, value, version);
    }

    @Override
    public String registerUser(String seed) {
        CryptoAlgorithm algorithm = Crypto.getAlgorithm("ed25519");
        SignatureFunction signFunc = Crypto.getSignatureFunction(algorithm);
        AsymmetricKeypair cryptoKeyPair = signFunc.generateKeypair(seed.getBytes());
        BlockchainKeypair keypair = new BlockchainKeypair(cryptoKeyPair.getPubKey(), cryptoKeyPair.getPrivKey());
        eventContext.getLedger().users().register(keypair.getIdentity());

        return keypair.getAddress().toBase58();
    }

    @Override
    public String registerDataAccount(String seed) {
        CryptoAlgorithm algorithm = Crypto.getAlgorithm("ed25519");
        SignatureFunction signFunc = Crypto.getSignatureFunction(algorithm);
        AsymmetricKeypair cryptoKeyPair = signFunc.generateKeypair(seed.getBytes());
        BlockchainKeypair keypair = new BlockchainKeypair(cryptoKeyPair.getPubKey(), cryptoKeyPair.getPrivKey());
        eventContext.getLedger().dataAccounts().register(keypair.getIdentity());

        return keypair.getAddress().toBase58();
    }

    @Override
    public String registerEventAccount(String seed) {
        CryptoAlgorithm algorithm = Crypto.getAlgorithm("ed25519");
        SignatureFunction signFunc = Crypto.getSignatureFunction(algorithm);
        AsymmetricKeypair cryptoKeyPair = signFunc.generateKeypair(seed.getBytes());
        BlockchainKeypair keypair = new BlockchainKeypair(cryptoKeyPair.getPubKey(), cryptoKeyPair.getPrivKey());
        eventContext.getLedger().eventAccounts().register(keypair.getIdentity());

        return keypair.getAddress().toBase58();
    }

    @Override
    public void publishEventWithSequence(String address, String topic, String content, long sequence) {
        eventContext.getLedger().eventAccount(Bytes.fromBase58(address)).publish(topic, content, sequence);
    }

    @Override
    public void publishEvent(String address, String topic, String content) {
        // 查询最新序号，初始为-1
        // 查询已提交区块数据，不包括此操作所在未提交区块的所有数据
        // Event event = eventContext.getLedger().getRuntimeLedger().getLatestEvent(address, topic);
        // 可查询包括此操作所在未提交区块的所有数据
        Event event = eventContext.getUncommittedLedger().getLatestEvent(address, topic);
        long sequence = -1;
        if (null != event) {
            sequence = event.getSequence();
        }
        eventContext.getLedger().eventAccount(Bytes.fromBase58(address)).publish(topic, content, sequence);
    }

    /**
     * 合约方法调用前操作
     *
     * @param eventContext
     */
    @Override
    public void beforeEvent(ContractEventContext eventContext) {
        this.eventContext = eventContext;
    }

    /**
     * 合约方法调用后操作
     *
     * @param eventContext
     * @param error
     */
    @Override
    public void postEvent(ContractEventContext eventContext, Exception error) {

    }

    /**
     * 交易
     *
     * @result model链接
     */
    public String trade() {
        return "dfenjogjre32344tljjroghfjej";
    }
}
