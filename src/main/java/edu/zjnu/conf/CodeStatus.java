package edu.zjnu.conf;

/**
 * @description: CodeStatus
 * @author: 杨海波
 * @date: 2022-01-14
 **/
public enum  CodeStatus {

    SUCCESS(200,"SUCCESS");

    CodeStatus(Integer status, String desc) {
        this.status = status;
        this.desc = desc;
    }

    private Integer status;

    private String desc;

    public Integer getStatus() {
        return status;
    }

    public String getDesc() {
        return desc;
    }
}
