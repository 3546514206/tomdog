package edu.zjnu.core;

import java.util.Map;

/**
 * @description: HttpResponse
 * @author: 杨海波
 * @date: 2022-01-14
 **/
public class HttpResponse {

    private String version; // http版本

    private int code;       // 响应码

    private String status;  // 状态信息

    private Map<String, String> headers; // 响应头

    private String body;    // 响应数据

    public String getVersion() {
        return version;
    }

    public void setVersion(String version) {
        this.version = version;
    }

    public int getCode() {
        return code;
    }

    public void setCode(int code) {
        this.code = code;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public Map<String, String> getHeaders() {
        return headers;
    }

    public void setHeaders(Map<String, String> headers) {
        this.headers = headers;
    }

    public String getBody() {
        return body;
    }

    public void setBody(String body) {
        this.body = body;
    }
}
