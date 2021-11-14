package com.block.chain.jd.contract;

import com.jd.blockchain.contract.Contract;
import com.jd.blockchain.contract.ContractEvent;

@Contract
public interface DataContract {

    /**
     * 查找对应model名称
     */
    @ContractEvent(name = "getDataName")
    String getDataName();

    /**
     * 查找model具体信息
     */
    @ContractEvent(name = "getDataInfo")
    String getDataInfo();

    /**
     * 设置KV
     *
     * @param address 数据账户地址
     * @param key     键
     * @param value   值
     * @param version 版本
     */
    @ContractEvent(name = "setKVWithVersion")
    void setKVWithVersion(String address, String key, String value, long version);

    /**
     * 设置KV，基于最新数据版本
     *
     * @param address 数据账户地址
     * @param key     键
     * @param value   值
     */
    @ContractEvent(name = "setKV")
    void setKV(String address, String key, String value);

    /**
     * 注册用户
     *
     * @param seed 种子，不小于32个字符
     */
    @ContractEvent(name = "registerUser")
    String registerUser(String seed);

    /**
     * 注册数据账户
     *
     * @param seed 种子，不小于32个字符
     */
    @ContractEvent(name = "registerDataAccount")
    String registerDataAccount(String seed);

    /**
     * 注册事件账户
     *
     * @param seed 种子，不小于32个字符
     */
    @ContractEvent(name = "registerEventAccount")
    String registerEventAccount(String seed);

    /**
     * 发布事件
     *
     * @param address  事件账户地址
     * @param topic    消息名称
     * @param content  内容
     * @param sequence 当前消息名称下最大序号（初始为-1）
     */
    @ContractEvent(name = "publishEventWithSequence")
    void publishEventWithSequence(String address, String topic, String content, long sequence);

    /**
     * 发布事件，基于最新时间序号
     *
     * @param address 事件账户地址
     * @param topic   消息名称
     * @param content 内容
     */
    @ContractEvent(name = "publishEvent")
    void publishEvent(String address, String topic, String content);

    /**
     * 交易
     *
     * @result model链接
     */
    @ContractEvent(name = "trade")
    String trade();
}
