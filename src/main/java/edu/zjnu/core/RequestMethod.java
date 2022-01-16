package edu.zjnu.core;

/**
 * @description: 协议方法
 * @author: 杨海波
 * @date: 2022-01-14
 **/
public enum RequestMethod {
    POST("post", "post请求"), GET("get", "get请求");

    private String RequestMethod;
    private String desc;

    RequestMethod(String RequestMethod, String desc) {
        this.RequestMethod = RequestMethod;
        this.desc = desc;
    }

    public String getRequestMethod() {
        return RequestMethod;
    }

    public String getDesc() {
        return desc;
    }
}
