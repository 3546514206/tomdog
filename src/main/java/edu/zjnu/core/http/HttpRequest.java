package edu.zjnu.core.http;

import java.util.Map;

/**
 * @description: HttpRequest
 * @author: 杨海波
 * @date: 2022-01-14
 **/
public class HttpRequest {

    private String method;  // 请求方法

    private    String url;     // 请求地址

    private String version; // http版本

    private Map<String, String> headers; // 请求头

    private String body;    // 请求主体

    public String getMethod() {
        return method;
    }

    public void setMethod(String method) {
        this.method = method;
    }

    public String getUrl() {
        return url;
    }

    public void setUrl(String url) {
        this.url = url;
    }

    public String getVersion() {
        return version;
    }

    public void setVersion(String version) {
        this.version = version;
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
